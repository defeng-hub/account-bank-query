package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadCSV(t *testing.T) {
	ast := assert.New(t)
	list, err := ReadCSV("test.csv")
	ast.NoError(err)
	t.Log(list)
}

func TestWordReplace(t *testing.T) {
	ass := assert.New(t)
	err := WordReplace("测试数据1", "替换后", "demo.docx", "out.docx")
	ass.NoError(err)
}

func TestExcelReplace(t *testing.T) {
	ass := assert.New(t)
	err := ExcelReplace("山西省", "替换后", "demo.xlsx", "out.xlsx")
	ass.NoError(err)

}

func TestRun(t *testing.T) {
	Run()
}
