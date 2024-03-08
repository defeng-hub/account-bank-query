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
	WordReplace("学生毕业论文", "替换后", "demo.docx", "out.docx")
}

func TestExcelReplace(t *testing.T) {
	ExcelReplace("山西省", "替换后", "demo.xlsx", "out.xlsx")
}
