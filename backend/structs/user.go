package structs

type User struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	OpenId  string `json:"open_id"`
	IsAdmin bool   `json:"is_admin"`
}

type Review struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
}

type Property struct {
	Address  string   `json:"address"` // TODO: Change this to something we can search easily
	Landlord Landlord `json:"landlord"`
}

type Landlord struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
