package main

import (
	"testing"
)

func TestGenerateHTMLDoc(t *testing.T) {
	data := generateHTMLDoc("http://feeds.nos.nl/nosnieuwsalgemeen")
	t.Errorf(string(data))
}
