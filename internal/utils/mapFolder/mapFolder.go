package mapFolder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"paser/internal/utils/checkFolder"
	"path/filepath"
)

func MakeMapComics() {
	comicsMapDir := "../../result/comics"
	checkFolder.Check(comicsMapDir)
	comicsDir := "../../result/comics"
	fileMap := make(map[string]interface{})
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Добавляем информацию о файле или папке в карту
		fileMap[path] = info.Name()

		return nil
	}

	err := filepath.Walk(comicsDir, walkFn)
	if err != nil {
		fmt.Println("Ошибка при обходе директории:", err)
		return
	}

	// Преобразуем карту в JSON
	jsonData, err := json.MarshalIndent(fileMap, "", "    ")
	if err != nil {
		fmt.Println("Ошибка при преобразовании в JSON:", err)
		return
	}

	err = ioutil.WriteFile(comicsMapDir+"/map.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
	}

	//group.Group("/_map.json", comicsMapDir, comicsMapDir)
}
