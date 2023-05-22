package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func File(args []string) {
	var twc, tcc, tlc, fc int
	for _, fileName := range args {
		var lc, wc, cc int
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}
		scan := bufio.NewScanner(file)
		for scan.Scan() {
			s := scan.Text()
			wc += len(strings.Fields(s))
			twc += wc
			cc += len(s)
			tcc += cc
			lc++
			tlc += lc
		}
		var lines string
		if lc > 1 {
			lines = "lines"
		} else {
			lines = "line"
		}
		fmt.Printf("there are %v words, %v characters and %v %v in %v\n", wc, cc, lc, lines, fileName)
		fc++
		file.Close()
	}
	if fc > 1 {
		fmt.Printf("total: %v words, %v characters and %v lines\n ", twc, tcc, tlc)
	}

}
