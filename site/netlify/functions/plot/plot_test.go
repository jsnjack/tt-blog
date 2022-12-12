package plot

import (
	"os"
	"os/exec"
	"testing"
)

func TestGenerateSVG(t *testing.T) {
	data, err := generateSVG(
		"Amy Pond",
		[]string{"Quality", "Ownership", "Speed", "Independence", "Team work", "Reliability"},
		[]float64{9, 8, 7, 6, 7, 8},
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
