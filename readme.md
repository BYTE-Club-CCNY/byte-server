# API Documentation

## Base URL
`https://test.byteccny/`
*localhost runs on port 3000*

---

## Endpoints

### General
- **`GET /`**  
  - **Description:** Test endpoint to verify the server is running.  
  - **Response:** `"BYTE Server is running!"`

---

### Projects

**Base Path:** `/projects`

#### `GET /projects/`
- **Description:** Test endpoint for the projects module.  
- **Response:** Status 200 (OK).

#### `GET /projects/get`
- **Query Parameters:**
  - `cohort` (optional): Cohort ID (integer). Defaults to `-1`.
  - `name` (optional): Project name (string).
  - `page` (optional): Page number (integer). Defaults to `1`.
- **Response:** JSON array of projects
```json
{
    "ID": "string", // UID of project AND team
    "Member1": "string",
    "Member2": "string",
    "Member3": "string",
    "Member4": "string",
    "ProjectName": "string",
    "ShortDesc": "string",
    "LongDesc": "string",
    "Link": "string",
    "Image": "string",
    "TechStack": "string array",
    "Topic": "string array",
    "Cohort": "string",
}
```

#### `POST /projects/add`
- **Request Body (JSON or Form-Data):**
  ```json
  {
    "name": "string",
    "short_desc": "string",
    "long_desc": "string",
    "member1": "UUID",
    "member2": "UUID (optional)",
    "member3": "UUID (optional)",
    "member4": "UUID (optional)",
    "link": "string",
    "image": "string",
    "tech_stack": "string (JSON array format)",
    "topic": "string (JSON array format)",
    "cohort": "integer"
  }
- **Headers**: `application/json` or `multipart/form-data` (both are accepted)
- **Validation**: All fields are required unless marked optional.
- **Response**: Status 200 on success.

### Users
**Base Path:** `/users`

---

### `GET /users/`
- **Description:** Test endpoint for the users module.
- **Response:** Status 200 (OK).

### `POST /users/add`
- **Description:** adds a user to the database
- **Request Body (JSON):**
```json
    {
    "name": "string",
    "cuny_email": "string",
    "emplid": "string",
    "personal_email": "string (optional)",
    "discord": "string (optional)"
    }
```
- **Validation:** All fields are required unless marked optional.
- **Response:** Status 200 on success.