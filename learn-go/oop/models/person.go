package models

type GetNamePerson struct {
	Name string `json:"name"`
}

func SetName(name string) *GetNamePerson {
	return &GetNamePerson{
		Name: name,
	}
}

type Person struct {
	GetNamePerson
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	setTime
}

func NewPerson(name, username, password, email string) *Person {
	return &Person{
		GetNamePerson: *SetName(name),
		Username:      username,
		Password:      password,
		Email:         email,
		setTime:       *NewSetTime(),
	}
}
