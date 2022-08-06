package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestGenerateHTMLDoc(t *testing.T) {
	data := generateHTMLDoc("http://feeds.nos.nl/nosnieuwsalgemeen")
	err := os.WriteFile("output.html", data, 0644)
	if err != nil {
		t.Error(err)
		return
	}
	exec.Command("xdg-open", "output.html").Start()
}

func TestExtractDomain(t *testing.T) {
	result := extractDomain("http://feeds.nos.nl/nosnieuwsalgemeen")
	expected := "nos.nl"
	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}
