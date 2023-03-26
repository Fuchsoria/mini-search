package app

type TextData struct {
	Hash     uint64
	Filename string
	Content  string
}

type TextFinded struct {
	Hash    uint64
	Index   int
	Content string
}

type WatcherI interface {
	CachedData() map[uint64]TextData
}

type SearchI interface {
	Find(pattern string) []TextFinded
}
