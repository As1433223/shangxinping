package models

func GetUserInfo(user User) (User, error) {
	var NewUser User
	res := db.Where(user).First(&NewUser)
	return NewUser, res.Error
}
