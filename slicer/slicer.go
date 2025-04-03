package slicer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

func GetStats(dir string) (int, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return 0, fmt.Errorf("ошибка чтения директории: %v", err)
	}

	count := 0

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".jpg" {
			count++
		}
	}

	return count, nil
}
