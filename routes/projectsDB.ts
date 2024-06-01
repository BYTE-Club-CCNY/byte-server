import client from "../db.config";

const express = require("express");
const router = express.Router();

router.get("/", (req: any, res: any) => {
    res.status(200).send("Projects Database");
});

module.exports = router;