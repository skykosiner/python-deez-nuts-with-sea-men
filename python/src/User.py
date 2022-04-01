from dataclasses import dataclass, asdict
from enum import IntEnum
from random import choice
import json

from .Utils import getCurrentDate

class areasToVolunteer(IntEnum):
    none = 0
    entrance_gate = 1
    gift_shop = 2
    painting_decorating = 3

@dataclass
class fullName:
    first_name: str
    last_name: str

@dataclass
class User:
    name: fullName
    signUpDate: str
    paying: bool
    volunteer: bool
    areaVolunteering: areasToVolunteer

    @staticmethod
    def validate_name(s: str) -> bool:
        return len(s) > 1

    @staticmethod
    def validate_bool(s: str) -> bool:
        return s.lower() in {"yes", "no", "y", "n"}

    @staticmethod
    def get_input(message: str, valid):
        while True:
            response = input(f"{message}\n")
            if valid(response):
                return response
            print("Invalid input, please answer again!")

    @staticmethod
    def new_user():
        first_name: str = User.get_input("What is your first name?", User.validate_name)
        last_name: str = User.get_input("What is your last name?", User.validate_name)

        fullname = fullName(first_name=first_name, last_name=last_name)

        volunteer: bool = User.get_input("Would you like to volunteer? (yes/no)", User.validate_bool).lower() in "yes"
        pay: bool = User.get_input("Would you like to pay now? (yes/no)", User.validate_bool).lower() in "yes"
        areaVolunteering: areasToVolunteer = choice(list(areasToVolunteer)[1:]) if pay else areasToVolunteer.none

        return User(fullname, getCurrentDate(), pay, volunteer, areaVolunteering)

    @staticmethod
    def save_to_json(user):
        json_string = json.dumps(asdict(user))

        print(json_string, file=open("./users.json", "a"))

    @staticmethod
    def get_json():
        with open("./users.json", "r") as json_lines:
            content = json_lines.readlines()
            if content != None:
                for i in content:
                    json_object = json.loads(i)
                    back_from_json: User = User(**json_object)
                    print(back_from_json)
