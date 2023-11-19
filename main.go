package main

import (
	"Hotel_BE/component"
	"Hotel_BE/docs"
	"Hotel_BE/middlewares"
	"Hotel_BE/modules/rooms"
	"Hotel_BE/modules/users"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	migrateFlag := flag.Bool("migrate", false, "Perform database migration")
	flag.Parse()

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	if *migrateFlag {
		err := migrateDatabase(db)
		if err != nil {
			log.Fatalf("Database migration failed: %v", err)
		}
		fmt.Println("Database migration completed successfully.")
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func migrateDatabase(db *gorm.DB) error {
	models := []interface{}{
		&users.User{},
		&rooms.RoomType{},
		&rooms.Room{},
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
	r.Use(middlewares.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"author":   "Hotel_BE",
			"message":  "Welcome to Hotel_BE",
			"health":   "OK",
			"api-docs": "http://localhost:8080/swagger/index.html",
			"version":  "1.0.0",
			"status":   "running",
		})
	})

	v1 := r.Group("/v1")

	user := v1.Group("/users")
	{
		userController := users.NewUserController()
		user.POST("", userController.CreateUser(appCtx))
		user.GET("", userController.ListUser(appCtx))
	}

	roomType := v1.Group("/room-types")
	{
		roomTypeController := rooms.NewRoomTypeController(appCtx)
		roomType.POST("", roomTypeController.CreateRoomType())
		roomType.PUT("/:id", roomTypeController.UpdateRoomType())
		roomType.GET("", roomTypeController.ListRoomTypes())
	}

	room := v1.Group("/rooms")
	{
		roomController := rooms.NewRoomController(appCtx)
		room.POST("", roomController.CreateRoom())
		room.PUT("/:id", roomController.UpdateRoom())
		room.GET("", roomController.GetRooms())
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r.Run(":8080")
}
