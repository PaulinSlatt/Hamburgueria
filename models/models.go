package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Criar um tipo customizado para lidar com JSONB no GORM
type StringArray []string

// Converter para JSON ao salvar no banco
func (s StringArray) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Converter JSON para o struct ao recuperar do banco
func (s *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("tipo inválido para conversão de JSONB")
	}
	return json.Unmarshal(bytes, s)
}

const (
	CarneMaminha = "Maminha"
	CarneAlcatra = "Alcatra"
	CarnePicanha = "Picanha"
	CarneVeggie  = "Veggie Burger"
)

const (
	PaoIntegral         = "Integral"
	PaoItalianoBranco   = "Italiano Branco"
	PaoTresQueijos      = "3 Queijos"
	PaoParmesaoEOregano = "Parmesão e Oregano"
)

const (
	OpcionalBacon      = "Bacon"
	OpcionalSalame     = "Salame"
	OpcionalCebolaRoxa = "Cebola Roxa"
	OpcionalCheddar    = "Cheddar"
	OpcionalTomate     = "Tomate"
	OpcionalPepino     = "Pepino"
)

type Burger struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	Nome      string      `json:"nome"`
	Carne     string      `json:"carne"`
	Pao       string      `json:"pao"`
	Opcionais StringArray `json:"opcionais" gorm:"type:jsonb"`
	Status    string      `json:"status"`
}

type Ingrediente struct {
	ID   uint   `json:"id"`
	Nome string `json:"nome"`
}

type Pao struct {
	ID   uint   `json:"id"`
	Nome string `json:"nome"`
}

type Carne struct {
	ID   uint   `json:"id"`
	Nome string `json:"nome"`
}

type Status struct {
	ID   uint   `json:"id"`
	Nome string `json:"nome"`
}
