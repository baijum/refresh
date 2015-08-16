package runner

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func build() (string, bool) {
	buildLog("Building...")

	cmd := exec.Command("go", "build", "-o", buildPath(), root())

	stderr, err := cmd.StderrPipe()
	if err != nil {
		logger.Fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logger.Fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		logger.Fatal(err)
	}

	io.Copy(os.Stdout, stdout)
	errBuf, _ := ioutil.ReadAll(stderr)

	err = cmd.Wait()
	if err != nil {
		return string(errBuf), false
	}

	return "", true
}
