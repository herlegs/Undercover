package logic

import (
	"sync"
	"errors"
)

type GameCache struct{
	sync.RWMutex
	words []string
	index int
}

func (cache *GameCache) generateWords(word1, word2 string, majorityNum, minorityNum int){
	words := make([]string, majorityNum + minorityNum)
	for i := 0; i < majorityNum; i++ {
		words = append(words, word1)
	}
	for i := 0; i < minorityNum; i++ {
		words = append(words, word2)
	}
	cache.words = shuffle(words)
	cache.index = 0
}

func (cache *GameCache) getWord() (string, error){
	cache.Lock()
	defer cache.Unlock()
	if cache.index < len(cache.words){
		word := cache.words[cache.index]
		cache.index++
		return word, nil
	}
	return "", errors.New("Room is full")
}

//roomName -- wordList
var roomCache = make(map[string]*GameCache)