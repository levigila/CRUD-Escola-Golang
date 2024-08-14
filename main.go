package main

import (
	"database/sql"
	"escola-crud/controllers"
	_ "escola-crud/docs" // Importa a documentação gerada pelo Swag
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Escola CRUD API
// @version 1.0
// @description API para gerenciar alunos em uma escola.

// @host golangschoolcrud.onrender.com
// @BasePath /
// @schemes https

func main() {

	// Carrega o arquivo .env, se ele existir
    if err := godotenv.Load(); err != nil {
        log.Println("Arquivo .env não encontrado, carregando variáveis de ambiente padrão")
    }

	// Usando variáveis para armazenar os valores das variáveis de ambiente
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Construindo a string de conexão
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Conectando ao banco de dados
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verificando a conexão
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão bem-sucedida ao banco de dados!")

	// Criando tabela "alunos"
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS alunos (
            id SERIAL PRIMARY KEY,
            nome VARCHAR(100),
            idade INT,
            nomeProfessor VARCHAR(100),
            notaPrimeiroSemestre FLOAT,
            notaSegundoSemestre FLOAT,
            numeroSala INT
        );
    `)
    if err != nil {
        log.Fatal(err)
    }

	// Iniciando o servidor Gin
	r := gin.Default()

	// Configuração básica de CORS
    r.Use(cors.Default())

	// Documentação Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Definindo as rotas e os controladores
	r.GET("/alunos", func(c *gin.Context) {
		controllers.GetAlunos(c, db)
	})

	r.POST("/alunos", func(c *gin.Context) {
		controllers.CreateAluno(c, db)
	})

	r.PUT("/alunos/:id", func(c *gin.Context) {
		controllers.UpdateAluno(c, db)
	})

	r.DELETE("/alunos/:id", func(c *gin.Context) {
		controllers.DeleteAluno(c, db)
	})

	// Rodando o servidor
	r.Run()
}
