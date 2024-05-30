import express from "express";
import activateDb from "./db";
import logger from "./utils/logger";

// activateDb();
const app = express();
const PORT = 3000;
const projectsLocal = require("./routes/projectsLocal");
const projectsDB = require("./routes/projectsDB");

const connectDB = async (req: any, res: any, next: any) => {
    try {
        await activateDb();
        next();
    } catch (err: any) {
        logger.info(`Error connecting to database: ${err.message} \n Switching to local data.`);
        next("route");
    }
}

app.use((req, res, next) => {
    // Log an info message for each incoming request
    logger.info(`Received a ${req.method} request for ${req.url}`);
    next();
});

app.get("/", (req, res) => {
    res.send("BYTE @ CCNY").status(200);
});

app.use("/projects", connectDB, projectsDB);

app.use("/projects", projectsLocal);

app.listen(PORT, () => {
    console.log(`listening on port ${PORT}`);
});

export default app;
