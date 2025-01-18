package apps

import (
	"github.com/gofiber/fiber/v2"
)

func ApplicationRouting() *fiber.App {
	app := fiber.New()

	app.Get("/collection-data", GetCollectionData) // returns all documents from a collection
	app.Post("/new-collection", NewCohort) // creates a new collection
	app.Post("/edit-cohort", EditCohortDraft) // edit a cohort draft

	/*
	app.Get("/apps/view-draft") // view draft
	app.Post("/apps/create-draft") // create a new draft
	*/

	/*
	create cohort, cohort id, application layout, set deadline
	app layout will be stored as a template in mongo
	deadline, cohort_id stored in postgres
	*/

	/*
	app.Post("publish-draft")

	app.Get("/apps/get-template") // question layout per cohort (mongo)*/

	/* saves application w/ userid (uuid same as the one in postgres) 
	to mongo application collection */
	/*
	app.Post("/apps/submit-app") 

	app.Post("/apps/save-app-draft") // save draft only if app is not submitted

	app.Get("/apps/view") // get existing application(s) query by id and page
	*/
	
	return app
}