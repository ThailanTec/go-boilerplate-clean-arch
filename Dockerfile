### Step 1: Baixar dependenciar e compilar o binario
FROM golang:1.20-alpine as builder

WORKDIR /app

# Copia o go.mod e faz o download das dependencias.
COPY go.mod go.sum ./
RUN go mod download

# Copia o código da aplicação e compila o binario.
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o boilerplate
################################################

### Step 2: Copiar o binario do stage anterior para a imagem final.
FROM scratch

# Copia apenas o binario gerado no stage anterior.
COPY --from=builder /app/boilerplate /

# Define o ponto de entrada para o container como /meuExecutavel.
# O binario será executado quando o container for iniciado.
ENTRYPOINT ["/boilerplate"]