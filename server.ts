import { Client } from "pg";
import activateDb from "./db";
import logger from "./utils/logger";
import projectsDB from "./routes/projectsDB";
import express from "express";
import connectDB from "./connectDB";
import checkDB, { secondsToMs } from "./dbChecker";

const app = express();
const PORT = 3000;
const INTERVAL = secondsToMs(5);
const TIMEOUT = secondsToMs(2);
let dbAval: boolean = false;

setInterval(async () => {
    try {
        dbAval = await checkDB(TIMEOUT);
    } catch (e: any) {
        console.error("Error:", e.message);
        dbAval = false;
    }
}, INTERVAL);

app.use((req: any, res: any, next: any) => {
    logger.info(`Received a ${req.method} request for ${req.url}`);
    next();
});

app.get("/", (req: any, res: any) => {
    res.send(
        `BYTE @ CCNY. The database is ${dbAval ? "available" : "not available"}`,
    ).status(200);
});

// app.use("/projects", projectsLocal);
// app.use("/projects", projectsDB);

// // any other route will return a 404
// app.get("*", (req: any, res: any) => {
//     res.status(404).json({ message: "Page not found. Invalid path or method provided to make this request." });
// });

app.listen(PORT, () => {
    console.log(`listening on port ${PORT}`);
});

export default app;
