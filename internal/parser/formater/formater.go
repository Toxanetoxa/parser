package formater

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"paser/internal/utils/checkFolder"
	"sort"
	"strings"
)

type OldSitemap struct {
	Urls []struct {
		Loc string `json:"Loc"`
	} `json:"Urls"`
}

type NewSitemap struct {
	Urls []string `json:"urls"`
}

func Formating(nameFile string, pathParse string, pathFormated string) {
	const op = "parse.formater.Formating"

	data, err := ioutil.ReadFile(pathParse + nameFile)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", nameFile, err, op)
		return
	}

	var oldSitemap OldSitemap
	err = json.Unmarshal(data, &oldSitemap)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return
	}

	// Создание структуры нового формата
	var newSitemap NewSitemap
	for _, url := range oldSitemap.Urls {
		if strings.HasPrefix(url.Loc, "https://readcomicsonline.ru/comic/") {
			newSitemap.Urls = append(newSitemap.Urls, url.Loc)
		}
	}

	sort.Slice(newSitemap.Urls, func(i, j int) bool {
		return len(newSitemap.Urls[i]) < len(newSitemap.Urls[j])
	})

	// Преобразование в JSON нового формата
	newJSONData, err := json.MarshalIndent(newSitemap, "", "    ")
	if err != nil {
		fmt.Println("Ошибка при преобразовании в JSON:", err)
		return
	}

	checkFolder.Check(pathFormated)

	// Запись в файл нового JSON
	err = ioutil.WriteFile(pathFormated+nameFile, newJSONData, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	fmt.Println("Данные успешно записаны в файл:", pathFormated, nameFile)

}
