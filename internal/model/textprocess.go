package model

import (
	"fmt"
	"os"
	"strings"
)

// TextProcessor структура для обработки текстовых файлов
type TextProcessor struct {
	filePath   string
	searchWord string
}

// NewTextProcessor создает новый экземпляр TextProcessor
func NewTextProcessor(filePath, searchWord string) *TextProcessor {
	return &TextProcessor{
		filePath:   filePath,
		searchWord: strings.ToLower(searchWord),
	}
}

// ReadFile читает содержимое файла
func (tp *TextProcessor) ReadFile() ([]byte, error) {
	return os.ReadFile(tp.filePath)
}

// CountWords подсчитывает общее количество слов в тексте
func (tp *TextProcessor) CountWords(content []byte) int {
	words := strings.Fields(string(content))
	return len(words)
}

// SearchWord ищет слово в тексте и возвращает количество повторений
func (tp *TextProcessor) SearchWord(content []byte) int {
	words := strings.Fields(string(content))
	count := 0

	for _, word := range words {
		if strings.ToLower(word) == tp.searchWord {
			count++
		}
	}

	return count
}

// Process выполняет анализ текста и возвращает результаты
func (processor *TextProcessor) Process() (*Result, error) {
	content, err := processor.ReadFile()
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении файла: %w", err)
	}

	totalWords := processor.CountWords(content)
	occurrences := processor.SearchWord(content)

	return &Result{
		totalWords:      totalWords,
		wordOccurrences: occurrences,
	}, nil
}

// GetFilePath возвращает путь к файлу
func (tp *TextProcessor) GetFilePath() string {
    return tp.filePath
}

// GetSearchWord возвращает слово для поиска
func (tp *TextProcessor) GetSearchWord() string {
    return tp.searchWord
}

// SetFilePath устанавливает путь к файлу
func (tp *TextProcessor) SetFilePath(filePath string) {
    tp.filePath = filePath
}

// SetSearchWord устанавливает слово для поиска
func (tp *TextProcessor) SetSearchWord(searchWord string) {
    tp.searchWord = strings.ToLower(searchWord)
}
