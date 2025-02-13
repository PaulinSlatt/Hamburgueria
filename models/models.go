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
	Pao        *Pao       `json:"pao,omitempty"`
	CarneID    uint       `json:"carne_id"`
	Carne      *Carne     `json:"carne,omitempty"`
	Opicionais []Opcional `json:"opcionais,omitempty" gorm:"many2many:ingredientes_opcionais;"`
}

// Status de um Burger

type Status struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Tipo string `json:"tipo" gorm:"unique;not null"`
}

// Burger com relacionamentos
type Burger struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Nome      string     `json:"nome"`
	CarneID   uint       `json:"carne_id" gorm:"not null"` // Chave estrangeira oculta no JSON
	Carne     Carne      `json:"carne" gorm:"foreignKey:CarneID"`
	PaoID     uint       `json:"pao_id" gorm:"not null"` // Chave estrangeira oculta no JSON
	Pao       Pao        `json:"pao" gorm:"foreignKey:PaoID"`
	Opcionais []Opcional `json:"opcionais" gorm:"many2many:burger_opcionais;"`
	StatusID  uint       `json:"status_id" gorm:"not null"` // Chave estrangeira oculta no JSON
	Status    Status     `json:"status" gorm:"foreignKey:StatusID"`
}
