package group

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"paser/internal/utils/checkFolder"
	"strings"
)

type OldSitemap struct {
	Urls []string `json:"urls"`
}

func Group(name string, pathFormated string, pathGroup string) {
	const op = "parser.group.Group"

	data, err := ioutil.ReadFile(pathFormated + name)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err, op)
		return
	}

	var oldSitemap OldSitemap
	err = json.Unmarshal(data, &oldSitemap)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err, op)
		return
	}

	//TODO описать внятную типизацию
	comicsMap := make(map[string][]string)
	for _, url := range oldSitemap.Urls {
		parts := strings.Split(url, "/")
		comicName := parts[4] //
		comicsMap[comicName] = append(comicsMap[comicName], url)
	}

	newJSONData, err := json.MarshalIndent(comicsMap, "", "    ")
	if err != nil {
		fmt.Println("Ошибка при преобразовании в JSON:", err)
		return
	}

	checkFolder.Check(pathGroup)

	err = ioutil.WriteFile(pathGroup+name, newJSONData, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	fmt.Println("Данные успешно записаны в файл:", pathGroup, name)
}
