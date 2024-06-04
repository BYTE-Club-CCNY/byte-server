import { Client } from "pg";

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

  try {
    return client.query(query, values);
  } catch (e: any) {
    throw Error(e.toString());
  }
}
