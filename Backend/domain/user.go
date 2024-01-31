package domain

type User struct {
	BaseModel
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Image       string `json:"image"`
	Bio         string `json:"bio"`
}
