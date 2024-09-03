package main

import (
	adapters_http "auth-server/internal/adapters/incoming/http"
	"auth-server/internal/adapters/incoming/http/server"
	"auth-server/internal/adapters/outgoing/jwt"
	database "auth-server/internal/adapters/outgoing/repository"
	"auth-server/internal/application/services"
	"auth-server/internal/configs"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Println("Diretório atual:", dir)

	config, err := configs.LoadConfig(dir)
	if err != nil {
		rootDir := filepath.Join(dir, "..", "..")
		config, err = configs.LoadConfig(rootDir)
		if err != nil {
			fmt.Println("Erro ao carregar configurações:", err)
			panic(err)
		}
	}
	fmt.Printf("Configurações carregadas: %+v\n", config)

	db, err := sql.Open(config.DBDriver, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepository := database.NewUserRepository(db)
	tokenGen := jwt.NewTokenGenerator(config.JWTSecret)
	authService := services.NewAuthService(userRepository, tokenGen)
	authHandler := adapters_http.NewAuthHandler(authService)

	webServer := server.NewWebServer(":" + config.WebServerPort)

	webServer.AddHandler(http.MethodPost, "/login", authHandler.Login)
	webServer.AddHandler(http.MethodPost, "/register", authHandler.Register)

	fmt.Println("Starting web server on port", config.WebServerPort)
	go func() {
		err = webServer.Start()
		if err != nil {
			panic(err)
		}
	}()
	select {}
}
