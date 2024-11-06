# Use uma imagem base do Go com a versão desejada
FROM golang:1.23

# Instalação do Air
RUN go install github.com/air-verse/air@latest

# Configura o diretório de trabalho
WORKDIR /app

# Copia os arquivos de módulo e baixa as dependências (se houver)
COPY go.mod ./
RUN go mod download

# Copia o código fonte para o contêiner
COPY . .

# Exponha a porta que o seu app utiliza
EXPOSE 8080

# Comando padrão para iniciar o Air
CMD ["air"]
