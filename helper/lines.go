package helper

var LinesToRemove = []string{
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

	`  ["<leader>cl"] = { "<cmd>nohl<cr>", " Clear search highlights" },`,
	`  ["<c-p>"] = { telescope.find_files, "Find file" },`,
	`  ["<c-f>"] = { telescope.live_grep, "Find word" },`,
	`  ["<c-n>"] = { "<cmd>NvimTreeToggle<cr>", "Open file explorer", noremap = true },`,
	`  ["<leader>ev"] = { "<cmd>e $MYVIMRC<cr>", " Open init.lua" },`,
	`  ["<leader>eb"] = { "<cmd>e " .. constants.CONFIG_PATH .. "/better-vim.lua<cr>", " Open better-vim.lua" },`,
	`  ["<leader>q"] = { "<cmd>:Bdelete<cr>", " Close buffer, not window", noremap = true },`,

	`  helpers.should_load_theme("catppuccin", utils.load_theme({ "catppuccin/nvim", name = "catppuccin" })),`,
	`  helpers.should_load_theme("dracula", utils.load_theme({ "Mofiqul/dracula.nvim" })),`,
	`  helpers.should_load_theme("ayu", utils.load_theme({ "ayu-theme/ayu-vim" })),`,
	`  helpers.should_load_theme("palenight", utils.load_theme({ "drewtempelmeyer/palenight.vim" })),`,
	`  helpers.should_load_theme("tokyonight", utils.load_theme({ "folke/tokyonight.nvim" })),`,
	`  helpers.should_load_theme("nord", utils.load_theme({ "arcticicestudio/nord-vim" })),`,
	`  helpers.should_load_theme("onedarkpro", utils.load_theme({ "olimorris/onedarkpro.nvim" })),`,
}
