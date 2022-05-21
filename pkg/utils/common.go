package utils

import (
	"bufio"
	"log"
	"os/exec"
	"strings"
)

func String2Lines(s string) ([]string, error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err := scanner.Err()

	return lines, err
}

func ExecCmd(cmd string) string {
	cmdStruct := exec.Command("bash", "-c", cmd)
	rst, err := cmdStruct.Output()
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(rst))
}
