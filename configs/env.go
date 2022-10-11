package configs

import(
	"log"
    "os"
    "github.com/joho/godotenv"
)

func EnvMongoURI() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("бля не нашел переменные окружения")
    }
  
    return os.Getenv("MONGOURI")
}