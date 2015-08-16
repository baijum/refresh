package main

import (
	"io/ioutil"
	"os"
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
