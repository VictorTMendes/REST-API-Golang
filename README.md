# 🧩 REST-API-Golang

Uma API REST desenvolvida em **Go (Golang)**, com arquitetura limpa e modular, organizada em camadas (`controller`, `usecase`, `repository`, `model`) e preparada para rodar em ambiente **Docker**.  
Este projeto foi criado para estudos e como base para o desenvolvimento de aplicações mais completas.

---

## 🚀 Funcionalidades

- ✅ CRUD completo (Create, Read, Update (in progress), Delete(in progress))
- 🧠 Separação clara de responsabilidades (Controller → Usecase → Repository)
- 🐳 Configuração pronta com **Docker Compose**
- 🧱 Modelos e repositórios desacoplados
- ⚙️ Fácil manutenção e escalabilidade

---

## 📂 Estrutura de Pastas

```bash
.
├── cmd/               # Ponto de entrada da aplicação
├── controller/        # Camada responsável pelos handlers (entrada HTTP)
├── usecase/           # Camada de regras de negócio
├── repository/        # Camada de persistência de dados
├── model/             # Estruturas e entidades do domínio
├── db/                # Configurações e scripts de banco de dados
├── docker-compose.yml # Configuração para ambiente Docker
├── go.mod             # Definições de dependências do projeto
└── go.sum             # Checksum das dependências
```

## 🧰 Tecnologias Utilizadas

- [Go](https://go.dev/) — Linguagem principal  
- [Docker](https://www.docker.com/) — Containerização  
- [Docker Compose](https://docs.docker.com/compose/) — Orquestração de serviços  
- [PostgreSQL](https://www.postgresql.org/) — Banco de dados relacional (padrão)  
- [net/http](https://pkg.go.dev/net/http) — Pacote nativo do Go para servidor HTTP  
