package main

import (
    "github.com/yonikosiner/python-deez-nuts-with-sea-men/pkg/user"
)

func main() {
    var u *user.User

    user := u.NewUser()
    u.SaveUser(user)
}
