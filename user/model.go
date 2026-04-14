package user

import "time"

type User struct {

	// this will be unique
	Email string

	Password string

	// this will be the pk
	Id int

	// could be normal user or employee or manager
	Role string

	Created_at time.Time
}

func (m *User) GetPID() string        { return m.Email }
func (m *User) GetPassword() string   { return m.Password }
func (m *User) PutPID(email string)   { m.Email = email }
func (m *User) PutPassword(pw string) { m.Password = pw }

func NewUser() *User {

	return &User{}
}
