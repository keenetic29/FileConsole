package app

import (
	"fmt"
	"os"

	"file.com/internal/model"
)

func Run(){
	var processor *model.TextProcessor

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
				processor = model.NewTextProcessor(filePath, "")
			} else {
				processor.SetFilePath(filePath)
			}
			fmt.Println("Путь к файлу установлен")

		case 2:
			fmt.Print("Введите слово для поиска: ")
			var searchWord string
			fmt.Scan(&searchWord)

			if processor == nil {
				processor = model.NewTextProcessor("", searchWord)
			} else {
				processor.SetSearchWord(searchWord)
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
			fmt.Printf("\nОбщее количество слов: %d", result.GetTotalWords())
			fmt.Printf("\nКоличество повторений слова '%s': %d\n", processor.GetSearchWord(), result.GetWordOccurences())

		case 0:
			fmt.Println("До свидания!")
			return

		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}