package usermanager

import (
	"go-apps/auth.com/input"
	"go-apps/auth.com/model"
	sessionsdao "go-apps/auth.com/persistence/sessions-dao"
	usersdao "go-apps/auth.com/persistence/users-dao"
	"strconv"
)

func GetUserById(id string) *model.UserRecord {
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil
	}
	return usersdao.User(uint(userId))
}

func GetUserByEmail(email string) *model.UserRecord {
	return usersdao.UserByEmail(email)
}

func CreateUser(input input.UserInput) error {
	return usersdao.InsertUser(input)
}

func GetSessionByToken(token string) *model.SessionRecord {
	return sessionsdao.GetSessionByToken(token)
}

func GetUsers(limit int) ([]*model.UserRecord, error) {
	return usersdao.Users(limit)
}
