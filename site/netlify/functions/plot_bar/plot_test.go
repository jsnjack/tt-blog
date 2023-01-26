package main

import (
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func TestGenerateSVG1(t *testing.T) {
	chData := &ChartData{
		Title:  "Test Chart",
		Legend: []string{"min", "max", "p99"},
		Keys:   []string{"haproxy1.8", "haproxy2.1"},
		Values: [][]float64{
			{10, 100, 30},
			{20, 250, 60},
		},
		Format: "svg",
		Type:   "bar",
	}
	data, err := generateImage(chData)
	if err != nil {
		t.Error(err)
		return
	}
	err = os.WriteFile("output1.svg", data, 0644)
	if err != nil {
		t.Error(err)
		return
	}
	exec.Command("xdg-open", "output1.svg").Start()
}

func TestParseQuery(t *testing.T) {
	parsedQueryMap := map[string]string{
		"type":      "bar",
		"title":     "Haproxy+response+duration",
		"legend":    "min,max,p99",
		"haproxy18": "10,20,19",
		"haproxy20": "9,20,17",
		"format":    "png",
	}
	chData, err := parseQuery(parsedQueryMap)
	if err != nil {
		t.Error(err)
		return
	}
	expected := &ChartData{
		Title:  "Haproxy+response+duration",
		Legend: []string{"min", "max", "p99"},
		Keys:   []string{"haproxy18", "haproxy20"},
		Values: [][]float64{
			{10, 9},
			{20, 20},
			{19, 17},
		},
		Format: "png",
		Type:   "bar",
	}
	if !reflect.DeepEqual(chData, expected) {
		t.Errorf("Expected %v, got %v", expected, chData)
		return
	}

}
