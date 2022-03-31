from src import User

fristName: str = input("What is your first name? ")
lastName: str = input("What is your last name? ")

class FullName:
    def __init__(self, firstName: str, lastName: str):
        self.firstName = firstName
        self.lastName = lastName

def makeFullName(firstName: str, lastName: str) -> FullName:
    fullName = FullName(firstName, lastName)
    return fullName

user = User(makeFullName(firstName, lastName), "420", True, True)
