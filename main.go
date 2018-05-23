package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

type Icon struct {
	Name  string
	Value string
}

func main() {
	var icons []Icon

	removeNewlines := regexp.MustCompile(`\r?\n`)
	stripDoctype := regexp.MustCompile(`<\?xml.*><\!DOCTYPE.*><svg`)

	iconNames := readIconNames("./icons.txt")
	for _, v := range iconNames {

		if v == "" {
			continue
		}

		iconContent, err := ioutil.ReadFile(filepath.Join("MaterialDesign", "icons", "svg", v+".svg"))
		if err != nil {
			log.Fatalf("Could not read icon %s: %v\n", v, err)
			return
		}

		iconContentString := string(iconContent)
		iconContentString = removeNewlines.ReplaceAllString(iconContentString, "")
		iconContentString = stripDoctype.ReplaceAllString(iconContentString, "<svg")

		icons = append(icons, Icon{
			Name: v,
			// THE first three bytes are the BOM, which we need to skip
			Value: iconContentString[3:],
		})
	}

	applyIconsToTemplate(icons)
}

func readIconNames(path string) []string {
	icns, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Could not open icon list file: %v", err)
		return nil
	}
	iconList := strings.Split(string(icns), "\n")
	return iconList
}

func applyIconsToTemplate(icons []Icon) {
	tmpl := template.Must(template.ParseFiles("./template.ts"))
	f, err := os.Create("icons.ts")
	if err != nil {
		log.Fatalf("Could not open output file: %v\n", err)
		return
	}
	err = tmpl.Execute(f, icons)
	if err != nil {
		log.Fatalf("Could not execute template: %v\n", err)
	}
	f.Close()
}
