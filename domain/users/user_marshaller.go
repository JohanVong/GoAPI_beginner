package users

import "encoding/json"

// PublicUser is public info
type PublicUser struct {
	Username    string `json:"username"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// PrivateUser is for authenticated user
type PrivateUser struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	IsAdmin     bool   `json:"is_admin"`
}

// Marshall for array of User
func (users Users) Marshall(isPrivate bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPrivate)
	}
	return result
}

// Marshall determines what type of user to return
func (user *User) Marshall(isPrivate bool) interface{} {
	var privateUser PrivateUser
	var publicUser PublicUser

	userJSON, _ := json.Marshal(user)
	json.Unmarshal(userJSON, &privateUser) // Unmarshal(кого, куда)
	json.Unmarshal(userJSON, &publicUser)

	if isPrivate {
		return privateUser
	}

	return publicUser
}
