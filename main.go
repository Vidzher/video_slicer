package main

import (
	"fmt"
	"time"
	"video_slicer/slicer"
)

func main() {
	fmt.Println("--- Консольный инструмент раскадровки .mp4 ---")

	for {
		videoPath := printMessage("\nВведите абсолютный путь к видеофайлу:", true)
		outputDir := printMessage("\nВведите имя директории для сохранения результата", true)

		start := time.Now()
		if err := slicer.ExtractFrames(videoPath, outputDir); err != nil {
			fmt.Printf("Ошибка извлечения кадров: %v\n", err)
			return
		}

		elapsedTime := time.Since(start)
		if err := slicer.PrintStats(outputDir, elapsedTime); err != nil {
			fmt.Printf("Невозможно обработать статистику: %v\n", err)
			return
		}

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
