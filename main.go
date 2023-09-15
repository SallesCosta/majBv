package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// filePath := "~/.config/nvim/lua/better-vim-core/lsp.lua" // Caminho estático para o arquivo
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	filePath := filepath.Join(homeDir, ".config", "nvim", "lua", "better-vim-core", "lsp.lua")

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("deu erro: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Print(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}

}
