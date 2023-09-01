# Imagem base do Go
FROM golang:alpine

# Defina a variável de ambiente GOPATH
ENV GOPATH=/go

# Copie os arquivos do projeto para o diretório de trabalho
WORKDIR $GOPATH/src/app
COPY . .

# Instale as dependências
RUN go mod download 

# Instale o Fresh
RUN go install github.com/pilu/fresh@latest

# Compile o código da aplicação
RUN go build -o app

# Execute a aplicação
CMD if [ "$WATCH_FILES" == "1" ]; then fresh; else ./app; fi