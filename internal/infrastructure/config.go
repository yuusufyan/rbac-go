package infrastructure

import (
	"fmt"
	"os"
	"rbac-go/internal/infrastructure/utils"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string `validate:"required"`

	DBHost     string `validate:"required"`
	DBPort     string `validate:"required"`
	DBUser     string `validate:"required"`
	DBPassword string `validate:"required"`
	DBName     string `validate:"required"`

	RedisHost     string
	RedisPort     string
	RedisPassword string

	MinioAccessKey string
	MinioSecretKey string
	MinioUrl       string
	MinioBucket    string

	CORSAllowed []string `validate:"required,dive,required"`
	UserService string
	AppSecret   string
	VerifyTls   bool
}

var AppConfig *Config

func LoadConfig() {
	utils.NewLogger().Info("Loading config")

	AppConfig = &Config{
		AppPort:    getEnv("APP_PORT", "11000"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "postgres"),

		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),

		MinioAccessKey: getEnv("MINIO_ACCESS_KEY", ""),
		MinioSecretKey: getEnv("MINIO_SECRET_KEY", ""),
		MinioUrl:       getEnv("MINIO_URL", "http://localhost:9000"),
		MinioBucket:    getEnv("MINIO_BUCKET", "minio"),

		CORSAllowed: strings.Split(getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:5173"), ","),
		UserService: getEnv("USER_SERVICE", "http://localhost:13000"),
		AppSecret:   getEnv("APP_SECRET", ""),
		VerifyTls:   getEnv("VERIFY_TLS", "true") == "true",
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	log := utils.NewLogger()
	log.Info(
		fmt.Sprintf(
			"WARNING: Environment variable %s is not set. Using fallback value: %s",
			key,
			fallback,
		),
	)

	return fallback
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		utils.NewLogger().Error("failed to load .env file", err, false)
	}
}
