CREATE TABLE projects (
    uid SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    "short-desc" VARCHAR NOT NULL,
    "long-desc" VARCHAR,
    team TEXT[],  
    link VARCHAR,
    image VARCHAR,
    "tech-stack" TEXT[],
    cohort VARCHAR,
    topic TEXT[] 
);
