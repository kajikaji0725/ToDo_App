package db

import (
	"fmt"
	"log"

	"github.com/kajikaji0725/ToDo-App/pkg/db/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Contoroller struct {
	contoroller *gorm.DB
}

type Config struct {
	Host     string
	Username string
	Password string
	DBname   string
	Port     string
}

func Dsn(config Config) string {
	return fmt.Sprintf(
		"user=%s password=%s port=%s database=%s host=%s sslmode=disable",
		config.Username,
		config.Password,
		config.Port,
		config.DBname,
		config.Host,
	)
}

func NewContoroller() (*Contoroller, error) {
	config := Config{
		Host:     "homework-db",
		Username: "root",
		Password: "root",
		DBname:   "root",
		Port:     "5432",
	}

	db, err := gorm.Open(postgres.Open(Dsn(config)), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&model.Homework{},
	)
	if err != nil {
		return nil, err
	}
	return &Contoroller{db}, nil
}

func (contoroller *Contoroller) SetHomework(toDo model.ToDo) error {
	log.Println("111111111111111111111111111")
	homework := model.Homework{}
	homework.ID = 1
	homework.Homework.Id = toDo.Id
	homework.Homework.Subject = toDo.Subject
	homework.Homework.Date = toDo.Date
	log.Println("111111111111111111111111111")
	contoroller.contoroller.Create(&homework)
	if err := contoroller.contoroller.Create(&homework).Error; err != nil {
		log.Print(err)
	}
	return nil
}
