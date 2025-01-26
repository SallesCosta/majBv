#!/bin/bash

# Nome do executável
EXEC_NAME="majBv"

# Passo 1: Compilar o projeto Go
echo "Compilando o projeto Go..."
go build -o $EXEC_NAME

# Verificar se a compilação foi bem-sucedida
if [ $? -ne 0 ]; then
    echo "Erro: Falha na compilação!"
    exit 1
fi

# Passo 2: Mover o binário e configurar permissões
echo "Movendo o binário para /usr/local/bin e ajustando permissões..."
sudo mv $EXEC_NAME /usr/local/bin/ && sudo chmod +x /usr/local/bin/$EXEC_NAME

# Verificar se o movimento foi bem-sucedido
if [ $? -eq 0 ]; then
    echo "Deploy concluído com sucesso!"
else
    echo "Erro durante o deploy."
    exit 1
fi
