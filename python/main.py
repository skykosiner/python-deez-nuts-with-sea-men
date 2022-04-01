#!/usr/bin/env python3
from dataclasses import dataclass
from src import User, fullName, getCurrentDate
from typing import Union

def validateName(s: str) -> bool:
    return len(s) > 1

def validateBool(s: str) -> bool:
    return s.lower() in {"yes", "no", "y", "n"}

def ask_for_input(message: str, valid):
    while True:
        response = input(f"{message}\n")
        if valid(response):
            break
        print("Invalid input, please answer again!")
    return response

def ask_for_inputs():
    firstName: str = ask_for_input(
        "What is your first name?",
        validateName,
    )
    lastName: str = ask_for_input(
        "What is your last name?",
        validateName,
    )
    volunteer: bool = ask_for_input(
        "Would you like to volunteer? (yes/no)",
        validateBool
    ).lower() in "yes"
    pay: bool = ask_for_input(
        "Would you like to pay right now? (yes/no)",
        validateBool
    ).lower() in "yes"

    fullname = fullName(firstName, lastName)
    user = User(name=fullname, signUpDate=getCurrentDate(), volunteer=volunteer, paying=pay)
    print(user.signUpDate)
    print("user", user)
    return User(name=fullname, signUpDate=getCurrentDate(), volunteer=volunteer, paying=pay)

user: User = ask_for_inputs()
user.saveUser(user)
