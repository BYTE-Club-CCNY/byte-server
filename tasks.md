- [x] Go mongodb initialize and setup
- [x] Add collections to MongoDB
  - [x] Endpoint to create collection
  - [x] Endpoint to see data within a collection
- [x] Draft Endpoints
  - [x] Create draft endpoint
  - [x] Edit draft endpoint
  - [x] View draft endpoint
  - [x] Publish draft endpoint
- [ ] Cohort Applications 
  - [ ] Submit application
  - [ ] Save user's application as draft before submitting
  - [ ] Retrieve all user applications that are completed
    - [ ] Implement Pagination

## Brainstorm
> I think we should just have one collection for every semester, and then every document would have a "documentType" which would let us separate applications, draft, and voting.
- What does a draft need?
  - you can have arbritrary amount of questions
  - documentType 
  - deadline