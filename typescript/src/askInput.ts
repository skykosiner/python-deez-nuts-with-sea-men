const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
})

function checkInput(input: string): boolean {
    return Boolean(input.split("").
        map(x => x.charCodeAt(0)).
        filter(x => x < 32 || x > 127).length);
}

export default function askForInput(question: string): Promise<string> {
    while (true) {
        //@ts-ignore
        return new Promise((resolve, reject) => {
            readline.question(`${question}\n`, (answer: string) => {
                const valid = checkInput(answer);

                if (valid) {
                    return resolve(answer)
                }
            });
        });
    }
};
