package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestSetConfigPath(t *testing.T) {
	tmp, _ := ioutil.TempDir("", "refresh-test-")
	inConf := filepath.Join(tmp, ".refresh.conf")
	ioutil.WriteFile(inConf, []byte(""), 0644)
	setConfigPath(inConf)
	outConf := os.Getenv("REFRESH_CONFIG_PATH")
	if inConf != outConf {
		t.Error("Configuration not set")
	}
	os.RemoveAll(tmp)
}

func TestSetConfigPathWrongPath(t *testing.T) {
	tmp, _ := ioutil.TempDir("", "refresh-test-")
	inConf := filepath.Join(tmp, ".refresh.conf")
	err := setConfigPath(inConf)
	if err == nil {
		t.Error("Configuration set")
	}
	os.RemoveAll(tmp)
}

func TestMainNonExistingConf(t *testing.T) {
	cmd := exec.Command(os.Args[0], "-c", "/tmp/non-existing-conf")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Error("Process ran with err %q, want exit status 1", err)
}
