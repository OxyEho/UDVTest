package db

import (
	"UDVTest/entities"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type IRepo interface {
	GetBooks(pageNum, size int) ([]entities.Book, error)
	GetBook(id uint) (entities.Book, error)
	GetBookItems(id uint) (entities.BookItems, error)
}

type PostgresRepo struct {
	db *gorm.DB
}

func (p *PostgresRepo) GetBooks(limit, offset int) ([]entities.Book, error) {
	var books []entities.Book
	err := p.db.Preload("Publisher").Limit(limit).Offset(offset).Find(&books).Error
	return books, err
}

func (p *PostgresRepo) GetBook(id uint) (entities.Book, error) {
	var book entities.Book
	err := p.db.Preload("Publisher").First(&book, id).Error
	return book, err
}

func (p *PostgresRepo) GetBookItems(id uint) (entities.BookItems, error) {
	var count int64
	p.db.Model(&entities.BookEntity{}).Where("book_id = ?", id).Count(&count)
	var takers []entities.User
	err := p.db.Model(&entities.BookEntity{}).
		Select(`"users"."id", "users"."name", "users"."surname", "users"."patronymic"`).
		Joins(`join "users" on "book_entities"."taker_id" = "users"."id"`).
		Where("book_id = ?", id).
		Scan(&takers).Error
	bookEntities := entities.BookItems{
		ID:         id,
		ItemsCount: uint(count),
		Takers:     takers,
	}
	return bookEntities, err
}

func NewPostgresRepo() PostgresRepo {
	return PostgresRepo{db: getDb()}
}

func InitDb() {
	db := getDb()
	err := db.AutoMigrate(&entities.User{}, &entities.Publisher{}, &entities.Book{}, &entities.BookEntity{})
	if err != nil {
		log.Fatal(err)
	}
}

func getDb() *gorm.DB {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	passw := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		host,
		user,
		passw,
		dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
