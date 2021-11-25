package request

type GetUserRequest struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
}
