package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestGenerateSVG1(t *testing.T) {
	data, err := generateSVG(
		"Amy Pond",
		[]string{"Quality", "Ownership", "Speed", "Independence", "Team work", "Reliability"},
		[][]float64{{9, 8, 7, 6, 7, 8}, {5, 8, 7, 9, 7, 9}},
		nil,
	)
	if err != nil {
		t.Error(err)
		return
	}
	err = os.WriteFile("output.svg", data, 0644)
	if err != nil {
		t.Error(err)
		return
	}
	exec.Command("xdg-open", "output.svg").Start()
}

func TestGenerateSVG2(t *testing.T) {
	data, err := generateSVG(
		"Amy Pond",
		[]string{"Quality", "Ownership", "Speed", "Independence", "Team work", "Reliability"},
		[][]float64{{9, 8, 7, 6, 7, 8}, {5, 8, 7, 9, 7, 9}},
		[]string{"Q1", "Q2"},
	)
	if err != nil {
		t.Error(err)
		return
	}
	err = os.WriteFile("output.svg", data, 0644)
	if err != nil {
		t.Error(err)
		return
	}
	exec.Command("xdg-open", "output.svg").Start()
}

func TestDataExtractor1(t *testing.T) {
	data := map[string]string{
		"name":      "Amy Pond",
		"quality":   "a+",
		"Ownership": "A",
	}
	name, skills, scores, legend := dataExtractor(data)
	if name != "Amy Pond" {
		t.Error("expected Amy Pond, got", name)
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
	if skills[1] != "Ownership" {
		t.Error("expected Ownership, got", skills[1])
		return
	}
	if len(scores) != 1 {
		t.Error("expected 1 score, got", len(scores))
		return
	}
	if fmt.Sprintf("%v", scores[0]) != "[9 8]" {
		t.Error("expected [9 8], got", scores[0])
		return
	}
	if len(legend) != 0 {
		t.Error("expected 0 legend, got", len(legend))
		return
	}
}
func TestDataExtractor2(t *testing.T) {
	data := map[string]string{
		"name":    "Amy Pond",
		"quality": "a, b",
	}
	name, skills, scores, legend := dataExtractor(data)
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

}

func TestDataExtractor3(t *testing.T) {
	data := map[string]string{
		"name":    "Amy Pond",
		"quality": "a,b",
		"speed":   "a",
	}
	name, skills, scores, legend := dataExtractor(data)
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
}
