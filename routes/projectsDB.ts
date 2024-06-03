import { Client } from "pg";
import activateDb from "../db";
import express from "express";
let router = express.Router();
let client: Client | undefined;

router.get("/", async (req: any, res: any) => {
    try {
        client = await activateDb();
        await client.end();
    } catch (err: any) {
        return res.status(500).send(err.message);
    }
    return res.status(200).send("Projects Database");
});

export default router;