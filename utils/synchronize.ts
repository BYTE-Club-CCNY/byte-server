import { writeFile } from 'fs/promises';
import { Client } from 'pg';
import { queryDatabase } from '../routes/databaseFunctions';
import path from 'path';

const FILE_PATH = path.resolve(__dirname, 'data.json');

const synchronizeLocal = async (client: Client) => {
    const query = 'SELECT * FROM projects';
    try {
        const data = await queryDatabase(client, query, []);
        await writeFile(FILE_PATH, JSON.stringify(data.rows, null, 2));
    }
    catch (e: any) {
        throw Error(e);
    }
}

export default synchronizeLocal;