//Package middleware ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-15 14:48:46
 * @LastEditors: congz
 * @LastEditTime: 2020-08-08 16:33:29
 */
package middleware

import (
	"gintest/pkg/e"
	"gintest/pkg/util"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	//主要是验证token未过期，并未验证其携带的username和userid是否匹配
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			} else {
				//fmt.Println(claims)
				c.Set("usermsg", claims) //将这里的claims保存到c.keys中
			}
		}
		if code != e.SUCCESS {
			c.JSON(404, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// JWTAdmin token验证中间件
func JWTAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			} else if claims.Authority == 0 {
				code = e.ERROR_AUTH_INSUFFICIENT_AUTHORITY
			}
		}

		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
