//Package api ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 11:40:36
 * @LastEditors: congz
 * @LastEditTime: 2020-08-12 20:58:01
 */
package api1

// import (
// 	"cmall/pkg/logging"
// 	"cmall/service"
// 	"fmt"

// 	"github.com/gin-gonic/gin"
// )

// // CreateProduct 创建商品
// func CreateProduct(c *gin.Context) {
// 	service := service.CreateProductService{}
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.Create()
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 		logging.Info(err)
// 	}
// }

// // ListProducts 商品列表接口
// func ListProducts(c *gin.Context) {
// 	var service service.ListProductsService
// 	if err := c.ShouldBind(&service); err == nil {
// 		fmt.Printf("service: %v\n", service)
// 		res := service.List()
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 		logging.Info(err)
// 	}

// }

// // ShowProduct 商品详情接口
// func ShowProduct(c *gin.Context) {
// 	service := service.ShowProductService{}
// 	res := service.Show(c.Param("id"))
// 	c.JSON(200, res)
// }

// // DeleteProduct 删除商品的接口
// func DeleteProduct(c *gin.Context) {
// 	service := service.DeleteProductService{}
// 	res := service.Delete(c.Param("id"))
// 	c.JSON(200, res)
// }

// // UpdateProduct 更新商品的接口
// func UpdateProduct(c *gin.Context) {
// 	service := service.UpdateProductService{}
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.Update()
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 		logging.Info(err)
// 	}
// }

// // SearchProducts 搜索商品的接口
// func SearchProducts(c *gin.Context) {
// 	service := service.SearchProductsService{}
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.Show()
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 		logging.Info(err)
// 	}
// }
