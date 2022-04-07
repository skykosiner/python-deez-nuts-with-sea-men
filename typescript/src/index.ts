import { isFunction } from "util";
import { isDate } from "util/types";
import askInput from "./askInput";
import { currentDate } from "./getCurrentDate";
import User, { fullName } from "./user";

async function main() {
    const firstName = await askInput("What is your first name? ");
    const lastName = await askInput("What is your last name? ");

    const volunteer = await askInput("Do you want to volunteer (true for yes, and false for no)? ");
    const pay = await askInput("Pay the $75 fee (true for yes, and false for no)? ");

    const fullName: fullName = {
        firstName,
        lastName,
    };

    const user = new User(fullName, currentDate(), pay === "true" ? true : false, volunteer === "true" ? true : false);
    user.addUser(user);
}

main();
