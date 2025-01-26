package apps

import (
	"github.com/gofiber/fiber/v2"
)

func App() *fiber.App {
	app := fiber.New()

	app.Get("/collection-data", GetCollectionData) // returns all documents from a collection
	app.Post("/new-cohort", NewCohort) // creates a new collection
	app.Post("/edit-draft", EditCohortDraft) // edit a cohort draft

	app.Get("/view-draft", ViewDraft)
	
	app.Post("/create-draft", CreateDraft) // create a new draft

	/*
	create cohort, cohort id, application layout, set deadline
	app layout will be stored as a template in mongo
	deadline, cohort_id stored in postgres
	*/

	app.Post("publish-draft",PublishDraft)

	app.Get("/get-template", ViewTemplate) // question layout per cohort (mongo)

	/* saves application w/ userid (uuid same as the one in postgres) 
	to mongo application collection */

	app.Post("/submit-app", SubmitApp) 

	//app.Post("/save-app-draft") // save draft only if app is not submitted

	//app.Get("/view") // get existing application(s) query by id and page

	return app
}