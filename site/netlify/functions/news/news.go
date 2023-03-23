package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmcdole/gofeed"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const docTemplate = `
{{ define "inc" }}{{ len (printf "%*s " . "") }}{{ end -}}
<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		<title>{{.Title}} {{slice .Published 0 16}}</title>
	</head>
	<body>
		<a id="start"></a>
		<h2>Vandaag, in Nederland:</h2>
		{{range .Items}}
			<ul>
				<li>
					<a href="#{{.GUID}}">{{.Title}}</a>
				</li>
			</ul>
		{{else}}
			No news today. Has the World ended?
		{{end}}
		<mbp:pagebreak/>
		{{range $index, $element := .Items}}
			<article>
				<a id="{{$element.GUID}}"></a>
				<h2>{{$element.Title}}</h2>
				{{$nextIndex := inc $index}}
				{{if lt $nextIndex (len $.Items)}}
					{{$nextEl := index $.Items $nextIndex}}
					<a href="#{{$nextEl.GUID}}">Next article</a>
					<br/>
				{{end}}
				{{renderEnclosures $element.Enclosures}}
				{{safeHTML $element.Description}}
				<dl>
					<dt>Navigate:</dt>
					<dd><a href="{{$element.Link}}" target="_blank">Open the article online</a></dd>
					<dd><a href="#start">Home</a></dd>
				</dl>
				<mbp:pagebreak/>
			</article>
		{{end}}
	</body>
</html>`

// safeHTML marks content as safe HTML, so it is not escaped
func safeHTML(s string) template.HTML {
	return template.HTML(s)
}

func renderEnclosures(en []*gofeed.Enclosure) template.HTML {
	data := ""
	for _, item := range en {
		if strings.HasPrefix(item.Type, "image/") {
			dataURL := toDataURL(item.URL)
			if dataURL != "" {
				data += fmt.Sprintf(`<img src="%s">`, dataURL)
			}
		}
	}
	return template.HTML(data)
}

// toBase64 downloads file and returns it as a data url
func toDataURL(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Unable to download %s: %s\n", url, err)
		return ""
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Unable to read %s: %s\n", url, err)
		return ""
	}

	encoded := fmt.Sprintf("data:%s;base64,%s", http.DetectContentType(bytes), base64.StdEncoding.EncodeToString(bytes))
	return encoded
}

// generateHTMLDoc generates HTML document from RSS news feed
func generateHTMLDoc(rssFeedURL string) []byte {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(rssFeedURL)
	if err != nil {
		log.Fatal(err)
	}

	t := template.Must(template.New("news").Funcs(map[string]interface{}{
		"safeHTML":         safeHTML,
		"renderEnclosures": renderEnclosures,
		"inc": func(i int) int {
			return i + 1
		},
	}).Parse(docTemplate))

	var buf bytes.Buffer
	err = t.Execute(&buf, feed)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

// extractDomain extracts domain from RSS URL
func extractDomain(feedURL string) string {
	u, err := url.Parse(feedURL)
	if err != nil {
		fmt.Println(err)
		return "RSS feed"
	}
	splitted := strings.Split(u.Hostname(), ".")
	if len(splitted) >= 2 {
		return strings.Join(splitted[len(splitted)-2:], ".")
	}
	return "RSS feed"
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// request.QueryStringParameters https://github.com/aws/aws-lambda-go/blob/main/events/apigw.go
	started := time.Now()
	feedURL := "http://feeds.nos.nl/nosnieuwsalgemeen"

	docBytes := generateHTMLDoc(feedURL)

	title := extractDomain(feedURL) + " " + time.Now().Format("2006-01-02")

	// Sending email
	m := mail.NewV3Mail()

	from := mail.NewEmail("Yauhen's Netlify", "noreply@yauhen.cc")
	content := mail.NewContent("text/html", title)
	to := mail.NewEmail("Yauhen", "jsnjack@kindle.com")

	m.SetFrom(from)
	m.AddContent(content)

	// create new *Personalization
	personalization := mail.NewPersonalization()
	personalization.AddTos(to)
	personalization.Subject = title

	// add `personalization` to `m`
	m.AddPersonalizations(personalization)

	attachment := mail.NewAttachment()
	encoded := base64.StdEncoding.EncodeToString(docBytes)
	attachment.SetContent(encoded)
	attachment.SetType("text/html")
	attachment.SetFilename(title + ".html")
	attachment.SetDisposition("attachment")

	m.AddAttachment(attachment)

	req := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	req.Method = "POST"
	req.Body = mail.GetRequestBody(m)
	resp, err := sendgrid.API(req)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Body)
		fmt.Println(resp.Headers)
	}

	respText := fmt.Sprintf(
		"Generated %s and sent to your kindle in %s",
		title,
		time.Since(started),
	)
	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/plain"},
		Body:            respText,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call
	lambda.Start(handler)
}
