package terminal

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getWidth() (int, error) {
	cmd := exec.Command("tput", "cols")
	var out bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.TrimSpace(out.String()))
}

func getHeight() (int, error) {
	cmd := exec.Command("tput", "lines")
	var out bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.TrimSpace(out.String()))
}

func GetResolution() (int, int, error) {
	width, err := getWidth()
	if err != nil {
		return 0, 0, fmt.Errorf("getWidth: %w", err)
	}
	height, err := getHeight()
	if err != nil {
		return 0, 0, fmt.Errorf("getHeight: %w", err)
	}
	return width, height, nil
}
