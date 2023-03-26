package searchalg

import (
	"minisearch/server/packages/app"
)

type Settings struct{}

type Search struct {
	watcher  app.WatcherI
	settings Settings
	bm       *BoyerMoore
}

func (s *Search) Find(pattern string) []app.TextFinded {
	texts := s.watcher.CachedData()
	results := []app.TextFinded{}

	for hash, text := range texts {
		index := s.bm.Search(text.Content, pattern)

		if index > 0 {
			results = append(results, app.TextFinded{Hash: hash, Index: index, Content: texts[hash].Content})
		}
	}

	return results
}

func New(settings Settings, watcher app.WatcherI) *Search {
	return &Search{
		settings: settings,
		watcher:  watcher,
		bm:       NewBM(),
	}
}
