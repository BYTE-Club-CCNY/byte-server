import express from "express";
import activateDb from "./db";
import { readFile } from "fs";
import winston from "winston";

// activateDb();
const app = express();
const PORT = 3000;

const logger = winston.createLogger({
    // Log only if level is less than (meaning more severe) or equal to this
    level: "info",
    // Use timestamp and printf to create a standard log format
    format: winston.format.combine(
        winston.format.timestamp(),
        winston.format.printf(
            (info) => `${info.timestamp} ${info.level}: ${info.message}`
        )
    ),
    // Log to the console and a file
    transports: [
        new winston.transports.Console(),
        new winston.transports.File({ filename: "logs/app.log" }),
    ],
});

app.use((req, res, next) => {
    // Log an info message for each incoming request
    logger.info(`Received a ${req.method} request for ${req.url}`);
    next();
});

app.get("/", (req, res) => {
    res.send("BYTE @ CCNY").status(200);
});

app.get('/projects', (req, res) => {
    if (req.query) {
        logger.warn("Query parameters ignored");
    }

    readFile('./data.json', 'utf8', (error, content) => {
        if (error) {
            logger.error("Error reading data.json");
            return res.status(500).send("Error reading file");
        }
        return res.status(200).json(JSON.parse(content));
    })
})

app.get('/getProjectByTeam', (req, res) => {
    if (!req.query.team) {
        logger.error("Team query parameter missing");
        res.status(400).send("Missing team");
        return
    }
    readFile('./data.json', 'utf8', (error, content) => {
        if (error) {
            logger.error("Error reading data.json");
            return res.status(500).send("Error reading file");
        }
        const jsonData = JSON.parse(content);
        const filteredData = jsonData.filter((item: any) => {
            return item.team.toString().toLowerCase().split(',').includes(req.query.team?.toString().toLowerCase());
        })
        if (filteredData.length === 0) {
            logger.warn("No projects found");
            return res.status(404).send("The data you are looking for does not exist");
        }
        return res.status(200).send(filteredData);
    })
})

app.get("/getProjectByCohort", (req, res) => {
    if (!req.query.cohort) {
        logger.error("Cohort query parameter missing");
        res.send("Missing cohort").status(400);
        return;
    }
    readFile("data.json", "utf8", (err, data) => {
        if (err) {
            logger.error("Error reading data.json");
            res.send("Error reading file").status(500);
        } else {
            try {
                const jsonData = JSON.parse(data);
                const filteredData = jsonData.filter(
                    (item: any) =>
                        item.cohort.toLowerCase() ===
                        req.query.cohort?.toString().toLowerCase()
                );
                if (filteredData.length === 0) {
                    logger.warn("No projects found");
                    res.send("No projects found").status(404);
                    return;
                }
                res.send(filteredData).status(200);
            } catch (err) {
                logger.error("Error parsing JSON");
                res.send("Error parsing JSON").status(500);
            }
        }
    });
});

app.get("/getProjectByName", (req, res) => {
    if (!req.query.name) {
        logger.error("Name query parameter missing");
        res.send("Missing project name").status(400);
        return;
    }

    readFile("data.json", "utf8", (err, data) => {
        if (err) {
            logger.error("Error reading data.json");
            res.send("Error reading file").status(500);
        } else {
            try {
                const jsonData = JSON.parse(data);
                const filteredData = jsonData.filter((item: any) =>
                    item.name
                        .toLowerCase()
                        .includes(req.query.name?.toString().toLowerCase())
                );
                if (filteredData.length === 0) {
                    logger.warn("No projects found");
                    res.send("No projects found").status(404);
                    return;
                }
                res.send(filteredData).status(200);
            } catch (err) {
                logger.error("Error parsing JSON");
                res.send("Error parsing JSON").status(500);
            }
        }
    });
});

app.listen(PORT, () => {
    console.log(`listening on port ${PORT}`);
});

export default app;
