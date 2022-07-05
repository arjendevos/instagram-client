package models

type Count struct {
	Post      int `json:"post"`
	Followers int `json:"followers"`
	Followed  int `json:"Following"`
}

type Business struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Category string `json:"category"`
}

type User struct {
	ID                    string   `json:"id"`
	IsBusinessAccount     bool     `json:"is_business_account"`
	IsProfessionalAccount bool     `json:"is_professional_account"`
	IsVerfied             bool     `json:"is_verified"`
	Username              string   `json:"username"`
	FullName              string   `json:"full_name"`
	ProfilePicture        string   `json:"profile_picture"`
	Count                 Count    `json:"count"`
	ExternalUrl           string   `json:"external_url"`
	Business              Business `json:"business"`
	Category              string   `json:"category_name"`
	Biography             string   `json:"biography"`
	Email                 string   `json:"email"`
	EmailProvider         string   `json:"email_provider"`
}
