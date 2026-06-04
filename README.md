# 📚 Book Scraper (Go)

Projeto de web scraping desenvolvido em Go para extração de dados do site  
https://books.toscrape.com/

O sistema coleta informações estruturadas de livros, processa os dados e exporta em formatos JSON e CSV.  
Além disso, inclui pipeline CI/CD, testes automatizados e containerização com Docker.

---

# 🚀 Como rodar o projeto

## ▶️ Rodando localmente (sem Docker)

### 1. Clonar o repositório
```bash
git clone https://github.com/seu-usuario/book-scraper.git
cd book-scraper
2. Baixar dependências
go mod tidy
3. Executar o scraper
go run cmd/scraper/main.go
📦 Saída

Os arquivos serão gerados em:

data/books.json
data/books.csv
🐳 Rodando com Docker
1. Build da imagem
docker build -t book-scraper .
2. Executar container
docker run --rm book-scraper

📌 Os dados serão gerados dentro do container em /app/data.

📊 Estrutura dos dados
JSON Schema

Cada livro segue o seguinte formato:

{
  "title": "string",
  "price": "float64",
  "rating": "int (1-5)",
  "availability": "string",
  "image_url": "string",
  "product_url": "string"
}

Exemplo:

{
  "title": "A Light in the Attic",
  "price": 51.77,
  "rating": 3,
  "availability": "In stock",
  "image_url": "https://books.toscrape.com/media/...",
  "product_url": "https://books.toscrape.com/catalogue/..."
}
CSV Schema

O CSV segue a seguinte estrutura:

title,price,rating,availability,image_url,product_url

Exemplo:

A Light in the Attic,51.77,3,In stock,https://...,https://...
⚙️ Como o pipeline funciona

O scraper segue uma arquitetura em pipeline:

main.go
  ↓
ScrapeBooks (paginação de 1 a 50 páginas)
  ↓
ScrapePage (requisição HTTP + parsing HTML)
  ↓
Transformações:
  - parsePrice (conversão £ → float64)
  - ratingToNumber (texto → int)
  - buildImageURL (normalização URL imagem)
  - buildProductURL (normalização URL produto)
  ↓
models.Book (estrutura de dados)
  ↓
Exportação:
  - JSON (encoding/json)
  - CSV (encoding/csv)
🧠 Detalhes do pipeline
Rate limit de 1 segundo entre páginas para evitar bloqueios
HTTP client com timeout de 10 segundos
User-Agent customizado
Parsing com goquery baseado em seletores CSS
🧠 Decisões técnicas
1. goquery para parsing HTML

Escolhido por ser leve, estável e baseado em seletores CSS, facilitando scraping estruturado.

2. Separação por responsabilidades

O projeto foi dividido em:

scraper (extração)
exporter (persistência)
models (estrutura de dados)
3. Pipeline sequencial

Optei por execução sequencial para:

evitar bloqueio do site
simplificar controle de erro
facilitar debug
4. Rate limiting simples

Uso de time.Sleep(1s) entre requests para reduzir risco de bloqueio.

5. Conversão de dados
Preço tratado manualmente (£ → float64)
Rating convertido de texto para número via mapping
🔮 O que eu faria com mais tempo
🚀 Performance
Implementar scraping concorrente com goroutines + worker pool
Fila de URLs (queue system)
🛡 Anti-bot
Rotação de User-Agent
Proxy rotation
Backoff exponencial em retries
📊 Observabilidade
Logs estruturados (zap/logrus)
Métricas (Prometheus)
Tracing (OpenTelemetry)
💾 Persistência
Salvar dados em PostgreSQL ou MongoDB
Histórico de scraping
🤖 IA
Classificação automática de livros
Resumo de descrição com LLM
Detecção de duplicados inteligente
🤖 Uso de IA no desenvolvimento

Durante o desenvolvimento utilizei IA para:

✔ Estruturação do projeto
Definição da arquitetura (scraper → transformer → exporter)
✔ Debug e correção de parsing
Ajuste de funções como parsePrice e ratingToNumber
✔ Boas práticas
Sugestão de separação em packages (scraper, exporter, models)
✔ Docker e CI/CD
Criação de Dockerfile multi-stage
Estrutura de pipeline GitLab CI
✔ Testes
Geração de testes unitários para funções de parsing
O que funcionou bem
Aceleração na estruturação do projeto
Redução de erros comuns de arquitetura
Melhor organização do código
Limitações observadas
Necessidade de revisão manual para detalhes de scraping
Ajustes finos de CSS selectors ainda dependem de inspeção humana
👨‍💻 Autor

Projeto desenvolvido como parte de um desafio técnico para prática de:

Web scraping em Go
Arquitetura de software
CI/CD
Docker