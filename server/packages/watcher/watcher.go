package watcher

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"

	"github.com/radovskyb/watcher"
)

type Settings struct {
	Folder    string
	FileTypes *regexp.Regexp
}

type Watcher struct {
	instance *watcher.Watcher
	settings Settings
}

func (w *Watcher) RunFilesChecking() error {
	log.Print(fmt.Sprintf("Checking through all folders in %s starts...", w.settings.Folder))

	w.instance.AddFilterHook(watcher.RegexFilterHook(w.settings.FileTypes, false))

	if _, err := os.Stat(w.settings.Folder); os.IsNotExist(err) {
		err := os.Mkdir(w.settings.Folder, 0o755)
		if err != nil {
			return err
		}
	}

	if err := w.instance.AddRecursive(w.settings.Folder); err != nil {
		return err
	}

	watchedFiles := w.instance.WatchedFiles()

	log.Print("All files were checked", "total", len(watchedFiles))

	return nil
}

func (w *Watcher) GetFiles() map[string]fs.FileInfo {
	return w.instance.WatchedFiles()
}

func (w *Watcher) Len() int {
	return len(w.instance.WatchedFiles())
}

func New(settings Settings) *Watcher {
	return &Watcher{
		instance: watcher.New(),
		settings: settings,
	}
}
