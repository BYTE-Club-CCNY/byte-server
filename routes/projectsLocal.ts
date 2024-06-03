import logger from "../utils/logger";
import { Router } from "express";
import { readFile } from "fs";
import path from "path";

const router = Router();
const FILE_PATH: string = path.resolve(__dirname, "../data.json");

router.get("/", (req: any, res: any) => {
    if (req.query) {
        logger.warn("Query parameters ignored");
    }
    
    readFile(FILE_PATH, "utf8", (error: any, content: any) => {
        if (error) {
            logger.error("Error reading data.json");
            return res.status(500).send("Error reading file");
        }
        return res.status(200).json(JSON.parse(content));
    });
});

router.get("/team", (req: any, res: any) => {
    if (!req.query.team) {
        logger.error("Team query parameter missing");
        res.status(400).send("Missing team");
        return;
    }

    readFile(FILE_PATH, "utf8", (error: any, content: any) => {
        if (error) {
            logger.error("Error reading data.json");
            return res.status(500).send("Error reading file");
        }
        const jsonData = JSON.parse(content);
        const filteredData = jsonData.filter((item: any) => {
            const itemData = item.team.toString().toLowerCase().split(",");
            const queryData = req.query.team
                ?.toString()
                .toLowerCase()
                .split(",");
            return queryData?.every((query: any) => itemData.includes(query));
        });
        if (filteredData.length === 0) {
            logger.warn("No projects found");
            return res
                .status(404)
                .send("The data you are looking for does not exist");
        }
        return res.status(200).send(filteredData);
    });
});

router.get("/cohort", (req: any, res: any) => {
    if (!req.query.cohort) {
        logger.error("Cohort query parameter missing");
        res.send("Missing cohort").status(400);
        return;
    }

    readFile(FILE_PATH, "utf8", (err: any, data: any) => {
        if (err) {
            logger.error("Error reading data.json");
            res.send("Error reading file").status(500);
        }
        const jsonData = JSON.parse(data);
        const filteredData = jsonData.filter((item: any) => {
            const itemData = item.cohort.toString().toLowerCase().split(",");
            const queryData = req.query.cohort
                ?.toString()
                .toLowerCase()
                .split(",");
            console.log(itemData, queryData);
            return itemData.some((item: any) => queryData?.includes(item));
        });

        if (filteredData.length === 0) {
            logger.warn("No projects found");
            res.send("No projects found").status(404);
            return;
        }
        res.send(filteredData).status(200);
    });
});

router.get("/name", (req: any, res: any) => {
    if (!req.query.name) {
        logger.error("Name query parameter missing");
        res.send("Missing project name").status(400);
        return;
    }

    readFile(FILE_PATH, "utf8", (err: any, data: any) => {
        if (err) {
            logger.error("Error reading data.json");
            res.send("Error reading file").status(500);
        }
        const jsonData = JSON.parse(data);
        const filteredData = jsonData.filter((item: any) => {
            const itemData = item.name.toString().toLowerCase();
            const queryData = req.query.name
                ?.toString()
                .toLowerCase()
                .split(",");
            return queryData?.some((query: any) => itemData.includes(query));
        });
        if (filteredData.length === 0) {
            logger.warn("No projects found");
            res.send("No projects found").status(404);
            return;
        }
        res.send(filteredData).status(200);
    });
});

export default router;