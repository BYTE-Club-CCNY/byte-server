- [x] Go mongodb initialize and setup
- [x] Add collections to MongoDB
  - [x] Endpoint to create collection
  - [x] Endpoint to see data within a collection
- [x] Draft Endpoints
  - [x] Create draft endpoint
  - [x] Edit draft endpoint
  - [x] View draft endpoint
  - [x] Publish draft endpoint
- [x] Cohort Applications 
  - [x] Submit application
  - [x] Save user's application as draft before submitting
  - [x] Retrieve all user applications that are completed
    - [x] Implement Pagination
- [x] Deploy MongoDB 

## Next Steps:
### Endpoints
- /auth/register (postgres)
  - table columns
    - email
    - password
      - need to hash
    - user_id 
- /auth/login (postgres)
  - enter email + password
    - check if email matches/exists, then check if hashed password matches 
- /auth/forgot-password/ (postgres)
  - need email or user_id
- /voting/give-vote (mongodb)
  - need following info:
    - cohort_id
    - user_id of application submitter
    - user_id of voter
    - vote and comment
  - will be attached to a user's application
  - {
      "votes" : [
        {
          "cabinet" : "user_id of cabinet member",
          "vote" : "accept",
          "comment" : "looks passionate"
        }
      ]
    }
- /voting/update-vote (mongodb)
  - need following info:
    - cohort_id
    - user_id of application submitter
    - user_id of voter
    - vote and comment
  - check if vote exists, if it doesn't return error 
  - if vote exists, replace the json information with vote and comment 
- /voting/final-vote (postgres)
  - need following info:
    - user_id of application submitter
    - vote (accept/deny)
    - cohort_id
  - Submits user_id and vote
- /apps/get-specific-app (mongodb)
  - need following info: 
    - user_id of user 
    - cohort_id
  - return application information, without voting details
### Other Technical Stuff
- deployed resume link
  - Potential Ideas:
    - google drive
    - s3 bucket
    - mongodb feature 