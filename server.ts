import { Client } from "pg";
import activateDb from "./db";
import logger from "./utils/logger";
import projectsDB from "./routes/projectsDB";
import express from "express";

const app = express();
const PORT = 3000;

let client: Client | undefined;

const connectDB = async (req: any, res: any, next: any) => {
    try {
        client = await activateDb();
        await client.end();
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


// if (!client) {
//     app.use("/projects", projectsLocal);
// } else {
    app.use("/projects", connectDB, projectsDB);
// }


app.listen(PORT, () => {
    console.log(`listening on port ${PORT}`);
});

export default app;
