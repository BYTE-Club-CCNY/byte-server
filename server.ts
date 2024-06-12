import logger from "./utils/logger";
import projectsLocal from "./routes/projectsLocal";
import projectsDB from "./routes/projectsDB";
import express from "express";
import checkDB, { secondsToMs } from "./dbChecker";
import cors from "cors";
import http from 'http';
import https from 'https';

// const privateKey  = fs.readFileSync('sslcert/server.key', 'utf8');
// const certificate = fs.readFileSync('sslcert/server.crt', 'utf8');
// const credentials = {key: privateKey, cert: certificate};

const PORT = 3000;
const INTERVAL = secondsToMs(60 * 60); // 1 hr
const TIMEOUT = secondsToMs(10);
const app = express();
let dbAval: boolean = true;

// initial check
try{
    console.log(`Please wait ${TIMEOUT/1000}s for the database to connect`)
    dbAval = await checkDB(TIMEOUT);
} catch (e: any) {
    dbAval = false;
}

// routine
setInterval(async () => {
    try {
        dbAval = await checkDB(TIMEOUT);
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

// app.listen(PORT, () => {
//     console.log(`listening on port ${PORT}`);
// });

const httpServer = http.createServer(app);
// const httpsServer = https.createServer(credentials, app);

httpServer.listen(PORT);
// httpsServer.listen(PORT+1);