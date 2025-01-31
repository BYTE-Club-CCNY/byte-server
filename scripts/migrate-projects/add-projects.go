package main

import (
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
	schema "byteserver/pkg/schemas"
	"github.com/lib/pq"
	"byteserver/pkg/database"
	"github.com/google/uuid"
	"byteserver/pkg/utils"
	"log"
)

type JsonStruc struct {
	Name string `json:"name" form:"name" validate:"required"`
	ShortDesc  string `json:"short-desc" form:"short-desc" validate:"required"`
	LongDesc string `json:"long-desc" form:"long-desc" validate:"required"`
	Team []string `json:"team" form:"team" validate:"required"` 
	Link string `json:"link" form:"link" validate:"required"` 
	Image string `json:"image" form:"image" validate:"required"`
	TechStack []string `json:"tech-stack" form:"tech-stack" validate:"required"` 
	Cohort int `json:"cohort" form:"cohort" validate:"required"` 
	Topic []string `json:"topic" form:"topic" validate:"required"` 
}

type JsonData struct {
	JsonData []JsonStruc `json:"projects"`
}

func main () {
	utils.InitEnv()
	database.InitDB()
	jsonFile, err := os.Open("./scripts/migrate-projects/data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var jsondata JsonData
	json.Unmarshal(byteValue, &jsondata)
	for i := 0; i < len(jsondata.JsonData); i++ {
		var project schema.Project
		project.Name = jsondata.JsonData[i].Name
		project.CohortID = jsondata.JsonData[i].Cohort
		project.ShortDesc = jsondata.JsonData[i].ShortDesc
		project.LongDesc = jsondata.JsonData[i].LongDesc
		project.Image = jsondata.JsonData[i].Image
		project.Link = jsondata.JsonData[i].Link
		project.Topic = pq.StringArray(jsondata.JsonData[i].Topic)
		project.TechStack = pq.StringArray(jsondata.JsonData[i].TechStack)

		
		members := jsondata.JsonData[i].Team	
		team := new(schema.Team)
		for i, uuid_string := range members {
			member_uuid := uuid.MustParse(uuid_string)
			if i == 0 && member_uuid != uuid.Nil {
				team.Member1 = member_uuid
			} else if i == 1  && member_uuid != uuid.Nil{
				team.Member2 = member_uuid
			} else if i == 2 && member_uuid != uuid.Nil {
				team.Member3 = member_uuid
			} else if i == 3 && member_uuid != uuid.Nil {
				team.Member4 = member_uuid
			}
		}
		fmt.Println("In process of migrating: ", project)
		res := database.DB.Create(&team)
		if res.Error != nil {
			log.Println("error creating team")
			log.Println("error on project: ", project)
			panic(res.Error)
		}
		project.ID = team.ID
		res = database.DB.Create(&project)
		if res.Error != nil {
			log.Println("error creating project")
			log.Println("error on project: ", project)
			panic(res.Error)
		}
		fmt.Println("Successfully Migrated: ", project.Name)
	}
	
}