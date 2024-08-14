# Golang User CRUD Boilerplate

Este projeto é um boilerplate desenvolvido em Golang, contendo um CRUD completo de usuários, utilizando PostgreSQL como banco de dados relacional e Redis para caching.

## Arquitetura

A aplicação segue os princípios da Clean Architecture, utilizando Domain-Driven Design (DDD) para organizar o código de forma modular e sustentável.

### Estrutura do Projeto

A estrutura de pastas do projeto é organizada da seguinte forma:

### Descrição das Pastas

- **cmd/app**: Ponto de entrada da aplicação.
- **domain**: Entidades de negócio e interfaces principais.
- **infra**: Implementações de banco de dados, cache e outras integrações externas.
    - `auth`: Implementações de autenticação e autorização.
    - `database`: Configurações e conexão com o PostgreSQL, incluindo migrações.
    - `repositories`: Implementações de repositórios para acesso ao banco de dados.
- **src**: Lógica de negócio e camadas de interface.
    - `config`: Configurações da aplicação, como variáveis de ambiente e inicializações.
    - `dto`: Objetos usados para transferir dados entre camadas da aplicação.
    - `handlers`: Implementações dos handlers HTTP.
    - `middleware`: Implementações de middlewares.
    - `routes`: Definições das rotas da aplicação.
    - `usecases`: Lógica de negócio e interações entre entidades.
- **test**: Testes unitários e mocks para a aplicação.
    - `mocks`: Implementações de mocks para repositórios e use cases.
        - `repositories`: Mocks para repositórios.
        - `usecases`: Mocks para use cases.
    - `unit`: Testes unitários organizados por camadas.
        - `infra`: Testes unitários para a camada infra.
            - `test_respositories/repositories`: Testes unitários específicos para repositórios.
        - `src`: Testes unitários para a camada src.
            - `handlers`: Testes unitários para handlers.
            - `usecases`: Testes unitários para use cases.

## Requisitos

- **Golang**: A versão mais recente do Golang.
- **PostgreSQL**: Configurado como banco de dados relacional.
- **Redis**: Utilizado como cache para otimização de operações de leitura.

## Configuração e Inicialização

### Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto com as seguintes configurações:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdbname

REDIS_HOST=localhost
REDIS_PORT=6379

JWTSecret=exemplo
JWTExpirationMinutes=60
```

## Makefile
Para iniciar o projeto:

```bash 
make run
```
Para verificar os testes e também verificar a qualidade do código:

```bash 
make linter
```

