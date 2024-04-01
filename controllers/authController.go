package controllers

import (
	"fmt"
	"goProject/models"
	"goProject/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func signToken(id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expire time (1 day)
	return token.SignedString([]byte("your-secret-key"))  // Change "your-secret-key" with your actual secret key
}

func createSendToken(user *models.User, statusCode int, res http.ResponseWriter) {
	tokenString, err := signToken(user.ID)
	if err != nil {
		utils.RespondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}

	cookieOptions := http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24), // Cookie expire time (1 day)
		HttpOnly: true,
	}

	if utils.IsProductionEnvironment() {
		cookieOptions.Secure = true
	}

	http.SetCookie(res, &cookieOptions)

	user.Password = "" // Remove password before sending response

	utils.RespondWithJSON(res, statusCode, map[string]interface{}{
		"status": "success",
		"token":  tokenString,
		"data":   user,
	})
}

func Signup(w http.ResponseWriter, r *http.Request) {
	newUser := models.User{
		Name:            r.FormValue("name"),
		Email:           r.FormValue("email"),
		Password:        r.FormValue("password"),
		PasswordConfirm: r.FormValue("passwordConfirm"),
	}

	// Save the new user to the database
	err := models.AddUser(newUser)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Send welcome email
	go func() {
		url := fmt.Sprintf("%s://%s/me", r.URL.Scheme, r.Host)
		if err := utils.SendWelcomeEmail(newUser.Email, url); err != nil {
			fmt.Println("Error sending welcome email:", err)
		}
	}()

	createSendToken(&newUser, http.StatusCreated, w)
}

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := models.GetUserByEmail(email)
	if err != nil || !user.VerifyPassword(password) {
		utils.RespondWithError(w, http.StatusUnauthorized, "Incorrect email or password")
		return
	}

	createSendToken(user, http.StatusOK, w)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "loggedout",
		Expires:  time.Now().Add(10 * time.Second),
		HttpOnly: true,
	})
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
