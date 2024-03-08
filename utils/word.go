package utils

import (
	"fmt"
	"github.com/carmel/gooxml/document" //版本号：v0.0.0-20220216072414-40ff56130850
	"log"
	"strings"
)

func WordReplace(search, replace string, filepath, outpath string) (err error) {
	r, err := ReadDocxFile(filepath)
	if err != nil {
		log.Fatalf("error reading docx file: %s", err)
		return err
	}
	defer r.Close()

	doc := r.Editable()

	// Replace text
	doc.Replace(search, replace, -1)

	// Save the modified document
	err = doc.WriteToFile(outpath)
	if err != nil {
		log.Fatalf("error writing docx file: %s", err)
		return err
	}
	return nil
}

func TestTable() {
	// 打开Word文档
	doc, err := document.Open("demo.docx")
	if err != nil {
		log.Fatalf("error opening document: %s", err)
	}

	// 要查找和替换的字符串
	search := "测试数据1"
	replace := "newString"

	// 遍历文档中的表格
	for _, tbl := range doc.Tables() {
		// 遍历表格中的行
		for _, row := range tbl.Rows() {
			// 遍历行中的单元格
			for _, cell := range row.Cells() {
				// 遍历单元格中的段落
				for _, para := range cell.Paragraphs() {
					// 遍历段落中的runs
					for _, run := range para.Runs() {
						text := run.Text()
						fmt.Println(text)
						//// 如果找到了要替换的字符串，就进行替换
						if strings.Contains(text, search) {
							// 替换文本
							newText := strings.Replace(text, search, replace, -1)
							run.ClearContent()
							run.AddText(newText)
						}
					}
				}
			}
		}
	}

	// 保存修改后的文档
	err = doc.SaveToFile("modified_document.docx")
	if err != nil {
		log.Fatalf("error saving document: %s", err)
	}
}
