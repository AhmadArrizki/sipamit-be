package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path"
	"strconv"
)

var App struct {
	Debug            bool   `mapstructure:"APP_DEBUG"`
	Host             string `mapstructure:"APP_HOST"`
	Port             string `mapstructure:"APP_PORT"`
	CORSAllowOrigins string `mapstructure:"CORS_ALLOW_ORIGINS"`
	SwaggerHost      string `mapstructure:"SWAGGER_HOST"`
}

var JWT struct {
	Key    string `mapstructure:"AUTH_JWT_KEY"`
	Expire int    `mapstructure:"AUTH_JWT_EXPIRE"`
}

var Mongo struct {
	Url  string `mapstructure:"MONGODB_URI"`
	Name string `mapstructure:"MONGODB_NAME"`
}

func init() {
	rootPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error while getting root path: %v", err)
	}
	filepath := path.Join(rootPath, ".env")
	err = godotenv.Load(filepath)
	if err != nil {
		fmt.Println(".env file not found, using environment variables instead.")
	}

	App.Debug = os.Getenv("APP_DEBUG") == "true"
	App.Host = os.Getenv("APP_HOST")
	if App.Host == "" {
		panic("APP_HOST is not set")
	}
	App.Port = os.Getenv("APP_PORT")
	if App.Port == "" {
		panic("APP_PORT is not set")
	}
	App.CORSAllowOrigins = os.Getenv("CORS_ALLOW_ORIGINS")
	if App.CORSAllowOrigins == "" {
		panic("CORS_ALLOW_ORIGINS is not set")
	}
	App.SwaggerHost = os.Getenv("SWAGGER_HOST")
	if App.SwaggerHost == "" {
		panic("SWAGGER_HOST is not set")
	}

	JWT.Key = os.Getenv("AUTH_JWT_KEY")
	if JWT.Key == "" {
		panic("AUTH_JWT_KEY is not set")
	}
	jwtExp := os.Getenv("AUTH_JWT_EXPIRE")
	if jwtExp == "" {
		panic("AUTH_JWT_EXPIRE is not set")
	} else {
		JWT.Expire, err = strconv.Atoi(jwtExp)
		if err != nil {
			panic("AUTH_JWT_EXPIRE is not valid")
		}
	}

	Mongo.Url = os.Getenv("MONGODB_URI")
	if Mongo.Url == "" {
		panic("MONGODB_URI is not set")
	}
	Mongo.Name = os.Getenv("MONGODB_NAME")
	if Mongo.Name == "" {
		panic("MONGODB_NAME is not set")
	}
}
