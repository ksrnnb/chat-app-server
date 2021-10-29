package gateway

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() *SqlHandler {
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error while connecting db")
	}

	return &SqlHandler{Conn: db}
}

func (h *SqlHandler) Preload(query string, args ...interface{}) (*gorm.DB) {
	return h.Conn.Preload(query, args...)
}

func (h *SqlHandler) Select(query interface{}, args ...interface{}) (*gorm.DB) {
	return h.Conn.Select(query, args...)
}

func (h *SqlHandler) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return h.Conn.Find(dest, conds...)
}

func (h *SqlHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return h.Conn.Where(query, args...)
}

func (h *SqlHandler) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return h.Conn.First(dest, conds...)
}

func (h *SqlHandler) Create(value interface{}) *gorm.DB {
	return h.Conn.Create(value)
}

func (h *SqlHandler) Update(column string, value interface{}) *gorm.DB {
	return h.Conn.Update(column, value)
}

func (h *SqlHandler) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	return h.Conn.Delete(value, conds...)
}

func (h *SqlHandler) Error() error {
	return h.Conn.Error
}

func (h *SqlHandler) Debug() *gorm.DB {
	return h.Conn.Debug()
}
