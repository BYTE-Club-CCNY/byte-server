import {Router} from "express";
import connectDB from "../connectDB";

const router: Router = Router();

router.use(connectDB);

router.get("/", async (req: any, res: any) => {
    try {
        await req.client.query('SELECT * FROM projects');
        return res.status(200).send("No errors!!");
    } catch (err: any) {
        return res.status(500).send(err.message);
    }
});

export default router;