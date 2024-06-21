import { spawn } from "child_process";

export function secondsToMs(d: number) {
    return d * 1000;
}

/**
 * @param TIMEOUT seconds till return false
 * @returns true if DB avail, false otherwise
 */
async function checkDB(TIMEOUT: number): Promise<boolean> {
    return new Promise((resolve, reject) => {
        let dbAval: boolean = false;

        const database = spawn("bun", ["dbCheck.ts", TIMEOUT.toString()]);

        database.stdout.on("data", (data) => {
            console.log("Output from dbCheck.ts:", data.toString());
        });

        database.on("exit", (code) => {
            if (code === 0) {
                dbAval = true;
            } else {
                dbAval = false;
            }
            resolve(dbAval);
        });

        database.on("error", (error) => {
            console.error(error);
            reject(error);
        });
    });
}

export default checkDB;
