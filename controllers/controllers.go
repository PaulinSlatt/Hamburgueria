package controllers

import (
	"GoEcho/database"
	"GoEcho/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Criar um novo burger
func CreateBurger(c echo.Context) error {
	var burger models.Burger
	db := database.GetDB()

	// Bind para associar os dados da requisição ao struct
	if err := c.Bind(&burger); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados inválidos"})
	}

	if !isValidCarne(burger.Carne) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Carne inválida"})
	}

	if !isValidPao(burger.Pao) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Pão inválido"})
	}

	burger.Status = "Solicitado"

	if burger.Opcionais == nil {
		burger.Opcionais = []string{}
	}

	if err := db.Create(&burger).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao criar burger"})
	}
	return c.JSON(http.StatusCreated, burger)
}

func isValidCarne(carne string) bool {
	return carne == models.CarneMaminha || carne == models.CarneAlcatra || carne == models.CarnePicanha || carne == models.CarneVeggie
}

func isValidPao(pao string) bool {
	return pao == models.PaoIntegral || pao == models.PaoItalianoBranco || pao == models.PaoTresQueijos || pao == models.PaoParmesaoEOregano
}

func isValidStatus(status string) bool {
	return status == "Solicitado" || status == "Em produção" || status == "Finalizado"
}

func GetBurgers(c echo.Context) error {
	var burgers []models.Burger
	db := database.GetDB()

	if err := db.Find(&burgers).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao consultar burgers"})
	}

	return c.JSON(http.StatusOK, burgers)
}

func UpdateBurgerStatus(c echo.Context) error {
	id := c.Param("id")
	db := database.GetDB()

	// Criar struct temporária para pegar apenas o status
	var input struct {
		Status string `json:"status"`
	}

	if err := c.Bind(&input); err != nil || !isValidStatus(input.Status) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Status inválido"})
	}

	// Verificar se o burger existe
	var burger models.Burger
	if err := db.First(&burger, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Burger não encontrado"})
	}

	// Atualizar apenas o campo Status
	if err := db.Model(&burger).Update("status", input.Status).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao atualizar burger"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Status atualizado com sucesso"})
}

func DeleteBurger(c echo.Context) error {
	id := c.Param("id")
	db := database.GetDB()

	if err := db.Delete(&models.Burger{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao deletar burger"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Burger deletado com sucesso"})
}

// Listar ingredientes

func GetIngredientes(c echo.Context) error {
	db := database.GetDB()

	var paes []models.Pao
	var carnes []models.Carne
	var opcionais []models.Ingrediente

	// Buscar pães
	if err := db.Find(&paes).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao consultar pães"})
	}

	// Buscar carnes
	if err := db.Find(&carnes).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao consultar carnes"})
	}

	// Buscar opcionais
	if err := db.Find(&opcionais).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao consultar opcionais"})
	}

	// Retornar os ingredientes no formato esperado
	return c.JSON(http.StatusOK, map[string]interface{}{
		"paes":      paes,
		"carnes":    carnes,
		"opcionais": opcionais,
	})
}

// Listar pães
//func GetPaes(c echo.Context) error {
//var paes []models.Pao
//db := database.GetDB()

//if err := db.Find(&paes).Error; err != nil {
//return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao consultar pães"})
//}

//return c.JSON(http.StatusOK, paes)
//}

// Listar carnes
//func GetCarnes(c echo.Context) error {
//var carnes []models.Carne
//db := database.GetDB()

//if err := db.Find(&carnes).Error; err != nil {
//	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao consultar carnes"})
//}

//return c.JSON(http.StatusOK, carnes)
//}

// Listar status
func GetStatus(c echo.Context) error {
	var statuses []models.Status
	db := database.GetDB()

	if err := db.Find(&statuses).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao consultar status"})
	}

	return c.JSON(http.StatusOK, statuses)
}
