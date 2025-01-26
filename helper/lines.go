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

	`  ["<c-p>"] = { telescope.find_files, "Find file" },`,
	`  ["<c-f>"] = { telescope.live_grep, "Find word" },`,
	`  ["<c-n>"] = { "<cmd>NvimTreeToggle<cr>", "Open file explorer", noremap = true },`,

	`  helpers.should_load_theme("catppuccin", { "catppuccin/nvim", name = "catppuccin" }),`,
	`  helpers.should_load_theme("dracula", { "Mofiqul/dracula.nvim" }),`,
	`  helpers.should_load_theme("ayu", { "ayu-theme/ayu-vim" }),`,
	`  helpers.should_load_theme("palenight", { "drewtempelmeyer/palenight.vim" }),`,
	`  helpers.should_load_theme("tokyonight", { "folke/tokyonight.nvim" }),`,
	`  helpers.should_load_theme("nord", { "arcticicestudio/nord-vim" }),`,
	`  helpers.should_load_theme("onedark", { "olimorris/onedarkpro.nvim" }),`,

	`  { "<leader>cl", "<cmd>nohl<cr>", desc = " Clear search highlights" },`,

	`  { "<leader>ev", "<cmd>e $MYVIMRC<cr>", desc = " Open init.lua" },`,
	`  { "<leader>eb", "<cmd>e " .. constants.CONFIG_PATH .. "/better-vim.lua<cr>", desc = " Open better-vim.lua" },`,

	`  { "<leader>c", group = "LSP / Highlights" },`,
	`  { "<leader>e", group = "Configuration" },`,
	`	 { "<leader>w", group = "LSP Workspace" },`,
	`  { "<leader>w", group = "LSP Workspace" },`,
	`  { "<leader>q", "<cmd>:Bdelete<cr>", desc = " Close buffer, not window", noremap = true },`,
}
