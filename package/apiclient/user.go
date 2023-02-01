package apiclient

type Users []User

type User struct {
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}

func (u Users) FilterByGender(gender string) []User {
	var users []User

	if len(u) != 0 {
		for _, u := range u {
			if u.Gender == gender {
				users = append(users, u)
			}
		}
		return users
	}
	return nil
}
