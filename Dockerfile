# Estágio de construção
FROM golang:1.22-alpine AS stage1

WORKDIR /app

# Copie os arquivos de módulo e baixe as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copie o restante do código-fonte
COPY . .

# Compile o aplicativo
RUN CGO_ENABLED=0 GOOS=linux go build -o myBinary ./cmd/app

# Estágio final
FROM scratch

COPY --from=stage1 /app/myBinary /

ENTRYPOINT ["/myBinary"]