package controllers

import (
	"database/sql"
	"escola-crud/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAlunos retorna todos os alunos
// @Summary Lista todos os alunos
// @Description Retorna todos os alunos cadastrados no banco de dados
// @Tags alunos
// @Produce json
// @Success 200 {array} models.Aluno
// @Router /alunos [get]
func GetAlunos(c *gin.Context, db *sql.DB) {
	alunos, err := models.GetAlunos(db)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": alunos,
	})
}

// CreateAluno cria um novo aluno
// @Summary Cria um novo aluno
// @Description Adiciona um novo aluno ao banco de dados
// @Tags alunos
// @Accept json
// @Produce json
// @Param aluno body models.Aluno true "Dados do aluno"
// @Success 201 {string} string "Aluno criado com sucesso!"
// @Failure 400 {string} string "Erro na criação do aluno"
// @Router /alunos [post]
func CreateAluno(c *gin.Context, db *sql.DB) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.CreateAluno(db, aluno)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Aluno criado com sucesso!",
	})
}

// UpdateAluno atualiza um aluno existente
// @Summary Atualiza um aluno existente
// @Description Atualiza as informações de um aluno específico no banco de dados
// @Tags alunos
// @Accept json
// @Produce json
// @Param id path int true "ID do aluno"
// @Param aluno body models.Aluno true "Dados atualizados do aluno"
// @Success 200 {string} string "Aluno atualizado com sucesso!"
// @Failure 400 {string} string "Erro na atualização do aluno"
// @Router /alunos/{id} [put]
func UpdateAluno(c *gin.Context, db *sql.DB) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	err := models.UpdateAluno(db, id, aluno)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Aluno atualizado com sucesso!",
	})
}

// DeleteAluno remove um aluno pelo ID
// @Summary Remove um aluno
// @Description Exclui um aluno específico do banco de dados
// @Tags alunos
// @Param id path int true "ID do aluno"
// @Success 200 {string} string "Aluno removido com sucesso!"
// @Failure 400 {string} string "Erro na remoção do aluno"
// @Router /alunos/{id} [delete]
func DeleteAluno(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	err := models.DeleteAluno(db, id)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Aluno removido com sucesso!",
	})
}
