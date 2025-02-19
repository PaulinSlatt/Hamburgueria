package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Tipo customizado para lidar com JSONB
type StringArray []string

func (s StringArray) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("tipo inválido para conversão de JSONB")
	}
	return json.Unmarshal(bytes, s)
}

type Pao struct {
	ID   uint   `json:"id"`
	Tipo string `json:"tipo"`
}

type Carne struct {
	ID   uint   `json:"id"`
	Tipo string `json:"tipo"`
}

type Opcional struct {
	ID   uint   `json:"id"`
	Tipo string `json:"tipo"`
}

type Ingrediente struct {
	ID         uint       `json:"id"`
	PaoID      uint       `json:"pao_id"`
	PaoNome    string     `json:"pao_nome"`
	CarneID    uint       `json:"carne_id"`
	CarneNome  string     `json:"carne_nome"`
	Opicionais []Opcional `gorm:"many2many:ingredientes_opcionais;"` // associação many-to-many
}

// Status de um Burger

type Status struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Tipo string `json:"tipo" gorm:"unique;not null"`
}

// Burger com relacionamentos
type Burger struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Nome string `json:"nome"`

	CarneID   uint   `json:"carne_id"`
	CarneNome string `json:"carne_nome"`

	PaoID   uint   `json:"pao_id"`
	PaoNome string `json:"pao_nome"`

	Opcionais  []Opcional `json:"opcionais" gorm:"many2many:burger_opcionais;"` // Agora visível no JSON
	StatusID   uint       `json:"status_id"`
	StatusNome string     `json:"status_nome"`
}

type BurgerOpcional struct {
	BurgerID   uint `json:"burger_id"`
	OpcionalID uint `json:"opcional_id"`
}
