import dotenv from "dotenv";
import { Client } from "pg";
import logger from "utils/logger";
dotenv.config();

const secondsToMs = (seconds: number): number => {
    return seconds * 1000;
};

const client = new Client({
    host: process.env.POSTGRESQL_DB_HOST,
    user: process.env.POSTGRESQL_DB_USER,
    password: process.env.POSTGRESQL_DB_PASSWORD,
    database: process.env.POSTGRESQL_DB,
    port: process.env.POSTGRESQL_DB_PORT
        ? parseInt(process.env.POSTGRESQL_DB_PORT)
        : 5432,
    statement_timeout: secondsToMs(5),
    query_timeout: secondsToMs(5),
    application_name: "BYTE Server",
    connectionTimeoutMillis: secondsToMs(5),
    idle_in_transaction_session_timeout: secondsToMs(5),
});

client.on("end", () => console.log("Client has disconnected"));
client.on("error", (err) =>
    console.error("Unexpected error on idle client", err),
);
client.on("notification", (msg) => {
    logger.alert(msg.channel);
    logger.alert(msg.payload);
});
client.on("notice", (msg) => logger.warn("notice:", msg));

export default client;
