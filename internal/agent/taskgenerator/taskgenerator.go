package taskgenerator

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func GenerateRandomPath() string {
	res := "C:\\Users\\"

	for i := 0; i < rand.Intn(3)+1; i++ {
		files, err := os.ReadDir(res)
		if err != nil {
			fmt.Println("Error:", err)
		}

		// Randomly select a file or directory from the list
		if len(files) == 0 {
			break
		}

		randomIndex := rand.Intn(len(files))
		selected := files[randomIndex]

		// If selected element is a directory, list its contents for the next iteration
		if selected.IsDir() {
			res += selected.Name() + "\\"
			subFiles, err := os.ReadDir(res)
			if err != nil {
				fmt.Println("Error:", err)
				res = res[:strings.LastIndex(res, "\\")]
			}
			files = subFiles
		} else {
			break
		}
	}

	return res
}
