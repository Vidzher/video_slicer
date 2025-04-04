package converter

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
	"video_slicer/utils"
)

func convertToMP4(initSrc, targetSrc string) error {
	if _, err := os.Stat(initSrc); os.IsNotExist(err) {
		return fmt.Errorf("файл %s не найден", initSrc)
	}

	outputDir := "converter"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("не удалось создать директорию: %v", err)
	}

	outputPath := filepath.Join(outputDir, targetSrc+".mp4")

	cmd := exec.Command(
		"ffmpeg",
		"-i", initSrc,
		"-c:v", "libx264",
		"-preset", "fast",
		"-crf", "23",
		"-c:a", "aac",
		"-b:a", "192k",
		outputPath,
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("не удалось конвертировать видео: %v", err)
	}

	return nil
}

func Start() {
	time.Sleep(1000)
	initSrc := utils.PrintMessage("\nВведите абсолютный путь к видеофайлу:", true)
	targetSrc := utils.PrintMessage("\nВведите имя нового файла:", true)

	fmt.Println("Видео конвертируется...")
	if err := convertToMP4(initSrc, targetSrc); err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
}
