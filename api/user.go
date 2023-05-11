package api

import (
	"net/http"
	"time"
	"travelfanapi/dao"
	"travelfanapi/domain"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authenticateBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateResponseBody struct {
	Token       string                    `json:"token"`
	CurrentUser *domain.User `json:"currentUser"`
}

func Register(c echo.Context) error {
	u := new(domain.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	// Hasheamos la contraseña antes de guardarla
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	db := c.Get("DB").(*gorm.DB)
	if err := db.Create(u).Error; err != nil {
		return err
	}

	// Logueamos al usuario inmediatamente después del registro
	token, err := generateToken(u.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"token": token,
	})
}

func Authenticate(c echo.Context) error {
	var body authenticateBody

	db := c.Get("DB").(*gorm.DB)

	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := dao.GetUserByLogin(db, body.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		return err
	}

	if passwordMatch := user.CheckPassword(body.Password); !passwordMatch {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	tokenString, err := generateToken(user.ID)
	if err != nil {
		logrus.Errorf("Error in api.Authenticate : error while signing the token: %s", err)
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:   "token",
		Value:  tokenString,
		Path:   "/",
	})

	return c.JSON(http.StatusOK, AuthenticateResponseBody{
		tokenString,
		user,
	})
}

func generateToken(userID uint) (string, error) {
	// Creamos el token con una duración de 24 horas
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return token.SignedString([]byte("secret"))
}
