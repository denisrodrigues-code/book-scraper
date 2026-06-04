# =====================================================
# Stage 1 - Build da aplicação
# =====================================================
# Utiliza uma imagem leve do Go baseada em Alpine
FROM golang:1.26-alpine AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos de dependência primeiro
# para aproveitar o cache do Docker
COPY go.mod go.sum ./

# Baixa as dependências do projeto
RUN go mod download

# Copia o restante do código-fonte
COPY . .

# Compila a aplicação
# CGO_ENABLED=0 gera um binário estático
RUN CGO_ENABLED=0 GOOS=linux go build \
    -o scraper \
    ./cmd/scraper

# =====================================================
# Stage 2 - Runtime
# =====================================================
# Imagem final extremamente leve
FROM alpine:latest

# Cria um usuário sem privilégios
RUN addgroup -S scraper && \
    adduser -S scraper -G scraper

# Define o diretório de trabalho
WORKDIR /app

# Cria pasta para armazenar os arquivos gerados
RUN mkdir -p data

# Copia o binário compilado do estágio anterior
COPY --from=builder /app/scraper .

# Ajusta permissões
RUN chown -R scraper:scraper /app

# Executa a aplicação com usuário não-root
USER scraper

# O scraper não expõe API,
# mas o desafio pede exposição de porta.
# Utilizamos uma porta padrão documentada.
EXPOSE 8080

# Comando de inicialização
CMD ["./scraper"]