package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Folder struct {
	Dir     string    `json:"dir"`
	Files   []string  `json:"files"`
	Folders []*Folder `json:"folders"`
}

func checkSuffix(files *[]string) bool {

	for _, file := range *files {
		if strings.HasSuffix(file, ".hack") {
			return true
		}
	}

	return false
}

func dfs(folder *Folder, countVirusFiles *int, parentFolderHasVirus bool) {

	if parentFolderHasVirus || checkSuffix(&folder.Files) {
		*countVirusFiles += len(folder.Files)
		parentFolderHasVirus = true
	}

	for _, f := range folder.Folders {
		dfs(f, countVirusFiles, parentFolderHasVirus)
	}

}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	fmt.Fscan(in, &count)

	for i := 0; i < count; i++ {
		var linesCount int
		fmt.Fscan(in, &linesCount)

		jsonData := []byte{}

		for i := 0; i < linesCount+1; i++ {
			line, _ := in.ReadBytes('\n')
			jsonData = append(jsonData, line...)
		}

		var folder Folder
		err := json.Unmarshal([]byte(jsonData), &folder)
		if err != nil {
			fmt.Println(err)
		}

		countVirusFiles := 0

		dfs(&folder, &countVirusFiles, false)

		fmt.Fprintln(out, countVirusFiles)

	}
}
