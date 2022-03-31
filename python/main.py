from dataclasses import dataclass
from src import User, fullName, getCurrentDate
from typing import Union

def ask_for_input(message: str, valid):
    while True:
        response = input(message)
        if valid(response):
            break
        print("Invalid input, please answer again!")
    return response

def ask_for_inputs():
    # no real validation on first and last name for now...
    firstName: str = ask_for_input(
        "What is your first name? ",
        lambda s: len(s) > 1,
    )

    lastName: str = ask_for_input(
        "What is your last name? ",
        lambda s: len(s) > 1,
    )

    volunteer: bool = ask_for_input(
        "Would you like to volunteer? ",
        lambda s: s.lower() in {"yes", "no", "y", "n"},
    ).lower() in {"y", "yes"}

    pay: bool = ask_for_input(
        "Would you like to pay right now? ",
        lambda s: s.lower() in {"yes", "no", "y", "n"},
    ).lower() in {"y", "yes"}

    fullname = fullName(firstName, lastName)
    user = User(fullname, getCurrentDate(), pay, volunteer)
    return user

user = ask_for_inputs()
