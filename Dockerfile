# Estágio de construção
FROM golang:1.22-alpine AS stage1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myBinary ./cmd/app

# Estágio final
FROM scratch

COPY --from=stage1 /app/myBinary /

ENTRYPOINT ["/myBinary"]