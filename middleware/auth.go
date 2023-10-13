package middleware

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"startup/auth"
	"startup/helper"
	"startup/users"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authMidd struct {
	authService auth.Service
	userService users.Service
}

func NewAuthMiddleware(userService users.Service, authService auth.Service) *authMidd {
	return &authMidd{
		userService: userService,
		authService: authService,
	}
}

func (a *authMidd) AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.Contains(authHeader, "Bearer") {
		response := helper.ApiResponse(http.StatusUnauthorized, nil, "Unauthorization", "Error")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	var stringToken string
	arrayHeader := strings.Split(authHeader, " ")
	if len(arrayHeader) == 2 {
		stringToken = arrayHeader[1]
	}
	jwtToken, err := a.authService.ValidateToken(stringToken)
	if err != nil {
		response := helper.ApiResponse(http.StatusUnauthorized, nil, err.Error(), "Error")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	claim, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		response := helper.ApiResponse(http.StatusUnauthorized, nil, "Unauthorized", "Error")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	userId := int(claim["user_id"].(float64))
	user, err := a.userService.GetUserById(userId)
	if err != nil {
		response := helper.ApiResponse(http.StatusUnauthorized, nil, "Unauthorized", "Error")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	c.Set("currentUser", user)

	// token, err :=
}


func  (a *authMidd) LoggingMiddleware(c *gin.Context) {
    // Waktu awal permintaan
    startTime := time.Now()

    // Mengecek apakah ada body request
    if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
        // Membaca dan mencatat body permintaan
        requestBody, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
            log.Printf("Error reading request body: %v\n", err)
        }

        // Mengembalikan body request ke kondisi semula
        c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

        // Mencetak log
        log.Printf("Request Body: %s\n", string(requestBody))
    }

    // Memproses permintaan
    c.Next()

    // Waktu selesai permintaan
    endTime := time.Now()

    // Mendapatkan informasi permintaan
    method := c.Request.Method
    path := c.Request.URL.Path
    statusCode := c.Writer.Status()

    // Mencetak log
    log.Printf("[%s] %s %s %d %v\n", method, path, c.ClientIP(), statusCode, endTime.Sub(startTime))
}