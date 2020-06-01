package zip_test

import (
	"../zip"
	"os"
	"testing"
)

func TestCompress(t *testing.T) {
	f1, err := os.Open("D:\\workspace\\gopath\\src\\statMallGroupNotify\\shanghai.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	defer f1.Close()
	f2, err := os.Open("D:\\workspace\\gopath\\src\\statMallGroupNotify\\suzhou.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	defer f2.Close()
	f3, err := os.Open("D:\\workspace\\gopath\\src\\statMallGroupNotify\\tianNing.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	defer f3.Close()
	var files = []*os.File{f1, f2, f3}
	dest := "D:\\workspace\\gopath\\src\\statMallGroupNotify\\test.zip"
	err = zip.Compress(files, dest)
	t.Fatal(err)
}
