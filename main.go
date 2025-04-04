package main

import (
	"fmt"
	"video_slicer/converter"
	"video_slicer/slicer"
	"video_slicer/utils"
)

func init() {
	if !utils.CheckFFMPEG() {
		utils.InstallFFMPEG()
		fmt.Println("Установка завершена! Перезапустите приложение...")
		return
	}
}

func main() {
	fmt.Println("--- Консольный инструмент раскадровки .mp4 ---")

	for {
		getUserInput()
	}
}

func getUserInput() {
	fmt.Println("\nДля продолжения введите номер команды:")
	fmt.Println("1. Раскадровка видео")
	fmt.Println("2. Конвертировать видео в .mp4:")

	var input string
	fmt.Scanln(&input)

	switch input {
	case "1":
		slicer.Start()
	case "2":
		converter.Start()
	default:
		fmt.Println("Команда не найдена")
	}
}
