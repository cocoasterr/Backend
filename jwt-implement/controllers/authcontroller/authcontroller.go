package authcontroller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/cocoaster/golang-jwt/configs"
	"github.com/cocoaster/golang-jwt/helper"
	"github.com/cocoaster/golang-jwt/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)
func HealthCheck(w http.ResponseWriter, r *http.Request){
	response,_ := json.Marshal(map[string]string{"message": "success"})
	w.Header().Add("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func isValidEmail(email string) bool {
	// Regular expression untuk memeriksa bentuk email
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailPattern, email)
	return match
}


func checkingEmail(payload models.User, w http.ResponseWriter)int{
	if payload.Email == ""{
		response := map[string]string{"message": "Email Required!"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return http.StatusBadRequest
	}

	if !isValidEmail(payload.Email){
		response := map[string]string{"message": "Email Invalid!"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return http.StatusBadRequest
	}
	return http.StatusOK
}

func Register(w http.ResponseWriter, r *http.Request){
	//take input json
	var userInput models.User
	var response map[string]string

	//type assertion with go, 
	//because return is interface{} we should intepretation return that return is model.User 
	payload := helper.GetJson(userInput, w ,r).(*models.User)
	//checking json input
	checkEmail := checkingEmail(*payload, w)
	if checkEmail != 200{
		return
	}
	if len(payload.Password) < 8{
		response = map[string]string{"message": "Password must more 8 Characters!"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Check if email already exists in the database
	var existingUser models.User
	if err := models.DB.Where("email = ?", payload.Email).First(&existingUser).Error; err == nil {
		response = map[string]string{"message": "Email already exists!"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Hash Password
	hashPassword,err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil{
		response = map[string]string{"message": "Password Hashing failed!"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	payload.Password = string(hashPassword)

	//create json to database
	if err:= models.DB.Create(&payload).Error;err!=nil{
		response = map[string]string{"message": "Create user failed!"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	
	//return result
	response = map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}



func Login(w http.ResponseWriter, r *http.Request){
	var userInput models.User
	var response map[string]string
	
	err:= godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .Env file")
	}
	JWT_KEY := []byte(os.Getenv("JWT_KEY"))
	payload := helper.GetJson(&userInput, w, r).(*models.User)

	checkEmail := checkingEmail(*payload, w)
	if checkEmail != 200{
		return
	}
	hashPassword,err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil{
		response = map[string]string{"message": "Password Hashing failed!"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	payload.Password = string(hashPassword)
	if err := models.DB.Where("email=?, password=?",payload.Email, payload.Password).Error; err!= nil{
		response = map[string]string{"message": "Email and Password is wrong!"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	//JWT 
	expTime:= time.Now().Add(time.Minute * 1)
	claims := &configs.JwtClaims{
		Username: payload.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "golang-jwt",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	//choice algoritm
	algo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	//sign token
	token, err:= algo.SignedString(JWT_KEY)
	if err != nil{
		response = map[string]string{"message": "Create token failed!"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	//set cookie
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})

	response = map[string]string{"message": "login success!"}
	helper.ResponseJSON(w, http.StatusOK, response)
}


func Logout(w http.ResponseWriter, r *http.Request){
		//remove cookie
		http.SetCookie(w, &http.Cookie{
			Name: "token",
			Path: "/",
			Value: "",
			HttpOnly: true,
			MaxAge: -1,
		})
	
		response := map[string]string{"message": "logout success!"}
		helper.ResponseJSON(w, http.StatusOK, response)
}