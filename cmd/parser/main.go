package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"paser/internal/parser/formater"
	"paser/internal/parser/group"
	"paser/internal/parser/pages"
	"paser/internal/parser/sitemap"
	"paser/internal/utils/mapFolder"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("%s, Ошибка загрузки файла .env")
	}
	url := os.Getenv("URL_SITEMAP")
	pathParse := os.Getenv("PATH_PARSE")
	pathFormat := os.Getenv("PATH_FORMAT")
	pathGroup := os.Getenv("PATH_GROUP")
	nameFile, err := sitemap.New(url, pathParse)
	if err != nil {
		return
	}
	formater.Formating(nameFile, pathParse, pathFormat)
	group.Group(nameFile, pathFormat, pathGroup)
	pages.Get(url, nameFile, pathGroup)

	mapFolder.MakeMapComics()

	fmt.Print("---------DONE!!!!---------")
}
