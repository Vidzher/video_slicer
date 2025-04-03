package slicer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func ExtractFrames(videoPath, outputDir string) error {
	if _, err := os.Stat(videoPath); os.IsNotExist(err) {
		return fmt.Errorf("файл %s не найден", videoPath)
	}

	if err := os.MkdirAll(outputDir, os.FileMode(0755)); err != nil {
		return fmt.Errorf("ошибка создания директории: %v", err)
	}

	outputPattern := filepath.Join(outputDir, "frame_%04d.jpg")

	cmd := exec.Command(
		"ffmpeg",
		"-i", videoPath,
		"-q:v", "2",
		outputPattern,
	)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("не удалось извлечь кадры: %v", err)
	}

	return nil
}

func PrintStats(dir string, elapsedTime time.Duration) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("ошибка чтения директории: %v", err)
	}

	count := 0

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".jpg" {
			count++
		}
	}

	fmt.Printf("Сохранено кадров: %d\n", count)
	fmt.Printf("Время работы: %.2f секунд\n", elapsedTime.Seconds())

	return nil
}
