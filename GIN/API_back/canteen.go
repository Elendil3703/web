//Package api ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-12 22:15:31
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 17:47:42
 */
package api2

// import (
// 	"cmall/pkg/logging"
// 	"cmall/service"

// 	"github.com/gin-gonic/gin"
// )

// // CreateCategory 创建分类
// func CreateCanteen(c *gin.Context) {
// 	service := service.CreateCanteenService{}
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.Create()
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 		logging.Info(err)
// 	}
// }

// // ListCategories 分类列表接口
// func ListCanteens(c *gin.Context) {
// 	service := service.ListCanteensService{}
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.List()
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 		logging.Info(err)
// 	}
// }
