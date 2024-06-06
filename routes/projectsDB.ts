import { Router } from "express";
import connectDB from "../connectDB";
import express from "express";
import logger from "../utils/logger";
import { queryDatabase } from "./databaseFunctions";
import { QueryResult } from "pg";

const router: Router = Router();

router.use(connectDB);
router.use(express.json());

router.get("/", async (req: any, res: any) => {
  try {
    return res.status(200).json({"message": "API is operational."});
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
    const data: QueryResult = await queryDatabase(req.client, baseQuery, values);
    return res.status(200).send(data.rows);
  } catch {
    return res.status(500).json({"message": "Error retrieving data"});
  }
});

router.post("/add", (req: any, res: any) => {
  const values: Array<any> = Object.values(req.body);
  const query = `
  INSERT INTO projects (name, "short-desc", "long-desc", team, link, image, "tech-stack", cohort, topic)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`;
  try {
    queryDatabase(req.client, query, values);
    return res.status(200).json({"message": "Project added successfully"});
  } catch (err: any) {
    return res.status(400).json({"message": err.message});
  }
});

router.put("/update", async (req: any, res: any) => {
  const projectName = req.query.name;

  if (!projectName) {
    return res.status(400).json({ message: "Project name is required" });
  }

  const fields = req.body;
  if (!fields || Object.keys(fields).length === 0) {
    return res.status(400).json({ message: "No fields to update provided" });
  }

  const setClauses: string[] = [];
  const values: (string | number)[] = [];

  // Construct the set clauses and values array
  Object.keys(fields).forEach((key, index) => {
    setClauses.push(`"${key}" = $${index + 1}`);
    values.push(fields[key]);
  });

  // Add the project name to the values array for the WHERE clause
  values.push(projectName);

  const query = `
    UPDATE projects
    SET ${setClauses.join(", ")}
    WHERE name = $${values.length}`;
  console.log(query);
  console.log(values);
  try {
    await queryDatabase(req.client, query, values);
    return res.status(200).send("Project updated successfully");
  } catch (err: any) {
    return res.status(400).json({ message: err.message });
  }
});


export default router;
