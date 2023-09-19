package main

import (
	"github.com/sallescosta/majBv/helper"
)

func main() {
	helper.Atualizar("lsp.lua", helper.LinesToRemove)
	helper.Atualizar("shortcuts.lua", helper.LinesToRemove)
	helper.Atualizar("plugins.lua", helper.LinesToRemove)
}
