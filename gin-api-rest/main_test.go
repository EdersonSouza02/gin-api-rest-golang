package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/edersonSouza02/gin-api-rest/controllers"
	"github.com/edersonSouza02/gin-api-rest/database"
	"github.com/edersonSouza02/gin-api-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()

	return rotas

}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()

	r.GET("/:nome", controllers.Saudacao)

	req, _ := http.NewRequest("GET", "/ederson", nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")

	mockDaResposta := `{"API diz":"E ai ederson, tudo beleza?"}`

	respostaBody, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mockDaResposta, string(respostaBody))

	fmt.Println(string(respostaBody))

	fmt.Println(mockDaResposta)
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()

	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)

	fmt.Println(resposta.Body)
}

func CriaAlunoMock() {

	aluno := models.Aluno{Nome: "NomeMock", Cpf: "11111111111", RG: "111111111"}

	database.DB.Create(&aluno)

	ID = int(aluno.ID)

}
func DeletaAlunoMock() {
	var aluno models.Aluno

	database.DB.Delete(&aluno, ID)
}

func TestBuscaAlunoPorCpfHandler(t *testing.T) {

	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()

	r.GET("/alunos/cpf/:cpf", controllers.BuscaPorCpf)

	req, _ := http.NewRequest("GET", "/alunos/cpf/11111111111", nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)

}

func TestBuscaAlunoPorIdHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()

	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)

	patchDaBusca := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("GET", patchDaBusca, nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	var AlunoMock models.Aluno

	json.Unmarshal(resposta.Body.Bytes(), &AlunoMock)

	assert.Equal(t, "NomeMock", AlunoMock.Nome, "Os nomes devem ser iguais")

	assert.Equal(t, "11111111111", AlunoMock.Cpf)

	assert.Equal(t, "111111111", AlunoMock.RG)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()

	r := SetupDasRotasDeTeste()

	r.DELETE("/alunos/:id", controllers.DeletarAluno)

	patchDeBusca := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", patchDeBusca, nil)

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)

}

func TestEditaUmAlunoHandler(t *testing.T) {

	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()

	r.PATCH("/alunos/:id", controllers.EditaAluno)

	aluno := models.Aluno{Nome: "NomeMock", Cpf: "11111111119", RG: "111111119"}

	valorJson, _ := json.Marshal(aluno)

	patchParaEditar := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", patchParaEditar, bytes.NewBuffer(valorJson))

	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	var alunoMockAtualizado models.Aluno

	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)

	assert.Equal(t, "11111111119", alunoMockAtualizado.Cpf)
	assert.Equal(t, "111111119", alunoMockAtualizado.RG)
	assert.Equal(t, "NomeMock", alunoMockAtualizado.Nome)

}
