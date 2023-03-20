# Payment Transfer API

## O que é o projeto ?

 API para transferência monetaria, baseada em um desafio da [DevGym](https://app.devgym.com.br/challenges/9af13172-e1fe-4c2e-ac10-cb6b0bcf2efc)

## Tecnologias usadas

- Golang 1.20
- Postgres
- Docker

## Como rodar esse projeto de forma local ?

1. Clone esse repositorio
    - Via HTTP `git clone  https://github.com/RickChaves29/payment_tranfer_api.git`
    - Via SSH `git@github.com:RickChaves29/payment_tranfer_api.git`

2. Ainda no terminal, copie a variável de ambiente que está no arquivo .env.example e cole no arquivo .bashrc ou .profile adicionando a palavra chave export antes.

    > OBS: O Arquivo .bashrc ou .profile fica na pasta raiz do seu úsuario
    - Exemplo no WSL ou linux

    ```bash
    export PAYMENT_DB='<url de conexão do db>'
    ```

3. Voltando para pasta onde você clonou o projeto rode os seguintes comandos:

     - Baixar todas as dependências `go mod download`
     - Rodar o projeto `go run server/main.go`

    >OBS: caso não tenha setado a variavel de ambiente use o comando

    `export PAYMENT_DB='<url de conexão do db>' && go run server/main.go`

## Como rodar o projeto apartir da imagem **Docker**

1. Puxe a imagem no [Docker Hub](https://hub.docker.com/r/rickchaves29/payment_transfer_api)

    `docker pull rickchaves29/payment_transfer_api:<tag de versão>`

2. Crie um container baseado na imagem

    ```bash
    docker run --name 'nome do container' -e PAYMENT_DB='<numero da porta>' \
    -p 4040:'<numero da porta>'/tcp rickchaves29/payment_transfer_api:'tag version'
    ```

## Como rodar o projeto usando o **Docker Compose** no modo de desenvolvimento

>OBS: Por Padrão o Docker Compose sempre ir chamar o arquivo **compose.yaml**

1. Clonar esse repositório pelo terminal
   - Via HTTP
     `git clone https://github.com/RickChaves29/payment_transfer_api.git`
   - Via SSH
     `git clone git@github.com:RickChaves29/payment_transfer_api.git`
2. Ao entrar na pasta do projeto rode o seguinte comando no terminal

   - Copie do arquivo .env.example as variaveis de ambiente e crie o arquivo .env e cole as variaveis, setando seus valores
   - Para subir os containers `docker compose up --remove-orphans`
   - Para remove os containers `docker compose down`

  >OBS: A flag **--remove-orphans** é opcinal, ela remove todos snapshots criados pelo comando up

## Criação da tabela no banco de daos

> OBS: Caso não exista a tabela no banco de dados, a tabela é criada automaticamente pela aplicação

## Rotas da API

### GET - /api/v1/users/:id

Retorna um json com base no id enviado na uri:

```json
{
    "id": "0000",
    "balance": 1000
}
```

### GET - /api/v1/transfer

Recebe um json:

```json
{
    "payer": "0000",
    "receive": "1111",
    "amount": 0
}
```

## Usuários para testar a API

```json
[
    {
        "id": "1234",
        "balance": 1000
    }
    {
        "id": "4321",
        "balance": 2000
    }
]
```
