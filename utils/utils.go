package utils

import (
	"fmt"
	"os/exec"
)

func CheckFFMPEG() bool {
	_, err := exec.LookPath("ffmpeg")
	return err == nil
}

func InstallFFMPEG() {
	fmt.Println("Установка FFmpeg через winget...")
	cmd := exec.Command(
		"winget", "install", "-e", "--id", "Gyan.FFmpeg",
		"--silent", "--accept-package-agreements", "--accept-source-agreements",
	)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("ошибка winget: %v", err)
	}
}

func PrintMessage(message string, isScan bool) (input string) {
	fmt.Println(message)
	if isScan {
		fmt.Scanln(&input)

		return input
	}

	return ""
}
