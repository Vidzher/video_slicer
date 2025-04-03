package main

import (
	"fmt"
	"video_slicer/slicer"
)

func main() {
	fmt.Println("--- Консольный инструмент раскадровки .mp4 ---")

	for {
		videoPath := printMessage("\nВведите абсолютный путь к видеофайлу:", true)
		outputDir := printMessage("\nВведите имя директории для сохранения результата", true)

		if err := slicer.ExtractFrames(videoPath, outputDir); err != nil {
			fmt.Printf("Ошибка извлечения кадров: %v\n", err)
			return
		}

		result, err := slicer.GetStats("output")
		if err != nil {
			fmt.Printf("Невозможно обработать статистику файлов: %v\n", err)
			return
		}

		fmt.Printf("Сохранено кадров: %d\n", result)
	}
}

func printMessage(message string, isScan bool) (input string) {
	fmt.Println(message)
	if isScan {
		fmt.Scanln(&input)

		return input
	}

	return ""
}
