package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/to77e/password-generator/tools/aes"
)

func ListNames(cipherKey string) error {
	const (
		perm = 0600
	)
	var (
		tmp        string
		decryptStr string
		values     []string
		err        error
		file       *os.File
	)

	if file, err = os.OpenFile(fileName, os.O_RDONLY, os.FileMode(perm)); err != nil {
		return fmt.Errorf("open file: %v\n", err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmp = scanner.Text()
		if decryptStr, err = aes.Decrypt(tmp, cipherKey); err != nil {
			return fmt.Errorf("decrypt: %w", err)
		}
		values = strings.Split(decryptStr, ";")
		fmt.Printf("%v %v\n", values[0], values[2])
		// fmt.Printf("%v %v %v %v\n", values[0], values[1], values[2], values[3])
	}
	return nil
}
