package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func File(args []string) {
	var totalwc, totalcc, totallc, filec int
	for _, fileName := range args {
		var linec, wordc, characterc int
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}
		scan := bufio.NewScanner(file)
		for scan.Scan() {
			s := scan.Text()
			wordc += len(strings.Fields(s))
			totalwc += wordc
			characterc += len(s)
			totalcc += characterc
			linec++
			totallc += linec
		}
		var lines string
		if linec > 1 {
			lines = "lines"
		} else {
			lines = "line"
		}
		fmt.Printf("there are %v words, %v characters and %v %v in %v\n", wordc, characterc, linec, lines, fileName)
		filec++
		file.Close()
	}
	if filec > 1 {
		fmt.Printf("total: %v words, %v characters and %v lines\n ", totalwc, totalcc, totallc)
	}

}
