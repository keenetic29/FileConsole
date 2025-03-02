package model

type Result struct {
	totalWords      int
	wordOccurrences int
}

func (res *Result) GetTotalWords() int {
	return res.totalWords
}

func (res *Result) GetWordOccurences() int {
	return res.wordOccurrences
}