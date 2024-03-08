package main

import (
	"account-bank-query/model"
	"account-bank-query/model/dto"
	"account-bank-query/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var BankBinCodeMap = make(map[string]*model.BankBinCode, 2000)

func init() {
	list, err := utils.ReadCSV("doc/bank-bin-code.csv")
	if err != nil {
		panic(err)
	}
	for _, row := range list[1:] {
		var atoi int
		var err1 error
		atoi, err1 = strconv.Atoi(row[3])
		if err1 != nil {
			atoi = len(row[0])
		}

		BankBinCodeMap[row[0]] = &model.BankBinCode{
			BankCode:       row[0],
			BankName:       row[1],
			BankShortName:  row[2],
			BankCardLength: int64(atoi),
		}
	}
	fmt.Println(BankBinCodeMap)

}

// 传入银行卡号, 返回银行信息
func getBank(bankCode dto.BankCode) *model.BankBinCode {
	var k int = 6 // 从第6位开始检索, 最大为10
	for {
		if k > 10 {
			break
		}
		// 取出
		if v, ok := BankBinCodeMap[bankCode.Prefix(k)]; ok {
			return v
		}
		k++
		continue
	}
	return nil
}

func main() {
	var err error

	r := gin.Default()

	//配置模板文件
	r.LoadHTMLGlob("templates/*")
	//配置静态文件目录
	r.Static("/static", "static")

	//配置路由
	r.GET("/", func(c *gin.Context) {
		q := c.Request.URL.Query()
		if v, ok := q["bank_code"]; ok {
			if len(v) > 0 {
				value := getBank(dto.BankCode(v[0]))
				c.HTML(http.StatusOK, "index.html", gin.H{"show": true, "data": value})
				return
			}
		}
		c.HTML(http.StatusOK, "index.html", nil)
		return
	})

	r.POST("/search", func(c *gin.Context) {
		var req dto.BankCodeDto
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(req.BankCodes) >= 0 {
			var list []*model.BankBinCode
			for _, v := range req.BankCodes {
				// 弄个列表
				list = append(list, getBank(v))
			}
			c.JSON(http.StatusOK, gin.H{"data": list})
			return
		}

		value := getBank(req.BankCode)
		c.JSON(http.StatusOK, gin.H{"data": value})
		return
	})

	//配置端口
	err = r.Run(":8099")
	if err != nil {
		return
	}
}
