from dataclasses import dataclass, asdict
import json
from .User import fullName, User

@dataclass
class Sponser:
    name: fullName
    message: str

    @staticmethod
    def new_sponsor():
        print("Have your name and a message of your choise be put on a brass plaque, and also support as with a $200 fee <3")
        first_name = User.get_input("What is your first name?", lambda s: len(s) > 1)
        last_name = User.get_input("What is your last name?", lambda s: len(s) > 1)
        message = User.get_input("What would you like your plank to say? ", lambda s: len(s) <= 128)

        fullname: fullName = fullName(first_name, last_name)

        sponser = Sponser.check_sponsor_dep(Sponser(fullname, message))

        print("Are you sure you want to continue with this info?")
        print(sponser)

        continue_sponser = User.get_input("Continue? (yes/no)", User.validate_bool).lower() in "yes"

        if continue_sponser:
            return sponser

    @staticmethod
    def check_sponsor_dep(sponser: 'Sponser'):
        sponsers = []
        with open("./sponsers.json", "r") as json_lines:
            content = json_lines.readlines()
            for i in content:
                json_object = json.loads(i)
                back_from_json: Sponser = Sponser(**json_object)

                # Make sure the name is formated correctly in the object
                fullname: fullName = fullName(**json_object["name"])
                back_from_json.name = fullname
                sponsers.append(back_from_json)

            for i in sponsers:
                if i.name == sponser.name:
                    print("You're already sponsoring")
                    continue_sponser = User.get_input("Continue as someone else? (yes/no)", User.validate_bool) in "yes"
                    if continue_sponser:
                        return Sponser.new_sponsor()
                    else:
                        exit(0x45)
            return sponser

    @staticmethod
    def save_sponsor(sponser):
        json_string = json.dumps(asdict(sponser))

        print(json_string, file=open("./sponsers.json", "a"))

    @staticmethod
    def get_sponsers():
        sponsers = []
        with open("./sponsers.json", "r") as json_lines:
            content = json_lines.readlines()
            for i in content:
                json_object = json.loads(i)
                back_from_json: Sponser = Sponser(**json_object)

                # Make sure the name is formated correctly in the object
                fullname: fullName = fullName(**json_object["name"])
                back_from_json.name = fullname

                name = f"name: {back_from_json.name.first_name} {back_from_json.name.last_name}\n"
                message = f"message: {back_from_json.message}\n"

                stringToReturn = name + message

                sponsers.append(stringToReturn)

            for i in sponsers:
                print(i)
