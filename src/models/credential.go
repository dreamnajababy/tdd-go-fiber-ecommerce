package models

type Credential struct {
	Username string
	Password string
}

var (
	CorrectUser = Credential{
		Username: "dreamnajababy", Password: "1234",
	}
	IncorrectUser = Credential{
		Username: "dreamnajababy", Password: "@#$%^&*()",
	}
)
