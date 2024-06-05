# Byte-server
Server for byte

# Description
Server is built with Node.js and Express. Provides endpoints for managing a collection of projects, including creating, reading and updating. Exists both database and a local JSON file backup for storing and quering projects

# How to use
Creating, Reading and Updating can be done from the database (default approach) endpoints. Only reading can be done for the local JSON file

## Endpoints for DB

### Get projects
`/projects/get`

Retrieves all projects in the database and supports additional query params for filtering

### Post new projects
`/projects/add`

**Please Fill Out All Fields!!!**

The schema for projects is defined as follows:
- **name**: TYPE string(s)
- **"short-desc"**: TYPE string(s)
- **"long-desc"**: TYPE string(s)
- **team**: TYPE array of strings
- **link**: TYPE string(s)
- **Image**: TYPE string(s)
- **"tech-stack"**: TYPE array of strings
- **cohort**: TYPE string(s)
- **topic**: TYPE array of strings

### UPDATE projects
`/projects/update?name={project_name}`

Replace `{project_name}` with the desired project name you want to update

## Endpoints for Local File

### Get all projects
`/projects`

Retrieves all projects from the local JSON file and also supports additional query params for filtering

### Get projects based on team members
`/projects?team={member_name}`

Replace `{member_name}` with a member that is in the project you would like to query

### Get projects based on cohort
`/projects?cohort={cohort}`

Replace `{cohort}` with the desired cohort to filter projects

### Get projects based on name
`/projects?name={project_name}`

Replace `{project_name}` with the desired project name to filter projects