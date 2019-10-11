package main

import (
	"github.com/tealeg/xlsx"
)

func main() {
	file := xlsx.NewFile()
	sheet := file.AddSheet("Sheet1")
	row := sheet.AddRow()
	row.SetHeightCM(1) //设置每行的高度
	cell := row.AddCell()
	cell.Value := "haha"
	cell = row.AddCell()
	cell.Value = "xxx"

	err:=file.Save("file.xlsx")
	if err != nil{
		panic(err)
	}
}
