import client from "./db.config";

async function getDB() {
    await client.connect();
}

async function checkDB(timeout: number) {
    setTimeout(async () => {
        try {
            await getDB();
            process.exit(0);
        } catch (e: any) {
            console.error("Error Connecting to DB:", e.message);
            process.exit(1);
        }
    }, timeout);
}

const args: string[] = process.argv;
const timeout: number = parseInt(args[2]);

await checkDB(timeout);
