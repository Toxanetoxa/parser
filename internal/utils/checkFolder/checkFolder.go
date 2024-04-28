package checkFolder

import (
	"fmt"
	"os"
)

func Check(path string) {
	const op = "checkFolder.checkFolder.Check"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			fmt.Println("Ошибка при создании директории:", err, op)
			return
		}
		fmt.Println("Директория создана:", path)
	}
}
