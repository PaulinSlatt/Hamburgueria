package controllers

import (
	"GoEcho/database"
	"GoEcho/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Criar um novo burger
func CreateBurger(c echo.Context) error {
	var input struct {
		Nome        string `json:"nome"`
		CarneID     uint   `json:"carne_id"`
		PaoID       uint   `json:"pao_id"`
		OpcionalIDs []uint `json:"opcionais"`
	}
	db := database.GetDB()

	// Bind para associar os dados da requisição ao struct
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados inválidos"})
	}

	// Verificar se a carne existe
	var carne models.Carne
	if err := db.First(&carne, input.CarneID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Carne não encontrada"})
	}

	// Verificar se o pão existe
	var pao models.Pao
	if err := db.First(&pao, input.PaoID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Pão não encontrado"})
	}

	// Buscar os opcionais com base nos IDs
	var opcionais []models.Opcional
	if len(input.OpcionalIDs) > 0 {
		if err := db.Where("id IN ?", input.OpcionalIDs).Find(&opcionais).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Erro ao buscar opcionais"})
		}
	}

	// Definir o status como "Solicitado"
	var status models.Status
	if err := db.Where("tipo = ?", "Solicitado").First(&status).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Erro ao definir o status"})
	}

	// Criar o burger no banco
	burger := models.Burger{
		Nome:      input.Nome,
		CarneID:   input.CarneID,
		PaoID:     input.PaoID,
		Opcionais: opcionais, // Associa os opcionais corretamente
		StatusID:  status.ID,
	}

	if err := db.Create(&burger).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao criar burger"})
	}

	// Atualizar resposta com dados completos
	burger.Carne = carne
	burger.Pao = pao
	burger.Status = status

	return c.JSON(http.StatusCreated, burger)
}

func isValidStatus(status string) bool {
	validStatuses := []string{"Solicitado", "Em produção", "Finalizado"}
	for _, valid := range validStatuses {
		if status == valid {
			return true
		}
	}
	return false
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
	var opcionais []models.Opcional // Usei o tipo correto, que é 'Opcional'

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

	// Estrutura do JSON final
	response := map[string]interface{}{
		"carnes":    []interface{}{}, // Lista de carnes
		"opcionais": []interface{}{}, // Lista de opcionais
		"paes":      []interface{}{}, // Lista de pães
	}

	// Preencher com os dados de carnes
	for _, carne := range carnes {
		response["carnes"] = append(response["carnes"].([]interface{}), map[string]interface{}{
			"id":   carne.ID,
			"tipo": carne.Tipo,
		})
	}

	// Preencher com os dados de pães
	for _, pao := range paes {
		response["paes"] = append(response["paes"].([]interface{}), map[string]interface{}{
			"id":   pao.ID,
			"tipo": pao.Tipo,
		})
	}

	// Preencher com os dados de opcionais
	for _, opcional := range opcionais {
		response["opcionais"] = append(response["opcionais"].([]interface{}), map[string]interface{}{
			"id":   opcional.ID,
			"tipo": opcional.Tipo, // Acessando o campo 'Tipo' da struct Opcional
		})
	}

	// Retornar a resposta no formato esperado
	return c.JSON(http.StatusOK, response)
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
