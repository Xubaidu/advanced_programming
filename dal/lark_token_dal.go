package dal

import (
	"advanced_programming/clients"
	"advanced_programming/common"
	"advanced_programming/constant"
	"advanced_programming/models"
	"log"
)

func GetLarkToken(larkTokenID int) (token string, err error) {
	var DB = clients.DB
	larkToken := &models.LarkToken{}
	DB.Where(map[string]interface{}{"id": larkTokenID}).First(larkToken)
	if larkToken == nil {
		log.Printf("没有符合条件的查询")
		err = common.NewError(constant.DataQueryError, "没有符合条件的查询")
		return "", err
	}
	return larkToken.Token, nil
}
