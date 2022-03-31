from dataclasses import dataclass
from enum import IntFlag

class areasToVolunteer(IntFlag):
    entrance_gate = 0
    gift_shop = 1
    painting_decorating = 2

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

        if self.volunteer == True:
            # Get a random item from the areas where the user can volunteer
            for i in areasToVolunteer:
                print(i)
            # self.areaVolunteering
