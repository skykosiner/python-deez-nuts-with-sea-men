#!/usr/bin/env python3

from src import User, Members, areasToVolunteer, Sponser

while True:
    inputInfo = input("""S to sign up
V to see all memmbers who are working as volunteers
E for members who are volunteering at the entrance gate
G for members who are volunteering at the gift shop
P for members who are helping with painting and dectoring
$ For members who have not yet paid there $75 fee
£ For members who subscription has expried and not yet paied for the year
N To become a new sponser, and also learn more about becoming a sponser
L to list all sponsers
C to exit
""")

    members = Members()

    match inputInfo:
        case "S":
            user = User.new_user()
            User.save_to_json(user)
        case "V":
            members.get_volunteering_members()
        case "E":
            members.get_area_of_volunteer(areasToVolunteer.entrance_gate)
        case "G":
            members.get_area_of_volunteer(areasToVolunteer.gift_shop)
        case "P":
            members.get_area_of_volunteer(areasToVolunteer.painting_decorating)
        case "$":
            members.get_non_payed_members(),
        case "£":
            members.subscription_expired(),
        case "N":
            sponser = Sponser.new_sponsor()
            Sponser.save_sponsor(sponser)
        case "L":
            Sponser.get_sponsers()
        case "C":
            break
        case _:
            print("Sorry you have entered an invalid option")
