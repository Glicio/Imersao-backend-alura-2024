Projeto do api para a imersão backend da Alura

## Disclaimer

O projeto da Alura usa JavaScript e MongoDB, mas estou utilizando Go para tentar aprender melhor essa linguagem e PostgreSQL pq não gosto de noSQL.

## Como rodar

1. Clone o projeto 
    `git clone https://github.com/Glicio/Imersao-backend-alura-2024.git`
    ou 
    `gh repo clone Glicio/Imersao-backend-alura-2024`

2. Crie um arquivo .env com as seguintes variáveis na raiz do projeto

exemplo:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
```

3. Rode o comando `go mod download` e `go mod tidy` para baixar as dependências

4. Rode o comando `go run cmd/service/main.go`

5. Acesse a api em `http://localhost:3000/post`

6. Acesse a página inicial em `http://localhost:3000/`

7. Acesse a pagina de upload em `http://localhost:3000/upload`

## Aula 4: Implementando Armazenamento e Upload de Imagens

    1. criado e servido arquivo html para a página index, disponibilizado em http://localhost:3000
       com um formulário simples para envio de uimagens junto com dados para criação de posts
    2. criada rota upload, que recebe a imagem, salva em disco e criar um post no banco de dados
