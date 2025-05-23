package env

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Guizzs26/currency-converter-go/internal/config"
	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ error loading .env file")
	} else {
		log.Println("✔️ .env file loaded")
	}
}

func InitConfig() config.Config {
	loadEnv()

	addr := mustGetString("ADDR", ":3333")
	env := mustGetString("ENV", "development")
	connStr := mustGetString("DB_ADDR", "postgres://pguser:pgpass@localhost:5455/currency_converter?sslmode=disable")
	maxOpenConns := mustGetInt("DB_MAX_OPEN_CONNS", 10)
	maxIdleConns := mustGetInt("DB_MAX_IDLE_CONNS", 5)
	maxIdleTime := mustGetString("DB_MAX_IDLE_TIME", "15m")

	db := config.DBConfig{
		ConnStr:      connStr,
		MaxOpenConns: maxOpenConns,
		MaxIdleConns: maxIdleConns,
		MaxIdleTime:  maxIdleTime,
	}
	cfg := config.Config{
		Addr: addr,
		Env:  env,
		DB:   db,
	}

	return cfg
}

func getString(key, fallback string) (string, error) {
	val, ok := os.LookupEnv(key)

	if !ok || strings.TrimSpace(val) == "" {
		if strings.TrimSpace(fallback) == "" {
			return "", fmt.Errorf("missing or empty environment variable '%s' and fallback value is also empty", key)
		}
		return fallback, nil
	}

	return val, nil
}

func getInt(key string, fallback int) (int, error) {
	val, ok := os.LookupEnv(key)

	if !ok || strings.TrimSpace(val) == "" {
		return fallback, nil
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback, fmt.Errorf("invalid integer value for '%s': %s", key, val)
	}

	return valAsInt, nil
}

func mustGetString(key, fallback string) string {
	val, err := getString(key, fallback)
	if err != nil {
		log.Fatalf("❌ %v", err)
	}

	return val
}

func mustGetInt(key string, fallback int) int {
	val, err := getInt(key, fallback)
	if err != nil {
		log.Fatalf("❌ %v", err)
	}

	return val
}
