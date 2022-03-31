from dataclasses import dataclass
from enum import IntEnum
import random

class areasToVolunteer(IntEnum):
    entrance_gate = 1
    gift_shop = 2
    painting_decorating = 3

@dataclass
class fullName:
    firstName: str
    lastName: str

@dataclass
class User:
    def __init__(self, name: fullName, signUpDate: str, paying: bool, volunteer: bool):
        self.name = name
        self.signUpDate = signUpDate
        self.paying = paying
        self.volunteer = volunteer

        # If the use choses to volunteer get a random job for them to do, and
        # add it to there user object
        if self.volunteer:
            self.areaVolunteering = random.choice(list(areasToVolunteer))
