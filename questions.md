## Questions
- What's the point of creating another instance with app.mount? 
  - Seems like it keeps the main.go code clean by abstracting the api routes 
- What's the point of utils.go when you can use godotenv? 
  - I'm assuming it's because of deployment, though I'd assume doing loadenv would achieve the same thing
- Do we initialize collections for the new semester or should that be a manual process?
  
## Comments
- I used godotenv to load environments since os.Setenv was giving me issues. 