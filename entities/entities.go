package entities

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}
type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}

//TODO: application code
type Todo struct {
	UserID uint64 `json:"user_id"`
	Title  string `json:"title"`
}
