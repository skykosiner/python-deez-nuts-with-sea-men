const readline = require('readline').createInterface({
  input: process.stdin,
  output: process.stdout
})


export default function askInput(question: string): Promise<string> {
    //@ts-ignore
    return new Promise((resolve, reject) => {
        readline.question(question, (answer: string) => {
            resolve(answer)
        });
    });
};
