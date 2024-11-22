package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rclone/rclone/fs/config/obscure"
)

func main() {
	configPath := os.ExpandEnv("$HOME/.config/rclone/rclone.conf")

	file, err := os.Open(configPath)
	if err != nil {
		fmt.Printf("Error opening rclone.conf: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "password = ") || strings.HasPrefix(line, "secret = ") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				obscuredValue := strings.TrimSpace(parts[1])

				revealedValue := obscure.MustReveal(obscuredValue)
				fmt.Printf("%s = %s\n", key, revealedValue)
			}
		} else {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading rclone.conf: %v\n", err)
	}
}
