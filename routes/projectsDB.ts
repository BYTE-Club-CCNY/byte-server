import { Client } from "pg";
import activateDb from "../db";
import express from "express";
const router = express.Router();

let client: Client | undefined;

router.get("/", async (req: any, res: any) => {
    try {
        client = await activateDb();
        await client.end();
        return res.status(200).send("No errors!!");
    } catch (err: any) {
        return res.status(500).send(err.message);
    }
});

export default router;