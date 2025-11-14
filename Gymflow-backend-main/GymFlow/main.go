package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gymflow/auth"
	"gymflow/config"
	"gymflow/handler"
	"gymflow/models"
	"gymflow/middleware"
	"gymflow/repository"
	"gymflow/service"
)

func main() {
	cfg := config.Load()

	// --- DB Init ---
	db, err := gorm.Open(postgres.Open(cfg.DB_DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	// ---- MIGRATIONS ----
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	// --- Dependency Injection ---
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	jwtService := auth.NewJWTService(cfg.JWTSecret, cfg.JWTExpireMins)
	userHandler := handler.NewUserHandler(userService, jwtService)

	r := gin.Default()

	// ------------------------------
	//       PUBLIC ROUTES
	// ------------------------------
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", userHandler.Register)
		authGroup.POST("/login", userHandler.Login)
		authGroup.GET("/verify", userHandler.VerifyEmail) // подтверждение email
	}

	// ------------------------------
	//      PROTECTED /api routes
	// ------------------------------
	api := r.Group("/api")
	api.Use(auth.JWTAuthMiddleware(jwtService))
	{
		api.GET("/me", userHandler.Me)
	}

	// ------------------------------
	//         ADMIN ROUTES
	// ------------------------------
	admin := r.Group("/admin")
	admin.Use(
		auth.JWTAuthMiddleware(jwtService),
		middleware.AdminOnlyMiddleware(), // проверка роли
	)
	{
		admin.POST("/trainers", userHandler.RegisterTrainer) // создать тренера
		admin.GET("/users", userHandler.GetAllUsers)        // пример
	}

	// ------------------------------
	//           RUN SERVER
	// ------------------------------
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Server running on %s", addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
