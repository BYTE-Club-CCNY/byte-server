# Byte-server
Server for byte

## Description
Server is built with Node.js and Express. Provides endpoints for managing a collection of projects, including creating, reading and updating. Exists both database and a local JSON file backup for storing and quering projects.

## How to use
Creating, Reading and Updating can be done from the database (default approach) endpoints. Only reading can be done for the local JSON file.

## AUTH
In order to `POST` or `PUT` data to the database, you must provide **Basic Auth**. *Only system admins and active developers have access to this.* 
- Username is your name (case doesn't matter)
- Password is your API Key provided

## Endpoints for DB
### Get projects
`/projects/get`

Retrieves all projects in the database and supports additional query params for filtering.

Filters include:
- **team**: TYPE string(s)
- **cohort**: TYPE string(s)
- **name**: TYPE string(s)

Filters are stackable, meaning you can do something similar to:  
`/projects/get?team=John&cohort=1&team=Jane`  
This will return all projects that have John **AND** Jane in their team and are in cohort 1. You can stack as many filters as you want but be aware that the more filters you add, the more specific the query will be, meaning it might return no results.

*Tip:* To get all projects leave `get` as blank


### Post new projects
`/projects/add`

**Please Fill Out All Fields!!!**  
These fields must be provided as JSON in the body of the request.

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
`/projects/update`

You **MUST** provide the project name you want to update as a query parameter like so:  
`/projects/update?name={project_name}`

You can provide any key value pair you want to update in the body of the request, but it **HAS** to be JSON. This doesn't have to be all the fields, only the ones you want to update.

Reminder that the schema for projects is defined as follows:
- **name**: TYPE string(s)
- **"short-desc"**: TYPE string(s)
- **"long-desc"**: TYPE string(s)
- **team**: TYPE array of strings
- **link**: TYPE string(s)
- **Image**: TYPE string(s)
- **"tech-stack"**: TYPE array of strings
- **cohort**: TYPE string(s)
- **topic**: TYPE array of strings

Any mismatch of the schema types (i.e. providing a string when we expect an array) will return an error.

## Endpoints for Local File

### Get all projects
`/projects`

Retrieves all projects from the local JSON file and also supports additional query params for filtering. 

### Get projects based on team members
`/projects?team={member_name}`

Replace `{member_name}` with a member that is in the project you would like to query. This filter is stackable, meaning you can do something similar to:  
`/projects?team=John&team=Jane`

### Get projects based on cohort
`/projects?cohort={cohort}`

Replace `{cohort}` with the desired cohort to filter projects. This filter is stackable, meaning you can do something similar to:  
`/projects?cohort=1&cohort=2`

### Get projects based on name
`/projects?name={project_name}`

Replace `{project_name}` with the desired project name to filter projects.
This filter is stackable, meaning you can do something similar to:  
`/projects?name=Website&name=Server`