import { AreasToVolunteer } from "./user/types"

export function stringToBoolean(value: string): boolean {
    switch (value) {
        case "yes" || "y":
            return true
        case "no" || "n":
            return false
        default:
            return false
    }
}

export function getAreareToVolunter(volunteer: boolean): AreasToVolunteer {
    if (volunteer) {
        const index: number = Math.floor(Math.random() * Object.keys(AreasToVolunteer).length);
        const value = Object.values(AreasToVolunteer)[index];

        return AreasToVolunteer[value];
    }

    return AreasToVolunteer.None
}

function test(str: string): boolean {
    return Boolean(str.split("").
        map((x) => x.charCodeAt(0)).
        filter(x => x < 32 || x > 27).length);
}
