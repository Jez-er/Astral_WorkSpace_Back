package middleware

import (
	auth "Astral/internal/jwt/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

/* Миддлвэйр, типа таможни или пограничников перед городом */
func AuthZ()gin.HandlerFunc{
	return func(c *gin.Context){
		clientToken := c.Request.Header.Get("Authorization")

		if clientToken == ""{
			c.JSON(403,"client token empty")
			c.Abort()
			return
		}

		extractedToken := strings.Split(clientToken,"Bearer ")
		if len(extractedToken) == 2 {
			
			clientToken = strings.TrimSpace(extractedToken[1])
		   } else {
		
			c.JSON(400, "Incorrect Format of Authorization Token")
			c.Abort()
			return
		   }
		   jwtWrapper := auth.JwtWrapper{
			SecretKey: "verysecretkey",
			Issuer: "AuthService",
		   }
		   claims, err := jwtWrapper.ValidationToken(clientToken)
		   if err != nil {
		
			c.JSON(401, err.Error())
			c.Abort()
			return
		   }
		  
		   c.Set("email", claims.Email)
		   
		   c.Next()
	}
}