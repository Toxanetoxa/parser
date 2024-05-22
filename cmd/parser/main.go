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
	"paser/internal/utils/checkFolder"
	"paser/internal/utils/mapFolder"
)

func main() {
	err := godotenv.Load(".env.prod")
	if err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
		return
	}
	url := os.Getenv("URL_SITEMAP")

	var command string
	fmt.Printf("\nwrite to start parsing: start\n")
	_, _ = fmt.Scan(&command)

	if command != "start" {
		fmt.Printf("\nERROR! parsing has not started\n")
		return
	}

	pathResult := os.Getenv("PATH_RESULT")
	checkFolder.Check(pathResult)
	pathParse := os.Getenv("PATH_PARSE")
	checkFolder.Check(pathParse)
	pathFormat := os.Getenv("PATH_FORMAT")
	checkFolder.Check(pathFormat)
	pathGroup := os.Getenv("PATH_GROUP")
	checkFolder.Check(pathGroup)

	nameFile, err := sitemap.New(url, pathParse)
	if err != nil {
		return
	}
	formater.Formating(nameFile, pathParse, pathFormat)
	group.Group(nameFile, pathFormat, pathGroup)
	pages.Get(url, nameFile, pathGroup)
	mapFolder.MakeMapComics()

	fmt.Print("<---------DONE!!!! parsing completed successfully!!! --------->")
}
