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
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Escola CRUD API
// @version 1.0
// @description API para gerenciar alunos em uma escola.

// @host localhost:8080
// @BasePath /

func main() {

	// Carregando o arquivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Erro ao carregar o arquivo .env")
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

	// Iniciando o servidor Gin
	r := gin.Default()

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
