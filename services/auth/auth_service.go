package authentication

import (
	"qmdapipoc/domain/users"
	"qmdapipoc/utils/errors"

	"golang.org/x/crypto/bcrypt"
)

func GetUser(user users.User) (*users.User, *errors.CustomError) {
	result := &users.User{Email: user.Email}

	if err := result.GetByEmail(); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return nil, errors.NewCustomError("failed to decrypt", "decrypt")
	}

	resultWp := &users.User{
		ID:            result.ID,
		UserGlobalKey: result.UserGlobalKey,
		FirstName:     result.FirstName,
		LastName:      result.LastName,
		Email:         result.Email,
	}

	return resultWp, nil
}
