package runner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRoot(t *testing.T) {
	r := root()
	if r != "." {
		t.Error("'root' is not '.'")
	}
}

func TestTmpPath(t *testing.T) {
	tmp := tmpPath()
	if tmp != "./tmp" {
		t.Error("'tmp_path' is not './tmp'")
	}
}

func TestBuildName(t *testing.T) {
	bname := buildName()
	if bname != "runner-build" {
		t.Error("'build_name' is not 'runner-build'")
	}
}

func TestBuildPath(t *testing.T) {
	bpath := buildPath()
	if bpath != filepath.Join("./tmp", "runner-build") {
		t.Error("'build_path' is not './tmp/runner-build'")
	}
}

func TestBuildErrorsFileName(t *testing.T) {
	blog := buildErrorsFileName()
	if blog != "runner-build-errors.log" {
		t.Error("'build_log' is not 'runner-build-errors.log'")
	}
}

func TestBuildErrorsFilePath(t *testing.T) {
	logPath := buildErrorsFilePath()
	if logPath != filepath.Join("./tmp", "runner-build-errors.log") {
		t.Error("'build_log_path' is not './tmp/runner-build-errors.log'")
	}
}

func TestConfigPath(t *testing.T) {
	conf := configPath()
	if conf != "./.refresh.conf" {
		t.Error("'config_path' is not './.refresh.conf'")
	}
}

func TestExcludeDir(t *testing.T) {
	exclude := excludeDir()
	if exclude != "" {
		t.Error("'exclude_dir' is not empty")
	}
}

func TestBuildDelay(t *testing.T) {
	delay := buildDelay()
	if delay != 600 {
		t.Error("Wrong delay:", delay)
	}
}

func TestBuildDelayWrongValue(t *testing.T) {
	os.Setenv("REFRESH_BUILD_DELAY", "wrong")
	loadEnvSettings()
	delay := buildDelay()
	if delay != 600 {
		t.Error("Wrong delay:", delay)
	}
}

func TestBuildDelayWithValue(t *testing.T) {
	os.Setenv("REFRESH_BUILD_DELAY", "700")
	loadEnvSettings()
	delay := buildDelay()
	if delay != 700 {
		t.Error("Wrong delay:", delay)
	}
}
