import express from "express";
import activateDb from "./db";
import logger from "./utils/logger";

// activateDb();
const app = express();
const PORT = 3000;
const projectsRoute = require("./routes/projects");

app.use((req, res, next) => {
    // Log an info message for each incoming request
    logger.info(`Received a ${req.method} request for ${req.url}`);
    next();
});

app.get("/", (req, res) => {
    res.send("BYTE @ CCNY").status(200);
});

app.use("/projects", projectsRoute);

app.listen(PORT, () => {
    console.log(`listening on port ${PORT}`);
});

export default app;
