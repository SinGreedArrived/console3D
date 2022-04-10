package terminal

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getWidth() (uint64, error) {
	cmd := exec.Command("tput", "cols")
	var out bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return 0, err
	}
	return strconv.ParseUint(strings.TrimSpace(out.String()), 10, 64)
}

func getHeight() (uint64, error) {
	cmd := exec.Command("tput", "lines")
	var out bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return 0, err
	}
	return strconv.ParseUint(strings.TrimSpace(out.String()), 10, 64)
}

func GetResolution() (uint64, uint64, error) {
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
