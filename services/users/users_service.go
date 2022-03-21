package user_service

import (
	"log"
	contracts "qmdapipoc/contracts/users"
	"qmdapipoc/domain/users"
	"qmdapipoc/utils/errors"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User) (*users.User, *errors.CustomError) {
	log.Println("CreateUser")
	if err := user.Validate(); err != nil {
		return nil, err
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		return nil, errors.NewCustomError("PasswordEncryptionFailed", "005")
	}

	user.Password = string(pwd[:])

	log.Println(user.Password)

	if err := user.Save(); err != nil {

		if err.ErrorCode == "UserEmailNotUnique" {
			log.Println("UserEmailNotUnique Service exception")
		}
		return nil, err
	}

	return &user, nil
}

func GetUserbyEmail(user users.User) (*users.User, *errors.CustomError) {
	result := &users.User{Email: user.Email}

	if err := result.GetByEmail(); err != nil {
		return nil, err
	}

	resultWp := &users.User{
		ID:        result.ID,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
	}

	return resultWp, nil
}

func GetUserCollection() ([]contracts.UserOutContract, *errors.CustomError) {

	user_collection, collection_error := users.GetUserCollection()

	if collection_error != nil {
		return nil, collection_error
	}

	out_contracts := make([]contracts.UserOutContract, 0)

	for _, user_contract := range user_collection {
		contract := contracts.UserOutContract{
			UserId:    user_contract.UserGlobalKey,
			FirstName: user_contract.FirstName,
			LastName:  user_contract.LastName,
			Email:     user_contract.Email,
		}
		out_contracts = append(out_contracts, contract)
	}

	return out_contracts, nil
}
