package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/wissemmansouri/OpenIT.one-Common/utils/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		// Allow cross-origin settings to return other subdomains and customize fields
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,Language,Content-Type,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Access-Control-Allow-Methods,Connection,Host,Origin,Referer,User-Agent,X-Requested-With")
		// Allow headers that the browser (client) can parse (important)
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		// c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")
		// Set cache time
		c.Header("Access-Control-Max-Age", "172800")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Set("Content-Type", "application/json")
		//}

		// Allow type validation
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}
		c.Request.Header.Del("Origin")
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		c.Next()
	}
}

func WriteLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.Request.URL.String(), "password") {
			logger.Info("request:", zap.Any("path", c.Request.URL.String()), zap.Any("param", c.Params), zap.Any("query", c.Request.URL.Query()), zap.Any("method", c.Request.Method))
			c.Next()
		}
	}
}
