package controllers

import (
	"GoEcho/database"
	"GoEcho/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Criar um novo burger
func CreateBurger(c echo.Context) error {
	var burger models.Burger
	db := database.GetDB()

	// 🔹 Bind do JSON
	if err := c.Bind(&burger); err != nil {
		fmt.Println("Erro ao bindar JSON:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados inválidos"})
	}

	// 🔹 Buscar a carne
	var carne models.Carne
	if err := db.First(&carne, burger.CarneID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Carne não encontrada"})
	}
	burger.CarneNome = carne.Tipo

	// 🔹 Buscar o pão
	var pao models.Pao
	if err := db.First(&pao, burger.PaoID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Pão não encontrado"})
	}
	burger.PaoNome = pao.Tipo

	// 🔹 Buscar os opcionais corretamente
	var opcionais []models.Opcional
	if len(burger.Opcionais) > 0 {
		var opcionalIDs []uint
		for _, opcional := range burger.Opcionais {
			opcionalIDs = append(opcionalIDs, opcional.ID)
		}

		if err := db.Where("id IN ?", opcionalIDs).Find(&opcionais).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Erro ao buscar opcionais"})
		}

		burger.Opcionais = opcionais // Associa os opcionais encontrados
	}

	// 🔹 Definir status automaticamente
	var status models.Status
	if err := db.Where("tipo = ?", "Solicitado").First(&status).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao definir status"})
	}
	burger.StatusID = status.ID
	burger.StatusNome = status.Tipo

	// 🔹 Criar o burger no banco de dados
	if err := db.Create(&burger).Error; err != nil {
		fmt.Println("Erro ao salvar no banco:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao criar burger"})
	}

	// 🔹 Retornar o burger criado, incluindo os opcionais
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

	// Bind e validação do JSON
	if err := c.Bind(&input); err != nil || input.Status == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Status inválido"})
	}

	// Buscar o status pelo nome
	var status models.Status
	if err := db.Where("tipo = ?", input.Status).First(&status).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Status não encontrado"})
	}

	// Verificar se o burger existe
	var burger models.Burger
	if err := db.First(&burger, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Burger não encontrado"})
	}

	// Atualizar o status no banco de dados
	if err := db.Model(&burger).Updates(map[string]interface{}{
		"status_id":   status.ID,
		"status_nome": status.Tipo,
	}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao atualizar burger"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Status atualizado com sucesso"})
}

func DeleteBurger(c echo.Context) error {
	id := c.Param("id")
	db := database.GetDB()

	// Verificar se o burger existe antes de tentar desvincular os opcionais
	var burger models.Burger
	if err := db.First(&burger, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Burger não encontrado"})
	}

	// Tentar desvincular os opcionais do burger
	if err := db.Model(&burger).Association("Opcionais").Clear(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao desvincular opcionais: " + err.Error()})
	}

	// Agora, excluir o burger
	if err := db.Delete(&models.Burger{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao deletar burger: " + err.Error()})
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
