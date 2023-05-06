package word

import (
	"sync"
)

type WordMatcher struct {
	Words map[string]struct{}
	mutex sync.RWMutex
}

func NewMatcher() *WordMatcher {
	return &WordMatcher{
		Words: make(map[string]struct{}),
	}
}

func (w *WordMatcher) AddWords(words ...string) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	for _, word := range words {
		w.Words[word] = struct{}{}
	}
}

func (w *WordMatcher) DeleteWords(words ...string) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	for _, word := range words {
		delete(w.Words, word)
	}
}

func (w *WordMatcher) CheckWords(words ...string) []bool {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	boolArr := make([]bool, len(words))

	for i, word := range words {
		_, boolArr[i] = w.Words[word]
	}

	return boolArr
}
