package pages

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"paser/internal/utils/checkFolder"
	"strconv"
	"strings"
)

type ComicData struct {
	ComicName string   `json:"comic_name"`
	Pages     []string `json:"pages"`
}

type Comics map[string][]string

func Get(url string, nameFile string, groupPath string) {
	const op = "parser.pages.Get"

	data, err := ioutil.ReadFile(groupPath + nameFile)
	if err != nil {
		fmt.Println("Ошибка при чтении файла: \n", nameFile, err, op)
		return
	}

	var comics Comics
	err = json.Unmarshal(data, &comics)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON: \n", err)
		return
	}

	checkFolder.Check("../../result/comics")

	for name, urls := range comics {
		_ = urls
		parseUrl := url + "/comic/" + name
		dirName := "../../result/comics/" + name
		checkFolder.Check(dirName)

		GetPreview(parseUrl, dirName)
		CheckChildren(parseUrl, dirName)

		for _, url := range urls {
			tom := strings.Split(url, "/")
			tomName := tom[len(tom)-1]
			_, err := strconv.Atoi(tomName)
			if err != nil {
				tomDirName := fmt.Sprintf("%s/%s", dirName, tomName)
				CheckChildren(url, tomDirName)
			}
		}
	}

}

func CheckChildren(url string, saveDir string) {
	const op = "parser.pages.CheckChildren"

	fmt.Printf("Начался поиск глав для:%s ...ожидайте \n", url)

	for i := 1; i > 0; i++ {
		parseUrl := fmt.Sprintf("%s/%d", url, i)
		newSaveDir := fmt.Sprintf("%s/%d", saveDir, i)

		resp, err := http.Get(parseUrl)
		if err != nil || resp.StatusCode != http.StatusOK {
			GetComicsTom(url, saveDir)
			break
		}

		CheckChildren(parseUrl, newSaveDir)
	}

	fmt.Printf("Поиск глав для:%s закончился \n", url)
}

func GetComicsTom(url string, dirSave string) {
	const op = "parser.pages.GetComicsTom"

	for i := 1; i > 0; i++ {
		parseUrl := fmt.Sprintf("%s/%d", url, i)
		resp, err := http.Get(parseUrl)
		if err != nil || resp.StatusCode != http.StatusOK {
			break
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Printf("Ошибка получения дом дерева:", err, op)
			break
		}

		checkFolder.Check(dirSave)

		el := doc.Find(".img-responsive.scan-page")
		imgName := fmt.Sprintf("%s/%d.jpg", dirSave, i)
		GetSrcImgMainContent(el, imgName)
	}
}

func GetPreview(url string, saveDir string) {
	const op = "parser.pages.GetPreview"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	el := doc.Find(".img-responsive")
	GetSrcImg(el, saveDir)
}

func GetSrcImg(tagImg *goquery.Selection, saveDir string) {
	const op = "parser.pages.GetSrcImg"

	fmt.Printf("Начался поиск src \n")

	tagImg.Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists {
			imgName := saveDir + "/" + "0.jpg"
			imgSrc := "http:" + src
			err := SaveImg(imgName, imgSrc)
			if err != nil {
				fmt.Printf("Не удалось сохранить картинку http:%s \n", src)
			}
		}
	})

	fmt.Printf("Поиск src закончился \n")
}

func GetSrcImgMainContent(tagImg *goquery.Selection, saveDir string) {
	const op = "parser.pages.GetSrcImgMainContent"

	fmt.Printf("Начался поиск src")

	tagImg.Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		src = strings.TrimSpace(src)
		if exists {
			err := SaveImg(saveDir, src)
			if err != nil {
				fmt.Printf("Не удалось сохранить картинку:%s \n", src)
			}
		}
	})

	fmt.Printf("Поиск src закончился \n")
}

func SaveImg(savePath string, imageUrl string) error {
	const op = "parse.page.SaveImg"

	file, err := os.Create(savePath)

	resp, err := http.Get(imageUrl)
	if err != nil {
		fmt.Println("Ошибка при получении изображения:", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка: неверный статус код %d %s\n", resp.StatusCode, op)
		return err
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err, op)
		return err
	}

	return nil
}
