package domain

type User struct {
	Id       uint32
	Username string
	Type     UserType
}

type UserType string

const (
	Owner    UserType = "Owner"
	Customer UserType = "Customer"
)
