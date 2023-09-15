package helper

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func Atualizar(fileName string, linesToRemove []string) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}

	filePath := filepath.Join(homeDir, ".config", "nvim", "lua", "better-vim-core", fileName)
	newFilePath := filePath + ".new"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("deu erro: ", err)
		return
	}

	newFile, err := os.Create(newFilePath)
	if err != nil {
		fmt.Println("Erro ao criar o novo arquivo:", err)
		return
	}

	defer file.Close()
	defer newFile.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		removeLine := false

		for _, lineToRemove := range linesToRemove {
			if line == lineToRemove {
				removeLine = true
				break
			}
		}

		if !removeLine {
			_, err := fmt.Fprintln(newFile, line)
			if err != nil {
				fmt.Println("Erro ao escrever no novo arquivo:", err)
				return
			}
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	err = os.Rename(newFilePath, filePath)
	if err != nil {
		fmt.Println("Erro ao renomear o novo arquivo:", err)
		return
	}

	fmt.Println(fileName, "atualizado.")

}
