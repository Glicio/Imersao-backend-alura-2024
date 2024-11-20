Projeto do api para a imersão backend da Alura

## Disclaimer

O projeto da Alura usa JavaScript e MongoDB, mas estou utilizando Go para tentar aprender melhor essa linguagem e PostgreSQL pq não gosto de noSQL.

## Como rodar

1. Crie um arquivo .env com as seguintes variáveis

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
```

2. Rode o comando `go run cmd/service/main.go`

3. Acesse a api em `http://localhost:3000/api`

4. Para testar a api, use o comando `curl -X POST -H "Content-Type: application/json" -d '{"key": "secret"}' http://localhost:3000/api`
