package service

import (
	"GoEcho/models"

	"gorm.io/gorm"
)

// Função para definir o status padrão do burger
func SetDefaultStatus(burger *models.Burger, db *gorm.DB) error {
	if burger.StatusID == 0 {
		// Buscar o status 'Solicitado' como padrão
		var status models.Status
		if err := db.Where("tipo = ?", "Solicitado").First(&status).Error; err != nil {
			return err
		}
		burger.StatusID = status.ID
	}
	return nil
}
