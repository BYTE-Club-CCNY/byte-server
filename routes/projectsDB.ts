import { Client } from "pg";
import activateDb from "../db";
import {Router} from "express";
const router: Router = Router();

router.get("/", async (req: any, res: any) => {
    let client: Client | undefined;

    try{
        client = await activateDb();
        await client.end()
    } catch (e) {
        return res.status(500).send("Internal Server Error");
    }

    return res.status(200).send("Projects Database");
});

export default router;