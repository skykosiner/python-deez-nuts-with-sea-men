from dataclasses import dataclass, asdict
# import datetime
from enum import IntEnum
from random import choice
import json
from .Date import Date

from .Utils import get_current_date, over_a_year

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
class Subscription:
    last_paid: Date

@dataclass
class User:
    # TODO: add in yearly subscription logic
    name: fullName
    sign_up_date: Date
    paying: bool
    volunteer: bool
    area_volunteering: areasToVolunteer
    subscritpion: Subscription

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

        subscription = Subscription(get_current_date())

        # If paying is not true then just say they have never payed even if there is a date there

        user = User(fullname, get_current_date(), pay, volunteer, areaVolunteering, subscription)

        return User.check_user_dup(user)

    @staticmethod
    def check_user_dup(user):
        users = []
        with open("./users.json", "r") as json_lines:
            content = json_lines.readlines()
            for i in content:
                json_object = json.loads(i)
                back_from_json: User = User(**json_object)

                # Make sure the name is formated correctly in the object
                fullname: fullName = fullName(**json_object["name"])
                back_from_json.name = fullname

                subscription: Subscription = Subscription(last_paid=Date(**json_object["subscritpion"]["last_paid"]))
                back_from_json.subscritpion = subscription

                users.append(back_from_json)

        for i in users:
            if i.name == user.name:
                print("Sorry, that users allready exists")
                return User.new_user()
        return user

    @staticmethod
    def save_to_json(user):
        json_string = json.dumps(asdict(user))

        print(json_string, file=open("./users.json", "a"))

    @staticmethod
    def get_working_area(workingAreaValue: areasToVolunteer):
        return areasToVolunteer(workingAreaValue)

    def subscription_expired(self) -> bool:
        # Check if it has been over a year since the user payed
        if not self.subscritpion.last_paid.day == 0:
            print("testing 1234")
            return over_a_year(self.subscritpion.last_paid)
        return False
