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
client.on("error", (err) => console.error("Error:", err));
client.on("notice", (msg) => console.warn("Notice:", msg));
client.on("notification", (msg) => {
    console.log("Notification:", msg);
    console.log("Payload:", msg.payload);
})

export default client;
