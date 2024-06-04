import { Client } from "pg";

/**
 * Function to add all items from values to database
 * Assumes values array correctly maps to the database schema (no empty values, etc.)
 */
export function addToDB(client: Client, query: string, values: Array<any>) {
    try {
        return client.query(query, values);
    } catch (e: any) {
        throw Error(e);
    }
}

export function getFromDB(client: Client, query: string, values: Array<any>) {
    try {
        return client.query(query, values);
    } catch (e: any) {
        throw Error(e);
    }
}
