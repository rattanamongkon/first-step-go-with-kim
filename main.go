package main

import (
	"example/go-rest-api/conf"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	t_db := conf.Dbconnection()
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Book{})
	handler := newHandler(db)

	r := gin.New()

	r.POST("/login", loginHandler)

	protected := r.Group("/", authorizationMiddleware)

	protected.GET("/books", handler.listBooksHandler)
	protected.POST("/books", handler.createBookHandler)
	protected.DELETE("/books/:id", handler.deleteBookHandler)

	r.Run()
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// var books = []Book{
// 	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
// 	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
// 	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
// }

type Handler struct {
	db *gorm.DB
}

func newHandler(db *gorm.DB) *Handler {
	return &Handler{db}
}

func loginHandler(c *gin.Context) {
	// implement logic login here!!!!

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
	})

	ss, err := token.SignedString([]byte("Test_CPF_Website"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": ss,
	})
}

func (h *Handler) listBooksHandler(c *gin.Context) {
	// s := c.Request.Header.Get("Authorization")
	// token := strings.TrimPrefix(s, "Bearer ")

	// if err := validateToken(token); err != nil {
	// 	c.AbortWithStatus(http.StatusUnauthorized)
	// 	return
	// }

	var books []Book
	if result := h.db.Find(&books); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *Handler) createBookHandler(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := h.db.Create(&book); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// books = append(books, book)
	c.JSON(http.StatusCreated, book)
}

func (h *Handler) deleteBookHandler(c *gin.Context) {
	id := c.Param("id")

	if result := h.db.Delete(&Book{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.Status(http.StatusNoContent)
}

func validateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte("Test_CPF_Website"), nil
	})
	// if token == "" {
	// 	return fmt.Errorf("token should not be empty")
	// } else {
	// 	if token != "ACCESS_TOKEN" {
	// 		return fmt.Errorf("token provided was invalid")
	// 	}
	// }
	return err
}

func authorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	if err := validateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
