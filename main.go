package main

import (
	"Hotel_BE/component"
	"Hotel_BE/docs"
	"Hotel_BE/modules/users"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	if err := migrateDatabase(db); err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func migrateDatabase(db *gorm.DB) error {
	models := []interface{}{
		&users.User{},
		// Add more models as needed
	}

	for _, model := range models {
		fmt.Println("Migrating model: ", model)
		if err := db.AutoMigrate(model); err != nil {
			return err
		}
	}

	return nil
}

func runService(db *gorm.DB) error {
	appCtx := component.NewAppContext(db)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	v1 := r.Group("/v1")

	user := v1.Group("/users")
	{
		userController := users.NewUserController()
		user.POST("", userController.CreateUser(appCtx))
		user.GET("", userController.ListUser(appCtx))
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r.Run(":8080")
}
