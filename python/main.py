from src import User, fullName

firstName: str = input("What is your first name? ")
lastName: str = input("What is your last name? ")

fullname = fullName(firstName, lastName)

user = User(fullname, "420", True, True)
