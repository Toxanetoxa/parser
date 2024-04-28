package sitemap

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"paser/internal/utils/checkFolder"
	"time"
)

type URL struct {
	Loc string `xml:"loc"`
}

type SitemapIndex struct {
	Urls []URL `xml:"url"`
}

func New(url string, pathParse string) (string, error) {
	const op = "parser.sitemap.New"

	resp, err := http.Get(url + "/sitemap.xml")
	if err != nil {
		fmt.Println("Ошибка при загрузке sitemap:", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка:", resp.Status, op)
		return "", err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err, op)
		return "", err
	}

	// Парсинг XML
	var sitemap SitemapIndex
	err = xml.Unmarshal(data, &sitemap)
	if err != nil {
		fmt.Println("Ошибка при разборе XML:", err, op)
		return "", err
	}

	// Преобразование в JSON
	jsonData, err := json.MarshalIndent(sitemap, "", "    ")
	if err != nil {
		fmt.Println("Ошибка при преобразовании в JSON:", err, op)
		return "", err
	}

	// Запись в файл
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02-15:04:05")

	nameFile := formattedTime + "-sitemap.json"

	checkFolder.Check(pathParse)

	err = ioutil.WriteFile(pathParse+nameFile, jsonData, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return "", err
	}

	fmt.Println("Данные успешно записаны в файл:", pathParse, nameFile)

	return nameFile, nil
}
