package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmcdole/gofeed"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const docTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
		<title>{{.Title}} {{slice .Published 0 16}}</title>
	</head>
	<body>
		{{range .Items}}
		<h2>Today, in the Netherlands:</h2>
		<ul>
			<li>
				<a href="#{{.GUID}}">{{{{.Title}}}}</a>
			</li>
		</ul>
		<mbp:pagebreak/>
		{{end}}
		{{range .Items}}
			<article>
				<h2>{{.Title}}</h2>
				<a id={{.GUID}}></a>
				{{safeHTML .Description}}
				<a href="{{.Link}}" target="_blank">Open the article</a>
				<mbp:pagebreak/>
			</article>
		{{else}}
			No news today. Has the World ended?
		{{end}}
	</body>
</html>`

func safeHTML(s string) template.HTML {
	return template.HTML(s)
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// request.QueryStringParameters https://github.com/aws/aws-lambda-go/blob/main/events/apigw.go

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("http://feeds.nos.nl/nosnieuwsalgemeen")
	if err != nil {
		log.Fatal(err)
	}

	t := template.Must(template.New("news").Funcs(map[string]interface{}{"safeHTML": safeHTML}).Parse(docTemplate))
	var buf bytes.Buffer
	err = t.Execute(&buf, feed)
	if err != nil {
		log.Fatal(err)
	}

	title := feed.Title + " " + feed.Published

	// Sending email
	m := mail.NewV3Mail()

	from := mail.NewEmail("Yauhen's Netlify", "noreply@yauhen.space")
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
	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
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

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/plain"},
		Body:            fmt.Sprintf("Collected %d articles", len(feed.Items)),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call
	lambda.Start(handler)
}
