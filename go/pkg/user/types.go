package user

type AreasToVolunteer int

const (
	None AreasToVolunteer = iota
	EntranceGate
	GiftShop
	PaintingDecorating
)

type FullName struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Date struct {
	Month int `json:"month"`
	Day   int `json:"day"`
	Year  int `json:"year"`
}

type Subscription struct {
	LastPaid Date `json:"last_paid"`
}

type User struct {
	Name            FullName         `json:"name"`
	SignUpDate      string           `json:"signUpDate"`
	Paying          bool             `json:"paying"`
	Volunteer       bool             `json:"volunteer"`
	AreaToVolunteer AreasToVolunteer `json:"AreaToVolunteer"`
	Subscription    Subscription     `json:"subscription"`
}
