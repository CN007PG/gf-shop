package model

type UserCreateInput struct {
	Phone    string
	Password string
	Nickname string
}
type UserChangePasswordInput struct {
	Password string
	Nickname string
}
type UserSignInInput struct {
	Nickname string
	Password string
}

type UserUpdateInput struct {
	Name       string
	Password   string
	Phone      string
	BuyHistory string
	BuyCart    string
	Token      string
}
