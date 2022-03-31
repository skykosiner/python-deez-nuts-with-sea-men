export function currentDate(): string {
    const date: Date = new Date();

    //DD-MM-YY format
    return `${date.getDate()}/${date.getMonth() + 1}/${date.getFullYear()}`
};
