package service

import (
	"auth-service/dto"
	"auth-service/model"
	"auth-service/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type AuthService struct {
	AuthRepository *repository.AuthRepository
}

func (service *AuthService) RegisterUser (dto dto.RegisterDTO) error {
	hashPw, _ := HashPassword(dto.Password)
	user := model.User{ Email: dto.Email, UserRole: 0, Password: hashPw, Username: dto.Username}
	err := service.AuthRepository.RegisterUser(&user)
	if err != nil {
		return err
	}
	requsetBody , jerr := json.Marshal(dto)
	if jerr != nil{
		return jerr
	}
	client := &http.Client{}
	requestUrl := fmt.Sprintf("http://%s:%s/registerUser", os.Getenv("USER_SERVICE_DOMAIN"), os.Getenv("USER_SERVICE_PORT"))
	req, _ := http.NewRequest("POST", requestUrl, bytes.NewBuffer(requsetBody))
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
		return nil
	}
	body, err5 := ioutil.ReadAll(res.Body)
	if err5 != nil {
		log.Fatalln(err5)
	}
	//Convert the body to type string
	sb := string(body)
	fmt.Println(sb)

	return nil
}


func (service *AuthService) FindByUsername (dto dto.LogInDTO) (*model.User, error){
	user, err := service.AuthRepository.FindUserByUsername(dto.Username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


func (service *AuthService) ChangePassword(username string, passwords dto.PasswordChangerDTO) (*model.User, error) {
	user, err := service.AuthRepository.FindUserByUsername(username)
	if CheckPasswordHash(passwords.PasswordOld,user.Password){
		user.Password, _ = HashPassword(passwords.PasswordNew)
		// TODO
		//err = service.AuthRepository.UpdateUser(user)
	}
	return user, err
}

func (service *AuthService) Authenticate(username string) (model.Role, error){
	user, err := service.AuthRepository.FindUserByUsername(username)
	if err != nil {
		return -1, err
	}
	return user.UserRole, nil

}

func (service *AuthService) DeleteUser(username string)  {
	service.AuthRepository.Delete(username)
}

