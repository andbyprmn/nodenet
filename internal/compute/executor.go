package compute

import (
	"errors"
	"os/exec"
)

func ExecuteProgram(programCode string) (string, error) {
	// Contoh sederhana mengeksekusi program Go
	cmd := exec.Command("go", "run", programCode)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New("execution failed: " + string(output))
	}
	return string(output), nil
}
