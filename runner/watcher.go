package runner

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/fsnotify.v1"
)

func watchFolder(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fatal(err)
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Events:
				if isWatchedFile(ev.Name) {
					watcherLog("sending event %s", ev)
					startChannel <- ev.String()
				}
			case err := <-watcher.Errors:
				watcherLog("error: %s", err)
			}
		}
	}()

	watcherLog("Watching %s", path)
	err = watcher.Add(path)

	if err != nil {
		fatal(err)
	}
}

func watch() {
	root := root()
	ex := excludeDir()

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && !isTmpDir(path) {
			if len(path) > 1 && strings.HasPrefix(filepath.Base(path), ".") {
				return filepath.SkipDir
			}

			if len(ex) > 0 {
				for _, ep := range strings.Split(ex, ",") {
					if strings.Contains(path, strings.TrimSpace(ep)) {
						// Skip this path
						watcherLog("Excluding %s", path)
						return filepath.SkipDir
					}
				}
			}

			watchFolder(path)
		}

		return err
	})
}
