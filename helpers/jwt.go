package helpers

import (
	"final-task-pbi-rakamin/app"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(user app.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}

// var secretKey = "rahasia"

// func GenerateToken(id uint64, email string) string {
// 	claims := jwt.MapClaims{
// 		"id":    id,
// 		"email": email,
// 	}

// 	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	signedToken, _ := parseToken.SignedString([]byte(secretKey))

// 	return signedToken
// }

// var secretKey = os.Getenv("JWT_SECRET")

// func GenerateToken(id uint, email string) string {
// 	claims := jwt.MapClaims{
// 		"id":    id,
// 		"email": email,
// 	}

// 	// creates a new token with the specified signing method and claims.
// 	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// creates and returns a complete, signed JWT.
// 	signedToken, _ := parseToken.SignedString([]byte(secretKey))

// 	return signedToken
// }

// func VerifyToken(c *gin.Context) (interface{}, error) {
// 	// init error message
// 	errResponse := errors.New("sign in to proceed")
// 	// get Authorization header value
// 	headerToken := c.Request.Header.Get("Authorization")
// 	// check if Authorization header contains Bearer as a suffix
// 	if bearer := strings.HasPrefix(headerToken, "Bearer"); !bearer {
// 		return nil, errResponse
// 	}

// 	// headerToken: Bearer <token-here>
// 	// get the <token-here> value after splitting inside index 1
// 	stringToken := strings.Split(headerToken, " ")[1]

// 	// parse token into a pointer of struct jwt.Token
// 	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
// 		// check if signing method is HS256 by casting the method into pointer of struct jwt.SigningMethodHMAC
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, errResponse
// 		}
// 		return []byte(secretKey), nil
// 	})

// 	// check if the token is valid and not nil
// 	if token != nil && token.Valid {
// 		// perform a type assertion to check if token claims can be asserted as jwt.MapClaims
// 		// reference: https://go.dev/tour/methods/15
// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if !ok {
// 			return nil, errResponse
// 		}

// 		// return claims (contains id & email of the successfully logged in user)
// 		return claims, nil
// 	}

// 	return nil, errResponse
// }
