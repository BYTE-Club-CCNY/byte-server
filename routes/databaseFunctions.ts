import { Client } from "pg";

/**
 * Function to add all items from values to database
 * Assumes values array correctly maps to the database schema (no empty values, etc.)
 */
export function queryDatabase(client: Client, query: string, values: Array<any>) {
    try {
        return client.query(query, values);
    } catch (e: any) {
        throw Error(e);
    }
}
