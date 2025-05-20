# LoginGo-React-teste

## Descrição

Este projeto é uma aplicação fullstack composta por um backend em Go (Golang) utilizando o framework Gin e um frontend em React. O objetivo é fornecer um sistema simples de cadastro, autenticação e listagem de usuários.

---

## Tecnologias Utilizadas

- **Backend:** Go, Gin, GORM, PostgreSQL
- **Frontend:** React, TypeScript, Axios, React Router DOM

---

## Como rodar o projeto

### Pré-requisitos

- [Go](https://golang.org/dl/) instalado
- [Node.js](https://nodejs.org/) e [npm](https://www.npmjs.com/) instalados
- [PostgreSQL](https://www.postgresql.org/download/) instalado e rodando

### Backend

1. Instale as dependências Go:
    ```sh
    go mod tidy
    ```
2. Configure o banco de dados PostgreSQL e ajuste a string de conexão em `backend/db/postgres.go` se necessário.
3. Rode o backend:
    ```sh
    go run main.go
    ```
    O backend estará disponível em `http://localhost:5000`.

### Frontend

1. Acesse a pasta [frontend](http://_vscodecontentref_/0):
    ```sh
    cd frontend
    ```
2. Instale as dependências:
    ```sh
    npm install
    ```
3. Rode o frontend:
    ```sh
    npm start
    ```
    O frontend estará disponível em `http://localhost:3000`.

---

## Funcionalidades

- Cadastro de usuário
- Login/autenticação
- Listagem de usuários (exemplo) #não implementado
- Proteção de rotas no frontend

---


