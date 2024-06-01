import activateDb from "./db";
import logger from "./utils/logger";

const express = require("express");
const app = express();
const PORT = 3000;
const projectsLocal = require("./routes/projectsLocal");
const projectsDB = require("./routes/projectsDB");

let isActive = false;

const connectDB = async (req: any, res: any, next: any) => {
    try {
        if (isActive) {
            logger.info("Database already connected");
            return next();
        }
        await activateDb();
        isActive = true;
        next();
    } catch (err: any) {
        logger.info(err.message);
        next("route");
    }
}

app.use((req: any, res: any, next: any) => {
    // Log an info message for each incoming request
    logger.info(`Received a ${req.method} request for ${req.url}`);
    next();
});

app.get("/", (req: any, res: any) => {
    res.send("BYTE @ CCNY").status(200);
});

app.use("/projects", connectDB, projectsDB);

app.use("/projects", projectsLocal);

app.listen(PORT, () => {
    console.log(`listening on port ${PORT}`);
});

export default app;
