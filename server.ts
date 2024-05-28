import express from "express";
import activateDb from "./db";
import { readFile }  from "fs";


// activateDb();
const app = express();
const PORT = 3000;

app.get("/", (req, res) => {
    res.send("BYTE @ CCNY").status(200);
});

app.get("/getProjectByCohort", (req, res) => {
    if (!req.query.cohort) {
        res.send("Missing cohort").status(400);
        return;
    }
    readFile("data.json", "utf8", (err, data) => {
        if (err) {
            res.send("Error reading file").status(500);
        } else {
            try {
                const jsonData = JSON.parse(data);
                const filteredData = jsonData.filter((item: any) => item.cohort.toLowerCase() === req.query.cohort?.toString().toLowerCase());
                if (filteredData.length === 0) {
                    res.send("No projects found").status(404);
                    return;
                }
                res.send(filteredData).status(200);
            } catch(err) {
                res.send("Error parsing JSON").status(500);
            }
        }
    });
});

app.get("/getProjectByName", (req, res) => {
    if (!req.query.name) {
        res.send("Missing project name").status(400);
        return;
    }

    readFile("data.json", "utf8", (err, data) => {
        if (err) {
            res.send("Error reading file").status(500);
        } else {
            try {
                const jsonData = JSON.parse(data);
                const filteredData = jsonData.filter((item: any) => item.name.toLowerCase().includes(req.query.name?.toString().toLowerCase()));
                if (filteredData.length === 0) {
                    res.send("No projects found").status(404);
                    return;
                }
                res.send(filteredData).status(200);
            } catch(err) {
                res.send("Error parsing JSON").status(500);
            }
        }
    });

});

app.listen(PORT, () => {
    console.log(`listening on port ${PORT}`);
});

export default app;