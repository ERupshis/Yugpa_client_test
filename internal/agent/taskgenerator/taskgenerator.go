package taskgenerator

import (
	"fmt"
	"math/rand"
	"os"
)

func GenerateRandomPath(path string) string {
	res := path

	for i := 0; i < rand.Intn(3)+1; i++ {
		files, err := os.ReadDir(res)
		if err != nil {
			fmt.Println("Error:", err)
		}

		if len(files) == 0 {
			break
		}

		randomIndex := rand.Intn(len(files))
		selected := files[randomIndex]

		if selected.IsDir() {
			res += selected.Name() + "\\"
			subFiles, err := os.ReadDir(res)
			if err != nil {
				fmt.Println("Error:", err)
			}
			files = subFiles
		} else {
			break
		}
	}

	return res
}
