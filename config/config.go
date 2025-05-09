package config

import (
	"fmt"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadEnvs loads all .env.* files from given env (dev, prod, test)
func LoadEnvs(env string) {
	envDir := filepath.Join("env", env)

	files, err := filepath.Glob(filepath.Join(envDir, ".env.*"))
	if err != nil {
		fmt.Println("Error reading env files:", err)
		return
	}

	for _, file := range files {
		err := godotenv.Overload(file)
		if err != nil {
			fmt.Printf("Failed to load %s: %v\n", file, err)
		} else {
			fmt.Printf("Loaded env: %s\n", file)
		}
	}
}
