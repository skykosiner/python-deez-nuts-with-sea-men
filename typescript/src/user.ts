import { appendFile, readFile } from "fs";

export enum AreasToVolunteer {
    entrance_gate = 0,
    gift_shop = 1,
    painting_decorating = 2,
}

export interface fullName {
    firstName: string,
    lastName: string
}

export interface IUser {
    name: fullName,
    signUpDate: string,
    paying: boolean,
    volunteer: boolean
    volunteeringAt: AreasToVolunteer
}

export default class User implements IUser {
    public users: IUser[] = [];
    public volunteeringAt: AreasToVolunteer

    constructor(public name: fullName, public signUpDate: string, public paying: boolean, public volunteer: boolean) {
        this.name = name
        this.signUpDate = signUpDate
        this.paying = paying
        this.volunteer = volunteer

        if (this.volunteer) {
            // Select a random choice from the enum
            const enumValues = Object.keys(AreasToVolunteer)
                .map(n => Number.parseInt(n))
                .filter(n => !Number.isNaN(n))

            const randomIndex = Math.floor(Math.random() * enumValues.length)
            const randomEnumValue = enumValues[randomIndex]
            this.volunteeringAt = randomEnumValue
        }
    }

    public async addUser(user: IUser): Promise<string | void> {
        await this.getUserFile();
        if (this.users.includes(user)) return "Sorry that user is allready signed up";

        this.users.push(user);

        appendFile("./users.json", `${user}\n`, function(err: Error) {
            if (err) throw err;
            return;
        })

    }

    private getUserFile() {
        return new Promise((res, rej) => {
            readFile("./users.json", (err: Error | undefined, data: Buffer) => {
                if (err) {
                    console.log(err);
                    rej(err);
                } else {
                    const arr = data.toString().split("\n")
                    for (const a in arr) {
                        if (arr[a] !== "") this.users.push(arr[a]);
                    };
                    res(null);
                }
            });
        });
    }
}
