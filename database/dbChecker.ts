import getDB from "./db";

async function checkDB() {
    const db = await getDB();

    if (db) {
        await db.end();
        return true;
    }
    return false;
}

export default checkDB;
