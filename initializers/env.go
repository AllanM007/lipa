package initializers

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnvVariables()  {

	// err := godotenv.Load(".env")
	err := godotenv.Load(filepath.Join("./", ".env"))
	
	// envPath,_ := filepath.Abs("./.env");
	// environments,err := godotenv.Read(envPath)
	// fmt.Println(environments)
	if err != nil {
		fmt.Println("error loading env file")
		log.Fatal(err)
	}
}