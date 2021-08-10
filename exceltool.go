package main

import (
	"flag"
	"fmt"
	excelize "github.com/xuri/excelize"
	"strconv"
)

func main() {
	var file = flag.String("file", "", "")
	var sheet = flag.String("sheet", "", "")
	var axis = flag.String("axis", "", "")
	var celltype = flag.String("type", "", "")
	var value = flag.String("value", "", "")
	flag.Parse()
	f, err := excelize.OpenFile(*file)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SaveAs(*file); err != nil {
		fmt.Println(err)
	}
	if *celltype == "number" {
		value1, _ := strconv.Atoi(*value)
		f.SetCellInt(*sheet, *axis, value1)
	}
	if *celltype == "text" {
		f.SetCellValue(*sheet, *axis, *value)
	}
	if *celltype == "formula" {
		f.SetCellFormula(*sheet, *axis, *value)
	}
	if err := f.SaveAs(*file); err != nil {
		fmt.Println(err)
	}
}
