import {Router} from "express";
import connectDB from "../connectDB";
import express from "express";
import logger from "../utils/logger";

const router: Router = Router();

router.use(connectDB);
router.use(express.json());

router.get("/", async (req: any, res: any) => {
    try {
        await req.client.query('SELECT * FROM projects');
        return res.status(200).send("No errors!!");
    } catch (err: any) {
        return res.status(500).send(err.message);
    }
});

router.get("/get", (req: any, res: any) => {
    req.client.query('SELECT * FROM projects', (err: any, result: any) => {
        if (err) {
            logger.error("Error reading from database");
            return res.status(500).send(err.message);
        }
        return res.status(200).send(result.rows);
    });
});

router.post("/add" , (req: any, res: any) => {
    const keys = Object.keys(req.body);
    const values = Object.values(req.body);
    
    const query = `INSERT INTO projects (${keys.map(key => {
        if (key.includes("-")) {
            return `"${key}"`;
        }
        return key;
    }).join(", ")}) VALUES (${keys.map((key, i) => `$${i + 1}`).join(", ")})`;

    req.client.query(query, values, (err: any, result: any) => {
        if (err) {
            logger.error("Error adding to database");
            return res.status(500).send(err.message);
        }
        return res.status(200).send("Added new project");
    });
});

export default router;