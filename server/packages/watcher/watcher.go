package watcher

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"minisearch/server/packages/app"

	"github.com/cespare/xxhash"
	"github.com/radovskyb/watcher"
)

type Settings struct {
	Folder    string
	FileTypes *regexp.Regexp
}

type Watcher struct {
	data     map[uint64]app.TextData
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

func (w *Watcher) getFileContent(path string) (error, string, uint64) {
	file, err := os.Open(path)
	if err != nil {
		return err, "", 0
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err, "", 0
	}

	hash := xxhash.Sum64(bytes)

	return nil, string(bytes), hash
}

func (w *Watcher) CacheData() error {
	files := w.GetFiles()

	for _, v := range files {
		path := filepath.Join(w.settings.Folder, v.Name())

		err, content, hash := w.getFileContent(path)
		if err != nil {
			fmt.Println(err, path)

			continue
		}

		w.data[hash] = app.TextData{
			Hash:     hash,
			Filename: v.Name(),
			Content:  content,
		}
	}

	return nil
}

func (w *Watcher) CachedData() map[uint64]app.TextData {
	return w.data
}

func (w *Watcher) Len() int {
	return len(w.instance.WatchedFiles())
}

func New(settings Settings) *Watcher {
	return &Watcher{
		instance: watcher.New(),
		settings: settings,
		data:     make(map[uint64]app.TextData),
	}
}
