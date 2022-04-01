from dataclasses import dataclass, asdict
from enum import IntEnum
import random
import json

class areasToVolunteer(IntEnum):
    unspecified = 0
    entrance_gate = 1
    gift_shop = 2
    painting_decorating = 3

@dataclass
class fullName:
    firstName: str
    lastName: str

@dataclass
class User:
    name: fullName
    signUpDate: str
    paying: bool
    volunteer: bool
    areaVolunteering: areasToVolunteer

    def __init__(self, *, name: fullName, signUpDate: str, paying: bool, volunteer: bool):
        self.name = name
        self.signUpDate = signUpDate
        self.paying = paying
        self.volunteer = volunteer

        # If the use choses to volunteer get a random job for them to do, and
        # add it to there user object
        if self.volunteer:
            self.areaVolunteering = random.choice(list(areasToVolunteer)[1:])
        else:
            # If the user is not volunteering then make sure areasToVolunteer is asgined to the unspecified option
            self.areaVolunteering = areasToVolunteer.unspecified

    def saveUser(self, user):
        # Convert the user object to a json string
        json_strings = json.dumps(asdict(user))
        # Dump that bad boy into the json file "users.json"
        f = open("./users.json", "a")
        f.write(json_strings)
        f.close()

    def getUsers(self):
        json_string = open("./users.json", "r")
        json_object = json.loads(json_string.read())
        print("That json object gurlll", json_object)

        # use ** magic to pass all keyword-arguments immediately, way smarter
        print("user", User(**json_object))
        back_from_json: User = User(**json_object)
        print(back_from_json)
