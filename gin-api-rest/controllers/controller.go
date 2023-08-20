package controllers

import (
	"net/http"

	"github.com/edersonSouza02/gin-api-rest/database"
	"github.com/edersonSouza02/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno

	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {

	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"API diz": "E ai " + nome + ", tudo beleza?",
	})

}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error()})
		return
	}
	database.DB.Create(&aluno)

	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno

	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeletarAluno(c *gin.Context) {
	var aluno models.Aluno

	id := c.Params.ByName("id")

	database.DB.Delete(&aluno, id)

	c.JSON(http.StatusOK, gin.H{
		"Data": "Aluno deletado com sucesso"})

}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno

	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
		return

	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ERROR": err.Error()})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)

	c.JSON(http.StatusOK, aluno)

}

func BuscaPorCpf(c *gin.Context) {

	var aluno models.Aluno

	cpf := c.Param("cpf")

	database.DB.Where(&models.Aluno{Cpf: cpf}).First(&aluno)

	if aluno.Cpf == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado por cpf"})
		return
	}

	c.JSON(http.StatusOK, aluno)

}

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno

	database.DB.Find(&alunos)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada(c *gin.Context) {

	c.HTML(http.StatusNotFound, "404.html", nil)

}