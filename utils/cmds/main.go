package cmds

import (
	"bytes"
	"os/exec"
	"strings"
)

func ExecCMD(cmdStr string) (string, error) {
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.Trim(out.String(), "\n"), nil
}
