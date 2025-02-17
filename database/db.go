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

	ingredientes := []models.Ingrediente{
		{
			PaoID:      paes[0].ID,                         // Usar o ID do pão
			CarneID:    carnes[0].ID,                       // Usar o ID da carne
			Opicionais: []models.Opcional{{Tipo: "Bacon"}}, // Associar os opcionais corretamente
		},
	}

	for _, i := range ingredientes {
		// Preencher os campos de nome
		pao := models.Pao{}
		if err := db.First(&pao, i.PaoID).Error; err != nil {
			log.Printf("Erro ao encontrar pão com ID %d: %v", i.PaoID, err)
		} else {package database

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

	ingredientes := []models.Ingrediente{
		{
			PaoID:      paes[0].ID,                         // Usar o ID do pão
			CarneID:    carnes[0].ID,                       // Usar o ID da carne
			Opicionais: []models.Opcional{{Tipo: "Bacon"}}, // Associar os opcionais corretamente
		},
	}

	for _, i := range ingredientes {
		// Preencher os campos de nome
		pao := models.Pao{}
		if err := db.First(&pao, i.PaoID).Error; err != nil {
			log.Printf("Erro ao encontrar pão com ID %d: %v", i.PaoID, err)
		} else {
			i.PaoNome = pao.Tipo
		}

		carne := models.Carne{}
		if err := db.First(&carne, i.CarneID).Error; err != nil {
			log.Printf("Erro ao encontrar carne com ID %d: %v", i.CarneID, err)
		} else {
			i.CarneNome = carne.Tipo
		}

		// Inserir ingrediente (sem o FirstOrCreate, usando Create)
		if err := db.Create(&i).Error; err != nil {
			log.Printf("Erro ao inserir ingrediente: %v", err)
		} else {
			log.Println("Ingrediente inserido com sucesso")
		}

		// Para os opcionais, se houver, associar com o ingrediente (many2many)
		for _, opcional := range i.Opicionais {
			// Inserir ou associar os opcionais ao ingrediente
			err := db.FirstOrCreate(&opcional).Error
			if err != nil {
				log.Printf("Erro ao inserir opcional '%s': %v", opcional.Tipo, err)
			} else {
				// Associar o opcional ao ingrediente
				if err := db.Model(&i).Association("Opicionais").Append(&opcional).Error; err != nil {
					log.Printf("Erro ao associar opcional '%s' ao ingrediente: %v", opcional.Tipo, err)
				} else {
					log.Printf("Opcional '%s' associado ao ingrediente com sucesso", opcional.Tipo)
				}
			}
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
			i.PaoNome = pao.Tipo
		}

		carne := models.Carne{}
		if err := db.First(&carne, i.CarneID).Error; err != nil {
			log.Printf("Erro ao encontrar carne com ID %d: %v", i.CarneID, err)
		} else {
			i.CarneNome = carne.Tipo
		}

		// Inserir ingrediente (sem o FirstOrCreate, usando Create)
		if err := db.Create(&i).Error; err != nil {
			log.Printf("Erro ao inserir ingrediente: %v", err)
		} else {
			log.Println("Ingrediente inserido com sucesso")
		}

		// Para os opcionais, se houver, associar com o ingrediente (many2many)
		for _, opcional := range i.Opicionais {
			// Inserir ou associar os opcionais ao ingrediente
			err := db.FirstOrCreate(&opcional).Error
			if err != nil {
				log.Printf("Erro ao inserir opcional '%s': %v", opcional.Tipo, err)
			} else {
				// Associar o opcional ao ingrediente
				if err := db.Model(&i).Association("Opicionais").Append(&opcional).Error; err != nil {
					log.Printf("Erro ao associar opcional '%s' ao ingrediente: %v", opcional.Tipo, err)
				} else {
					log.Printf("Opcional '%s' associado ao ingrediente com sucesso", opcional.Tipo)
				}
			}
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
