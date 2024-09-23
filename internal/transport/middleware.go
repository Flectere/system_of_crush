package transport

import (
	"net/http"
	"strings"

	"github.com/Flectere/system_of_crush/internal/service"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлекаем заголовок Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Проверка что заголовок в заголовке лежит токен
		tokenString := strings.Split(authHeader, "Bearer ")
		if len(tokenString) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		user_id, err := service.ValidateJWT(tokenString[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Если всё корректно, продолжаем выполнение запроса
		c.Set("user_id", user_id)
		c.Next()
	}
}
