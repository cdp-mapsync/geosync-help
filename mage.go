//go:build mage
// +build mage

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var (
	categories = map[string]string{
		"5f09a32fb53ddf002e4f3982": "z-tools-app",
		"5fe792a0bde6ea002d31d8d8": "z-tools-app",
		"5fe791c7cd0751004af7191c": "z-tools-app",
		"5fe790fc5216d3004a7c4959": "z-tools-app",
		"5ffafa359c088100f2be1f6c": "z-tools-app",
		"5ffaeeee5557d5008c6d8295": "z-tools-app",
		"61ab6a7f4e6af500298d91d8": "z-tools-app",
		"5f08958fad72c1002d455db7": "qgis",
		"5fdd5d610bc1940017c6fbbe": "publisher-app",
		"60b14560c59996001555722b": "publisher-app",
		"60b145a66e703d0015bfd23d": "publisher-app",
		"60b14641b5e32100153df572": "publisher-app",
		"60b14683043a040015f1a475": "publisher-app",
		"60b14b0473c48300162e9677": "publisher-app",
		"618bd853edb2c50017410203": "publisher-app",
		"5fdd5dc0482fcd0017aac66d": "utility-solution",
		"5fdd5c361abc34001762414d": "am-tools-app",
		"5fdd5ddc543dd400439eabb1": "campus-solution",
		"5fdd5dfe4993a200170b5b00": "e911-solution",
		"5fdd5e2c482fcd0017aac67d": "geosync-app",
		"60b14c5263e25f0015557bb0": "geosync-app",
		"60b14ca928568000152db6ef": "geosync-app",
		"60b14d849c0a580015c62920": "geosync-app",
		"60b14ccf36e76e0015bb5494": "geosync-app",
		"60b14d3257a5960015dd3b4e": "geosync-app",
		"5fe0bd7506a32700177aee39": "go-plus-app",
	}
)

func Import() {
	err := os.RemoveAll("output")
	if err != nil {
		fmt.Println("f1 failed:", err)
		return
	}

	data, err := os.ReadFile("import.csv")
	if err != nil {
		fmt.Println("f6 failed:", err)
		return
	}

	r := csv.NewReader(strings.NewReader(string(data)))
	r.LazyQuotes = true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}

		now := time.Now()

		md := []byte(
			`---
title: ` + record[3] + `
lastmod: ` + now.Format("2006-01-02 15:04:05 -07:00") + `
---
			
` + record[4])

		category := record[1]
		path := "output/" + categories[category]
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.MkdirAll(path, 0700) // Create your file
		}
		_ = os.WriteFile(path+"/"+record[15]+".md", md, 0644)
	}

	return
}
