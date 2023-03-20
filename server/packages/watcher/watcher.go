package watcher

import "github.com/radovskyb/watcher"

type Settings struct {
	Folder string
}

type Watcher struct {
	w        *watcher.Watcher
	settings Settings
}

func New(settings Settings) *Watcher {
	return &Watcher{
		w:        watcher.New(),
		settings: settings,
	}
}
