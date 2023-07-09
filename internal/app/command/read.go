package command

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/to77e/paslok/tools/aes"
)

func ReadName(cipherKey, filePath, name string) error {
	const (
		perm = 0600
	)

	file, err := os.OpenFile(filepath.Clean(filePath), os.O_RDONLY, os.FileMode(perm))
	if err != nil {
		return fmt.Errorf("failed to open file: %v\n", err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	var (
		values     []string
		password   string
		decryptStr string
	)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmp := scanner.Text()
		if decryptStr, err = aes.Decrypt(tmp, cipherKey); err != nil {
			return fmt.Errorf("decrypt: %w", err)
		}
		values = strings.Split(decryptStr, ";")
		if values[0] == name {
			password = values[1]
		}
	}

	if err = scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan file: %v", err)
	}
	if len(password) == 0 {
		return fmt.Errorf("name %v doesn't exist", name)
	}

	if err = clipboard.WriteAll(password); err != nil {
		return fmt.Errorf("faild to clipboard password: %v", err)
	}
	return nil
}
