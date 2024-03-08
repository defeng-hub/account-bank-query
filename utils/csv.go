package utils

import (
	"encoding/csv"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"os"
)

func ReadCSV(filename string) ([][]string, error) {
	// 打开CSV文件
	file, err := os.Open(filename) // data.csv为要读取的CSV文件名
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// utf-8
	// reader := csv.NewReader(file)

	// GBK方式
	reader := csv.NewReader(transform.NewReader(file, simplifiedchinese.GBK.NewDecoder()))

	// 逐行读取数据
	return reader.ReadAll()
}
