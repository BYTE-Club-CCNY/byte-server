import cors from "cors";
import express from "express";
import fs from "fs";
import http from "http";
import https from "https";
import { secondsToMs } from "./database/db.config";
import checkDB from "./database/dbChecker";
import projectsDB from "./routes/projectsDB";
import projectsLocal from "./routes/projectsLocal";
import logger from "./utils/logger";

const privateKey = fs.readFileSync("cert/privkey.pem", "utf8");
const certificate = fs.readFileSync("cert/cert.pem", "utf8");
const credentials = { key: privateKey, cert: certificate };

const INTERVAL: number = secondsToMs(60 * 60);
const PORT: number = 3000;
const app = express();
let dbAval: boolean = true;
 
// initial check
(async () => {
    try {
        dbAval = await checkDB();
        logger.info(`Server is up, database is ${dbAval ? "up" : "not up"}`);
    } catch (e: any) {
        dbAval = false;
    }
})();

// routine
setInterval(async () => {
    try {
        dbAval = await checkDB();
    } catch (e: any) {
        console.error("Error:", e.message);
        dbAval = false;
    }
    logger.info(`Database is ${dbAval ? "available" : "not available"}`);
}, INTERVAL);

app.use(cors());
app.use((req: any, res: any, next: any) => {
    logger.info(`Received a ${req.method} request for ${req.url}`);
    next();
});

app.use("/projects", (req: any, res: any, next: any) => {
    if (dbAval) {
        projectsDB(req, res, next);
    } else {
        projectsLocal(req, res, next);
    }
});

app.get("/", (req: any, res: any) => {
    res.send(
        `BYTE @ CCNY. The database is ${dbAval ? "available" : "not available"}`,
    ).status(200);
});

// any other route will return a 404
app.get("*", (req: any, res: any) => {
    res.status(404).json({
        message:
            "Page not found. Invalid path or method provided to make this request.",
    });
});

const httpServer = http.createServer(app);
const httpsServer = https.createServer(credentials, app);

httpServer.listen(PORT);
httpsServer.listen(PORT + 1);
