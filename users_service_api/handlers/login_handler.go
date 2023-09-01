package handlers

import (
	"net/http"
	"time"
	"users_service_api/entities"
	"users_service_api/repositories"
	"users_service_api/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type LoginHandler struct {
	userRepository *repositories.UserRepository
	utilsFunctions *utils.Functions
}

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (rh *LoginHandler) Login(c *gin.Context) {
	var body LoginRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := rh.userRepository.GetUserByUsername(body.Username)
	if err != nil {
		c.JSON(500, gin.H{"Error": "Error al consultar el usuario", "message": err})
		return
	}

	compare := rh.utilsFunctions.ComparePassword(user.Password, body.Password)
	if compare != nil {
		c.JSON(403, gin.H{"Error": "Contraseña incorrecta", "message": err})
		return
	}

	token, err := rh.generateToken(*user)
	if err != nil {
		c.JSON(500, gin.H{"Error": "Error al generar el token", "message": err})
		return
	}

	c.JSON(200, gin.H{"Status": "logged", "Token": token})

}

func (rh *LoginHandler) generateToken(user entities.User) (string, error) {
	var jwtSecret = []byte("secret_key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expira en 24 horas
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
