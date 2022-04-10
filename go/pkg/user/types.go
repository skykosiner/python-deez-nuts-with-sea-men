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
    LastName string `json:"lastName"`
}

type User struct {
    Name FullName `json:"name"`
    SignUpDate string `json:"signUpDate"`
    Paying bool `json:"paying"`
    Volunteer bool `json:"volunteer"`
    AreaToVolunteer AreasToVolunteer `json:"AreaToVolunteer"`
}

