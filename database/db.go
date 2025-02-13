package database

import (
	"GoEcho/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	var err error
	connStr := "host=localhost user=root password=root dbname=root port=5433 sslmode=disable"
	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	fmt.Println("Conectado ao banco de dados!")
	SeedDatabase(db)

	err = db.AutoMigrate(&models.Ingrediente{}, &models.Burger{}, &models.Opcional{}, &models.Status{})

	if err != nil {
		log.Fatal("Erro ao migrar o banco de dados:", err)
	}

	log.Println("Banco de dados conectado com sucesso!")
}

func GetDB() *gorm.DB {
	return db
}

func SeedDatabase(db *gorm.DB) {
	// Inserir pães
	paes := []models.Pao{
		{Tipo: "Integral"},
		{Tipo: "Italiano Branco"},
		{Tipo: "3 Queijos"},
		{Tipo: "Parmesão e Orégano"},
	}
	for _, p := range paes {
		err := db.FirstOrCreate(&p, models.Pao{Tipo: p.Tipo}).Error
		if err != nil {
			log.Printf("Erro ao inserir pão: %v", err)
		} else {
			log.Printf("Pão '%s' inserido ou já existente", p.Tipo)
		}
	}

	// Inserir carnes
	carnes := []models.Carne{
		{Tipo: "Maminha"},
		{Tipo: "Alcatra"},
		{Tipo: "Picanha"},
		{Tipo: "Veggie Burger"},
	}
	for _, c := range carnes {
		err := db.FirstOrCreate(&c, models.Carne{Tipo: c.Tipo}).Error
		if err != nil {
			log.Printf("Erro ao inserir carne: %v", err)
		} else {
			log.Printf("Carne '%s' inserida ou já existente", c.Tipo)
		}
	}

	// Inserir opcionais
	opcionais := []models.Opcional{
		{Tipo: "Bacon"},
		{Tipo: "Salame"},
		{Tipo: "Cebola Roxa"},
		{Tipo: "Cheddar"},
		{Tipo: "Tomate"},
		{Tipo: "Pepino"},
	}
	for _, o := range opcionais {
		err := db.FirstOrCreate(&o, models.Opcional{Tipo: o.Tipo}).Error
		if err != nil {
			log.Printf("Erro ao inserir opcional: %v", err)
		} else {
			log.Printf("Opcional '%s' inserido ou já existente", o.Tipo)
		}
	}

	// Inserir ingredientes

	ingredientes := []models.Ingrediente{
		{
			Pao:        &paes[0],                        // Referência para o Pão
			Carne:      &carnes[0],                      // Referência para a Carne
			Opicionais: []models.Opcional{opcionais[0]}, // Associando os opcionais ao ingrediente
		},
	}

	for _, i := range ingredientes {
		// Primeiro, insira os opcionais, garantindo que já existem
		for _, opcional := range i.Opicionais {
			err := db.FirstOrCreate(&opcional).Error
			if err != nil {
				log.Printf("Erro ao inserir opcional: %v", err)
			}
		}

		// Agora, associe o ingrediente com os opcionais, carne e pão
		err := db.FirstOrCreate(&i).Error
		if err != nil {
			log.Printf("Erro ao inserir ingrediente: %v", err)
		} else {
			log.Println("Ingrediente inserido ou já existente")
		}
	}

	// Inserir status
	status := []models.Status{
		{Tipo: "Solicitado"},
		{Tipo: "Em produção"},
		{Tipo: "Finalizado"},
	}
	for _, s := range status {
		err := db.FirstOrCreate(&s, models.Status{Tipo: s.Tipo}).Error
		if err != nil {
			log.Printf("Erro ao inserir status '%s': %v", s.Tipo, err)
		} else {
			log.Printf("Status '%s' inserido ou já existente", s.Tipo)
		}
	}
}
