#!/usr/bin/env python3

from src import User

user: User = User.new_user()
print("User", user)
user.save_to_json(user)
