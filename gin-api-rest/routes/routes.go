package routes

import (
	"github.com/edersonSouza02/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/alunos", controllers.ExibeTodosAlunos)

	r.GET("/index", controllers.ExibePaginaIndex)

	r.GET("/:nome", controllers.Saudacao)

	r.NoRoute(controllers.RotaNaoEncontrada)

	r.POST("/alunos", controllers.CriaNovoAluno)

	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)

	r.GET("/alunos/cpf/:cpf", controllers.BuscaPorCpf)

	r.PATCH("/alunos/:id", controllers.EditaAluno)

	r.DELETE("/alunos/:id", controllers.DeletarAluno)

	r.Run()
}
