package searchalg

type BoyerMoore struct{}

func (this *BoyerMoore) Search(text, pattern string) int {
	shift := this.CreateShift(pattern)
	t := 0
	last := len(pattern) - 1

	for t < len(text)-last {
		p := last
		for p >= 0 && text[t+p] == pattern[p] {
			p--
		}

		if p == -1 {
			return t
		}

		t += shift[text[t+last]]
	}

	return -1
}

func (this *BoyerMoore) CreateShift(pattern string) []int {
	shift := make([]int, 256)

	for j := 0; j < len(shift); j++ {
		shift[j] = len(pattern)
	}

	for p := 0; p < len(pattern)-1; p++ {
		shift[pattern[p]] = len(pattern) - p - 1
	}

	return shift
}

func NewBM() *BoyerMoore {
	return &BoyerMoore{}
}
