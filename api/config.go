package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func envACCOUNTSID() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file")
		log.Fatal(err)
	}
	return os.Getenv("TWILIO_ACCOUNT_SID")
}
func envAUTHTOKEN() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
		log.Fatal(err)
	}
	return os.Getenv("TWILIO_AUTH_TOKEN")
}
func envSERVICESID() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
		log.Fatal(err)
	}
	return os.Getenv("TWILIO_SERVICES_ID")
}
