package utils

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestReadCSV(t *testing.T) {
	ast := assert.New(t)
	list, err := ReadCSV("test.csv")
	ast.NoError(err)
	t.Log(list)
}

func TestRun(t *testing.T) {
	// Load a docx file
	r, err := ReadDocxFile("王德丰开题报告表.docx")
	if err != nil {
		log.Fatalf("error reading docx file: %s", err)
	}
	defer r.Close()

	doc := r.Editable()

	// Replace text
	doc.Replace("已研读的有关文献资料", "newText", -1)

	// Save the modified document
	err = doc.WriteToFile("modified.docx")
	if err != nil {
		log.Fatalf("error writing docx file: %s", err)
	}
}
