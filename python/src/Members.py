from datetime import datetime
import json
from typing import List
from .User import User, fullName, areasToVolunteer, Subscription
from .Date import Date
from .Utils import get_current_date

class Members:
    def __init__(self):
        self.users = []

        # Each time we run the class we want to have a list of all the users in an array
        self.get_members()

    def get_members(self) -> None:
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

                self.users.append(back_from_json)

    # TODO: Make return type an array of the user object (for all the methods below)
    def get_non_payed_members(self):
        notPaying = []
        for i in self.users:
            if not i.paying:
                fullname = f"{i.name.first_name} {i.name.last_name}"
                notPaying.append(fullname)

        for i in notPaying:
            print(i)

    def get_volunteering_members(self):
        volunteers = []
        for i in self.users:
            if i.volunteer:
                fullname = f"{i.name.first_name} {i.name.last_name}"
                volunteers.append(fullname)

        for i in volunteers:
            print(i)

    # This works with any of the three areas, you just pass in
    # areasToVolunteer.painting_decorating or any other area one can volunteer
    def get_area_of_volunteer(self, area: areasToVolunteer):
        volunters = []

        for i in self.users:
            if i.area_volunteering == area:
                fullname = f"{i.name.first_name} {i.name.last_name}"
                volunters.append(fullname)
        for i in volunters:
            print(i)

    def subscription_expired(self):
        for i in self.users:
            i.subscription_expired()
