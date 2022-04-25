import askForInput from "src/askInput";
import { currentDate } from "src/getCurrentDate";
import { getAreareToVolunter, stringToBoolean } from "../utils";
import { AreasToVolunteer, FullName, Date, Subscription, User } from "./types";

export class UserClass {
    public async newUser(): Promise<User> {
        const firstName: string = await askForInput("What is your first name?");
        const lastName: string = await askForInput("What is your last name?");

        const paying: boolean = stringToBoolean(await askForInput("Would you like to pay now? (yes/no)"));
        const volunteer: boolean = stringToBoolean(await askForInput("Would you like to volunteer? (yes/no)"));

        const fullname: FullName = { firstName, lastName }

        const user: User = {
            name: fullname,
            signUpDate: currentDate(),
            paying,
            volunteer,
            areaToVolunteer: getAreareToVolunter(volunteer)
        }
    }
}
