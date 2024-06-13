import client from "./db.config";

async function getDB(): Promise<void> {
	await client.connect();
	await client.end()
}

function runWithTimeout(fn: () => Promise<void>, timeout: number, interval: number) : Promise<void> {
    return new Promise<void>((resolve, reject) => {
        const startTime = Date.now();

        const handle = setInterval(async () => {
            const elapsedTime = Date.now() - startTime;

            if (elapsedTime >= timeout) {
                clearInterval(handle);
                resolve();
            } else {
                try {
                    await fn();
                    clearInterval(handle);
                    resolve();
                } catch (error) {
                    clearInterval(handle);
                    reject(error);
                }
            }
        }, interval);
    });
}

runWithTimeout(getDB, parseInt(process.argv[2]), parseInt(process.argv[2]))
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
