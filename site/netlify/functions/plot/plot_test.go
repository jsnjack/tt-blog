package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestGenerateSVG1(t *testing.T) {
	data, err := generateImage(
		"Amy Pond",
		[]string{"Quality", "Ownership", "Speed", "Independence", "Team work", "Reliability"},
		[][]float64{{5, 8, 7, 9, 7, 9}},
		nil,
		"svg",
	)
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

func TestGenerateSVG2(t *testing.T) {
	data, err := generateImage(
		"Amy Pond",
		[]string{"Quality", "Ownership", "Speed", "Independence", "Team work", "Reliability"},
		[][]float64{{9, 8, 7, 6, 7, 8}, {5, 8, 7, 9, 7, 9}},
		[]string{"Q1", "Q2"},
		"svg",
	)
	if err != nil {
		t.Error(err)
		return
	}
	err = os.WriteFile("output2.svg", data, 0644)
	if err != nil {
		t.Error(err)
		return
	}
	exec.Command("xdg-open", "output2.svg").Start()
}

func TestGeneratePNG(t *testing.T) {
	data, err := generateImage(
		"Amy Pond",
		[]string{"Quality", "Ownership", "Speed", "Independence", "Team work", "Reliability"},
		[][]float64{{9, 8, 7, 6, 7, 8}, {5, 8, 7, 9, 7, 9}},
		[]string{"Q1", "Q2"},
		"png",
	)
	if err != nil {
		t.Error(err)
		return
	}
	err = os.WriteFile("output1.png", data, 0644)
	if err != nil {
		t.Error(err)
		return
	}
	exec.Command("xdg-open", "output1.png").Start()
}

func TestDataExtractor1(t *testing.T) {
	data := map[string]string{
		"name":      "Amy Pond",
		"quality":   "a+",
		"Ownership": "A",
	}
	name, skills, scores, legend, format := dataExtractor(data)
	if name != "Amy Pond" {
		t.Error("expected Amy Pond, got", name)
		return
	}
	if len(skills) != 2 {
		t.Error("expected 2 skills, got", len(skills), skills)
		return
	}
	if skills[0] != "Ownership" {
		t.Error("expected Ownership, got", skills[1])
		return
	}
	if skills[1] != "Quality" {
		t.Error("expected Quality, got", skills[0])
		return
	}
	if len(scores) != 1 {
		t.Error("expected 1 score, got", len(scores))
		return
	}
	if fmt.Sprintf("%v", scores[0]) != "[8 9]" {
		t.Error("expected [8 9], got", scores[0])
		return
	}
	if len(legend) != 0 {
		t.Error("expected 0 legend, got", len(legend))
		return
	}
	if format != "svg" {
		t.Error("expected svg, got", format)
		return
	}
}
func TestDataExtractor2(t *testing.T) {
	data := map[string]string{
		"name":    "Amy Pond",
		"quality": "a, b",
	}
	name, skills, scores, legend, format := dataExtractor(data)
	if name != "Amy Pond" {
		t.Error("name is not correct")
		return
	}
	if len(skills) != 1 {
		t.Error("expected 1 skill, got", len(skills))
		return
	}
	if skills[0] != "Quality" {
		t.Error("expected Quality, got", skills[0])
		return
	}
	if len(scores) != 2 {
		t.Error("expected 2 scores, got", len(scores))
		return
	}
	if fmt.Sprintf("%v", scores[0]) != "[8]" {
		t.Error("expected [8], got", scores[0])
		return
	}
	if fmt.Sprintf("%v", scores[1]) != "[5]" {
		t.Error("expected [5], got", scores[1])
		return
	}
	if len(legend) != 0 {
		t.Error("expected 0 legend, got", len(legend))
		return
	}
	if format != "svg" {
		t.Error("expected svg, got", format)
		return
	}
}

func TestDataExtractor3(t *testing.T) {
	data := map[string]string{
		"name":    "Amy Pond",
		"quality": "a,b",
		"speed":   "a",
		"type":    "png",
	}
	name, skills, scores, legend, format := dataExtractor(data)
	if name != "Amy Pond" {
		t.Error("name is not correct")
		return
	}
	if len(skills) != 2 {
		t.Error("expected 2 skills, got", len(skills))
		return
	}
	if skills[0] != "Quality" {
		t.Error("expected Quality, got", skills[0])
		return
	}
	if skills[1] != "Speed" {
		t.Error("expected Speed, got", skills[1])
		return
	}
	if len(scores) != 2 {
		t.Error("expected 2 scores, got", len(scores))
		return
	}
	if fmt.Sprintf("%v", scores[0]) != "[8 8]" {
		t.Error("expected [8 8], got", scores[0])
		return
	}
	if fmt.Sprintf("%v", scores[1]) != "[5]" {
		t.Error("expected [5], got", scores[1])
		return
	}
	if len(legend) != 0 {
		t.Error("expected 0 legend, got", len(legend))
		return
	}
	if format != "png" {
		t.Error("expected png, got", format)
		return
	}
}
