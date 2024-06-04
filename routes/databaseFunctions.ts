import { Client } from "pg";

// todo: wrap everything in a try catch block and pass the exception to the user
export function addToDB(client: Client, valuesToQuery: any) {
  const query = `
        INSERT INTO projects (name, "short-desc", "long-desc", team, link, image, "tech-stack", cohort, topic)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`;

  const values = [
    valuesToQuery.name,
    valuesToQuery.shortDescription,
    valuesToQuery.longDescription,
    valuesToQuery.team,
    valuesToQuery.link,
    valuesToQuery.image,
    valuesToQuery.techStack,
    valuesToQuery.cohort,
    valuesToQuery.topic,
  ];
  client.query(query, values);
}
