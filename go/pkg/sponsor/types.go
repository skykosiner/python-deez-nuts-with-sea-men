package sponsor

import "github.com/yonikosiner/python-deez-nuts-with-sea-men/pkg/user"

type Sponser struct {
	Name    user.FullName `json:"name"`
	Message string        `json:"message"`
}
