package main

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

// Result структура для хранения результатов обработки
type Result struct {
	totalWords      int
	wordOccurrences int
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

func main() {
	var processor *TextProcessor

	for {
		fmt.Println("\n=== Меню ===")
		fmt.Println("1. Установить путь к файлу")
		fmt.Println("2. Установить слово для поиска")
		fmt.Println("3. Выполнить анализ")
		fmt.Println("0. Выход")

		var choice int
		fmt.Print("\nВыберите действие: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Введите путь к файлу: ")
			var filePath string
			fmt.Scan(&filePath)

			if processor == nil {
				processor = NewTextProcessor(filePath, "")
			} else {
				processor.filePath = filePath
			}
			fmt.Println("Путь к файлу установлен")

		case 2:
			fmt.Print("Введите слово для поиска: ")
			var searchWord string
			fmt.Scan(&searchWord)

			if processor == nil {
				processor = NewTextProcessor("", searchWord)
			} else {
				processor.searchWord = searchWord
			}
			fmt.Println("Слово для поиска установлено")

		case 3:
			if processor == nil {
				fmt.Println("Ошибка: сначала укажите путь к файлу и слово для поиска")
				continue
			}

			result, err := processor.Process()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
				continue
			}

			fmt.Printf("\nРезультаты анализа:")
			fmt.Printf("\nОбщее количество слов: %d", result.totalWords)
			fmt.Printf("\nКоличество повторений слова '%s': %d\n", processor.searchWord, result.wordOccurrences)

		case 0:
			fmt.Println("До свидания!")
			return

		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}
