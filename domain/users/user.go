package users

import (
	"log"
	storage "qmdapipoc/storages"
	"qmdapipoc/utils/errors"
	"strings"
)

type User struct {
	ID            int64  `json:"ID"`
	UserGlobalKey string `json:"user_global_key"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Password      string `json:"password"`
	Email         string `json:"email"`
}

var (
	createUserQuery        = `INSERT INTO "User"."Users"(first_name, last_name, email, password) VALUES($1, $2, $3, $4);`
	getUserByEmailQuery    = `SELECT 1 user_id, user_global_key, first_name, last_name, email, password FROM "User"."Users" WHERE email = $1;`
	getUserByIdQuery       = `SELECT 1 user_id, user_global_key, first_name, last_name, email, password FROM "User"."Users" WHERE user_global_key = $1;`
	getUserCollectionQuery = `SELECT user_id, user_global_key, first_name, last_name, email, password FROM "User"."Users"`
)

func (user *User) Validate() *errors.CustomError {

	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if user.FirstName == "" {
		return errors.NewCustomError("UserFirstNameNotValid", "001")
	}

	if user.LastName == "" {
		return errors.NewCustomError("UserLastNameNotValid", "002")
	}

	if user.Password == "" {
		return errors.NewCustomError("UserPasswordNotValid", "003")
	}

	if user.Email == "" {
		return errors.NewCustomError("UserEmailNotValid", "004")
	}

	return nil
}

func (user *User) Save() *errors.CustomError {

	log.Println(createUserQuery)

	stmt, err := storage.Client.Prepare(createUserQuery)

	if err != nil {
		log.Println(err.Error())
		return errors.NewCustomError("DbError", "DB001", err.Error())
	}

	defer stmt.Close()

	result, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)

	if saveErr != nil {

		if strings.Contains(saveErr.Error(), "u_user_users_email") {
			log.Println("u_user_users_email Service db/domain exception")
			return errors.NewCustomError("UserEmailNotUnique", "011", saveErr.Error())
		}

		return errors.NewCustomError("FailedToSave", "", saveErr.Error())
	}

	log.Println(result)

	return nil
}

func (user *User) GetByEmail() *errors.CustomError {
	stmt, err := storage.Client.Prepare(getUserByEmailQuery)

	if err != nil {
		log.Println(err.Error())
		return errors.NewCustomError("invalid email", "DB001", err.Error())
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Email)

	if getErr := result.Scan(&user.ID, &user.UserGlobalKey, &user.FirstName, &user.LastName, &user.Email, &user.Password); getErr != nil {
		log.Println("GetByEmail user found")
		return errors.NewCustomError("GetByEmail", "DbError", getErr.Error())
	}

	log.Println("user found")

	return nil
}

func GetUserCollection() ([]User, *errors.CustomError) {

	user_collection := make([]User, 0)

	user_collection_result, err := storage.Client.Query(getUserCollectionQuery)

	if err != nil {
		log.Fatal(err)
	}

	defer user_collection_result.Close()

	log.Println("user_collection_result found", user_collection_result)

	for user_collection_result.Next() {
		var user User
		err := user_collection_result.Scan(&user.ID, &user.UserGlobalKey, &user.FirstName, &user.LastName, &user.Email, &user.Password)

		if err != nil {
			log.Fatal(err)
		}

		user_collection = append(user_collection, user)
	}

	return user_collection, nil
}

func GetUser(user_global_key string) (User, *errors.CustomError) {

	user := User{}

	user_result, err := storage.Client.Prepare(getUserByIdQuery)

	if err != nil {
		log.Fatal(err)
	}

	defer user_result.Close()

	result := user_result.QueryRow(user_global_key)

	if getErr := result.Scan(&user.ID, &user.UserGlobalKey, &user.FirstName, &user.LastName, &user.Email, &user.Password); getErr != nil {
		return user, errors.NewCustomError("UserNotFound", "0001", getErr.Error())
	}

	return user, nil
}
