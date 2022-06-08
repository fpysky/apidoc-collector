package main

import (
	"api-doc-generator/generator"
)

func main() {
	g := generator.New("./../")
	g.BeginScanningApiFolder()
}
