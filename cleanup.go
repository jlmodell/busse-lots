package main

import (
	"bufio"
	"os"
	"strings"
)

func clean_lot_csv(filePath *string) string {
	src, err := os.OpenFile(*filePath, os.O_RDONLY, 0666)
	check(err)
	defer src.Close()

	dst, err := os.OpenFile(*filePath+".new", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666)
	check(err)

	dst.Seek(0, 0)

	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "|P", "", -1)
		line = strings.Replace(line, "\"", "", -1)

		dst.WriteString(line + "\n")
	}

	dst.Close()

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return *filePath + ".new"
}
