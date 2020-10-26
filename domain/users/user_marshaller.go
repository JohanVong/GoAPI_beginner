package users

import "encoding/json"

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
func (users Users) Marshall() []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall()
	}
	return result
}

// Marshall determines what type of user to return
func (user *User) Marshall() interface{} {
	var privateUser PrivateUser

	userJSON, _ := json.Marshal(user)
	json.Unmarshal(userJSON, &privateUser) // Unmarshal(кого, куда)

	return privateUser
}
