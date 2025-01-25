## Questions
- What's the point of creating another instance with app.mount? 
  - Seems like it keeps the main.go code clean by abstracting the api routes 
- What's the point of utils.go when you can use godotenv? 
  - I'm assuming it's because of deployment, though I'd assume doing loadenv would achieve the same thing
  - (fahad) its cause i wanted to flex and make my own dotenv loader - cause its relatively easy to make
      it works just as long as the .env file is in the root of the application directory
- Do we initialize collections for the new semester or should that be a manual process?
  - Might not be worth the hassle / doesn't sound intuitive
  
## Comments
- I used godotenv to load environments since os.Setenv was giving me issues. 