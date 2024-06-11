import dotenv from "dotenv";
import { Client } from "pg";
dotenv.config();

const client = new Client({
    host: process.env.POSTGRESQL_DB_HOST,
    user: process.env.POSTGRESQL_DB_USER,
    password: process.env.POSTGRESQL_DB_PASSWORD,
    database: process.env.POSTGRESQL_DB,
    port: process.env.POSTGRESQL_DB_PORT
        ? parseInt(process.env.POSTGRESQL_DB_PORT)
        : 5432,
});

client.on("end", () => console.log("Client has disconnected"));
client.on("error", (err) =>
    console.error("Unexpected error on idle client", err),
);

export default client;
