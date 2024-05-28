import express from "express";
import activateDb from "./db";

// activateDb();
const app = express();
const PORT = 3000;

app.get("/", (req, res) => {
    res.send("BYTE @ CCNY").status(200);
})

app.listen(PORT, () => {
    console.log(`listening on port ${PORT}`);
});

export default app;