#!/usr/bin/env python3

from src import User

user: User = User.new_user()

user.get_json()
user.save_to_json(user)
