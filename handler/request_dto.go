package handler

type createUserRequestDto struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
}
