export enum AreasToVolunteer {
    None = 0,
    EntranceGate = 1,
    GiftShop = 2,
    PaintingDecorating = 3
}

export interface FullName {
    firstName: string
    lastName: string
}

export interface Date {
    month: number
    day: number
    year: number
}

export interface Subscription {
    LastPaid: Date
}

export interface User {
    name: FullName
    signUpDate: string
    paying: boolean
    volunteer: boolean
    areaToVolunteer: AreasToVolunteer
    subscirption: Subscription
}
