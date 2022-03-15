package lib

type bizError string

func (b bizError) Error() string {
	return string(b)
}

const (
	ErrUsernameExists              bizError = "Username has been used"
	ErrUsernameOrPasswordIncorrect bizError = "Incorrect username or password"
	ErrUserNotFound                bizError = "user not found"
)
