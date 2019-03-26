// +build darwin

package hostid

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
)

// getID returns unique string based on IOPlatformUUID string from ioreg return
func getID() (string, error) {
	out := &bytes.Buffer{}

	cmd := exec.Command("ioreg", "-rd1", "-c", "IOPlatformExpertDevice")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = out

	err := cmd.Run()

	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(out.String(), "\n") {
		if strings.Contains(line, "IOPlatformUUID") {
			return strings.TrimSpace(line), nil
		}
	}

	return "", errors.New("Ca nnot find IOPlatformUUID in result")
}
