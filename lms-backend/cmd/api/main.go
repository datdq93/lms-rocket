package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/lms-rocket/lms-backend/internal/config"
	"github.com/lms-rocket/lms-backend/internal/handler"
	"github.com/lms-rocket/lms-backend/internal/middleware"
	"github.com/lms-rocket/lms-backend/internal/repository"
	"github.com/lms-rocket/lms-backend/internal/service"
)

func main() {
	// Load .env file if exists
	_ = godotenv.Load()

	// Initialize logger
	logger, err := zap.NewProduction()
	if os.Getenv("APP_ENV") == "development" {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()

	// Initialize database
	db, err := config.InitDatabase()
	if err != nil {
		sugar.Fatalw("Failed to initialize database", "error", err)
	}

	// Initialize Redis (optional)
	redisClient := config.InitRedis()
	if redisClient != nil {
		sugar.Info("Redis connected successfully")
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	courseRepo := repository.NewCourseRepository(db)

	// Initialize services
	authService := service.NewAuthService(userRepo, redisClient)
	userService := service.NewUserService(userRepo)
	courseService := service.NewCourseService(courseRepo)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	courseHandler := handler.NewCourseHandler(courseService)

	// Setup Gin router
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.LoggerMiddleware(logger))
	router.Use(middleware.ErrorHandler())

	// CORS configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = getAllowedOrigins()
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Request-ID"}
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"time":   time.Now().UTC(),
		})
	})

	// API v1 routes
	api := router.Group("/api/v1")
	{
		// Public routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)
			auth.POST("/forgot-password", authHandler.ForgotPassword)
			auth.POST("/reset-password", authHandler.ResetPassword)
			auth.POST("/verify-email", authHandler.VerifyEmail)
		}

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// Auth
			protected.POST("/auth/logout", authHandler.Logout)
			protected.POST("/auth/resend-verification", authHandler.ResendVerification)

			// Users
			protected.GET("/users/me", userHandler.GetProfile)
			protected.PATCH("/users/me", userHandler.UpdateProfile)
			protected.POST("/users/me/change-password", userHandler.ChangePassword)
			protected.POST("/users/me/avatar", userHandler.UploadAvatar)

			// Courses
			protected.GET("/courses", courseHandler.ListCourses)
			protected.GET("/courses/:slug", courseHandler.GetCourse)
		}

		// Admin routes
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware())
		admin.Use(middleware.RoleMiddleware("admin"))
		{
			admin.GET("/users", userHandler.ListUsers)
			admin.GET("/users/:id", userHandler.GetUser)
			admin.PATCH("/users/:id", userHandler.UpdateUser)
			admin.DELETE("/users/:id", userHandler.DeleteUser)
		}

		// Teacher routes
		teacher := api.Group("/teacher")
		teacher.Use(middleware.AuthMiddleware())
		teacher.Use(middleware.RoleMiddleware("teacher", "admin"))
		{
			teacher.POST("/courses", courseHandler.CreateCourse)
			teacher.PATCH("/courses/:id", courseHandler.UpdateCourse)
			teacher.DELETE("/courses/:id", courseHandler.DeleteCourse)
			teacher.POST("/courses/:id/publish", courseHandler.PublishCourse)
			teacher.POST("/courses/:id/unpublish", courseHandler.UnpublishCourse)
		}
	}

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		sugar.Infow("Starting server", "port", port, "env", os.Getenv("APP_ENV"))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sugar.Fatalw("Failed to start server", "error", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	sugar.Info("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		sugar.Fatalw("Server forced to shutdown", "error", err)
	}

	// Close database connection
	sqlDB, _ := db.DB()
	sqlDB.Close()

	if redisClient != nil {
		redisClient.Close()
	}

	sugar.Info("Server exited gracefully")
}

func getAllowedOrigins() []string {
	origins := []string{
		"https://lmsrocket.com",
		"https://www.lmsrocket.com",
		"https://app.lmsrocket.com",
	}

	if os.Getenv("APP_ENV") == "development" {
		origins = append(origins, "http://localhost:3000")
	}

	return origins
}
