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
