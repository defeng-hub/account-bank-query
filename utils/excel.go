package utils

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func ExcelReplace(search, replace string, filepath, outpath string) (err error) {
	// 打开Excel文件
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	sheetName := f.GetSheetName(0)
	// 获取Sheet1上所有单元格
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for i, row := range rows {
		for j, colCell := range row {
			if colCell == search {
				// 构建单元格的坐标，如A1，B2等
				cell, _ := excelize.CoordinatesToCellName(j+1, i+1)
				//设置新的值
				f.SetCellValue(sheetName, cell, replace)
			}
		}
	}

	// 保存修改后的Excel文件
	if err = f.SaveAs(outpath); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
