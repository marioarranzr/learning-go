package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	fmt.Println("\n", "*******************    parseFiles    *******************")
	parseFiles()
	fmt.Println("\n", "*******************    parseGlob     *******************")
	parseGlob()
}

func parseFiles() {
	tpl, err := template.ParseFiles("tpl1.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Create new file
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}
	defer nf.Close()

	// Print in stdout
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Print in the new file created
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func parseGlob() {
	// Loads all files with that extension
	tpl, err := template.ParseGlob("*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Execute all in tpl
	// Print in the new file created
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Execute only tpl.gohtml file stored in tpl
	// Print in stdout
	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
