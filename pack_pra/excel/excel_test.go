package pack_pra

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"testing"
)

func TestCreate(t *testing.T) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	row := sheet.AddRow()
	row.SetHeightCM(1) //设置每行的高度
	cell := row.AddCell()
	cell.Value = "haha"
	cell = row.AddCell()
	cell.Value = "xxx"

	err = file.Save("file.xlsx")
	if err != nil {
		panic(err)
	}
}

func TestRead(t *testing.T) {
	wb, err := xlsx.OpenFile("file.xlsx")
	if err != nil {
		return
	}
	fmt.Println("Sheets in this file:")
	for i, sh := range wb.Sheets {
		fmt.Println(i, sh.Name)
	}
	fmt.Println("----")
}
