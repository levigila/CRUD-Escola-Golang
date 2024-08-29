# **Escola CRUD API**

Este é um projeto de **API em Golang** para gerenciar alunos em uma escola. A API permite realizar operações de **CRUD** (Create, Read, Update, Delete) sobre os dados dos alunos, armazenados em um banco de dados **PostgreSQL**.

## **Índice**

- [Visão Geral](#visão-geral)
- [Funcionalidades](#funcionalidades)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Configuração do Ambiente](#configuração-do-ambiente)
- [Deploy](#deploy)
- [Documentação da API](#documentação-da-api)
- [Contribuição](#contribuição)
- [Licença](#licença)

## **Visão Geral**

A **Escola CRUD API** foi desenvolvida para gerenciar informações de alunos, como nome, idade, notas dos semestres, nome do professor e número da sala. O projeto segue a arquitetura **MVC** (Model-View-Controller) e foi implementado utilizando a linguagem **Golang** e **PostgreSQL** como banco de dados.

## **Funcionalidades**

- **Listar Alunos:** Retorna todos os alunos cadastrados.
- **Adicionar Aluno:** Adiciona um novo aluno ao banco de dados.
- **Atualizar Aluno:** Atualiza as informações de um aluno específico.
- **Remover Aluno:** Exclui um aluno do banco de dados.

## **Tecnologias Utilizadas**

- **Golang:** Linguagem principal para desenvolvimento da API.
- **PostgreSQL:** Banco de dados utilizado para armazenar as informações dos alunos.
- **Swagger:** Documentação interativa da API.
- **Docker:** Para facilitar a configuração e execução do ambiente.
- **Render:** Plataforma utilizada para o deploy da aplicação.

## **Configuração do Ambiente**

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/seu-usuario/seu-repositorio.git
   cd seu-repositorio
   ```

2. **Configuração do Banco de Dados:**
   - Certifique-se de que o PostgreSQL está instalado e configurado.
   - Crie um banco de dados chamado `golangschool_6d8b`.
   - Atualize as variáveis de ambiente no arquivo `.env` com as informações do banco de dados.

3. **Rodar a aplicação localmente:**
   - Instale as dependências necessárias:
     ```bash
     go mod tidy
     ```
   - Execute a aplicação:
     ```bash
     go run main.go
     ```

4. **Configuração com Docker:**
   - Construa e execute o container Docker:
     ```bash
     docker-compose up --build
     ```

## **Deploy**

O deploy foi realizado na plataforma [Render](https://render.com). Para acessar a API em produção, utilize a seguinte URL:

```
https://golangschoolcrud.onrender.com
```

## **Documentação da API**

A documentação interativa da API está disponível através do **Swagger**. Para acessá-la, basta acessar a seguinte URL após rodar a aplicação localmente:

```
http://localhost:8080/swagger/index.html
```

## **Contribuição**

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.

## **Licença**

Este projeto está licenciado sob a **Licença MIT**. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---
