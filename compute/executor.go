package compute

import (
	"os"
	"os/exec"
)

func ExecuteProgram(code string) error {
	err := os.WriteFile("program.go", []byte(code), 0644)
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "run", "program.go")
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
