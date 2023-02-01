package apiclient

type Users []User

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}

func (u Users) FilterByGender(gender string) map[string][]User {
	var users map[string][]User = map[string][]User{
		"male":   make([]User, 0),
		"female": make([]User, 0),
	}

	if len(u) != 0 {
		for _, u := range u {
			users[gender] = append(users[gender], u)
		}
		return users
	}
	return nil
}
