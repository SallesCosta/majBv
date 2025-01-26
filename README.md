# majBv - Mettre à Jour BetterVim

majBv is a Go-based utility designed to automate the process of updating and 
customizing Neovim configurations ([BetterVim](https://bettervim.com/)). It is particularly useful for 
developers who frequently update their Neovim plugins and prefer to modify some of the default commands and functionalities.

## Features

- Automatic removal of specific lines in configuration files.
- Supports multiple configuration files.
- Easy to customize the lines to be removed.

## Usage

After building the project, you can run the utility with the following command:

```bash
./majBv
```

This command will update the following Neovim configuration files by removing the specified lines:

- `lsp.lua`
- `shortcuts.lua`
- `plugins.lua`
- `cmp.lua`

## Customization

The lines to be removed are defined in the `helper/lines.go` file. You can modify the `LinesToRemove` variable to customize the lines you want the utility to remove.

## Building the Project

To build the project, you need to have Go installed on your system. You can then use the `go build` command to build the project.

## 

`sudo mv majBv /usr/local/bin/ && sudo chmod +x /usr/local/bin/majBv`

## Contributing

Contributions are welcome. Please feel free to submit a pull request or open an issue.
