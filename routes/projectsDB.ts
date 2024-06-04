import { Router } from "express";
import connectDB from "../connectDB";
import express from "express";
import logger from "../utils/logger";
import { addToDB, getFromDB } from "./databaseFunctions";
import { QueryResult } from "pg";

const router: Router = Router();

router.use(connectDB);
router.use(express.json());

router.get("/", async (req: any, res: any) => {
  try {
    return res.status(200).send("No errors!!");
  } catch (err: any) {
    return res.status(500).json({ "message": err.message });
  }
});

router.get("/get", async (req: any, res: any) => {
  let baseQuery = "SELECT * FROM projects";
  const filters: string[] = [];
  const values: (string | number)[] = [];

  // if the name filter was provided
  if (req.query.name) {
    filters.push(`name ILIKE $${filters.length + 1}`);
    values.push(`%${req.query.name}%`);
  }

  // if the cohort filter was provided
  if (req.query.cohort) {
    filters.push(`cohort ILIKE $${filters.length + 1}`);
    values.push(`%${req.query.cohort}%`);
  }
  // if the team filter was provided
  if (req.query.team) {
    filters.push(`team ILIKE $${filters.length + 1}`);
    values.push(`%${req.query.team}%`);
  }

  // combine all the filters into a single query
  if (filters.length > 0) {
    baseQuery += " WHERE " + filters.join(" AND ");
  }

  // execute the query, making sure to provide the values for the filters
  try {
    const data: QueryResult = await getFromDB(req.client, baseQuery, values);
    return res.status(200).send(data.rows);
  } catch {
    return res.status(500).json({"message": "Error retrieving data"});
  }
});

router.post("/add", (req: any, res: any) => {
  const values: Array<any> = Object.values(req.body);

  try {
    addToDB(req.client, values);
    return res.status(200).send("Project added successfully");
  } catch (err: any) {
    return res.status(400).send(err.message);
  }
});

router.post("/update", (req: any, res: any) => {
  // TODO: Implement this endpoint

});

export default router;
