package datastruct

type User struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}

func NewUser(Id, name, email, password string, balance float64) *User {
	return &User{
		ID:       Id,
		Name:     name,
		Email:    email,
		Password: password,
		Balance:  balance,
	}
}
