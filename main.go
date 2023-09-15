package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

var linesToRemove = []string{
	`  nmap("<leader>rn", vim.lsp.buf.rename, " Rename symbol")`,
	`  nmap("<leader>ca", vim.lsp.buf.code_action, " Show code actions")`,
	`  nmap("<leader>gd", vim.lsp.buf.definition, " Go to definition")`,
	`  nmap("<leader>cr", telescope.lsp_references, " Show references")`,
	`  nmap("<leader>ci", vim.lsp.buf.implementation, " Show implementation")`,
	`  nmap("<leader>cy", vim.lsp.buf.type_definition, " Type definition")`,
	`  nmap("<leader>ds", telescope.lsp_document_symbols, "[D]ocument [S]ymbols")`,
	`  nmap("<leader>ws", telescope.lsp_dynamic_workspace_symbols, "[W]orkspace [S]ymbols")`,
	`  nmap("<leader>cd", vim.lsp.buf.hover, " Show documentation")`,
	`  nmap("<leader>cs", vim.lsp.buf.signature_help, "Signature Documentation")`,
	`  nmap("<leader>ll", vim.lsp.codelens.run, "Run code lens action")`,
	`  -- Lesser used LSP functionality`,
	`  nmap("<leader>gD", vim.lsp.buf.declaration, "[G]oto [D]eclaration")`,
	`  nmap("<leader>wa", vim.lsp.buf.add_workspace_folder, "[W]orkspace [A]dd Folder")`,
	`  nmap("<leader>wr", vim.lsp.buf.remove_workspace_folder, "[W]orkspace [R]emove Folder")`,
	`  nmap("<leader>wl", function()`,
	`    print(vim.inspect(vim.lsp.buf.list_workspace_folders()))`,
	`  end, "[W]orkspace [L]ist Folders")`,
}

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}

	filePath := filepath.Join(homeDir, ".config", "nvim", "lua", "better-vim-core", "lsp.lua")
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

	fmt.Println("Linhas removidas com sucesso.")
}
