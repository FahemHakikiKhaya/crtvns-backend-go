package request

type AuthLoginRequest struct {
	Name     string `validate: "required" json:"name"`
	Email    string `validate: "required,email" json:"email"`
	Password string `validate: "required,min=8" json:"password"`
}