package main

import (
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
	schema "byteserver/pkg/schemas"
	"github.com/lib/pq"
	"strings"
	"byteserver/pkg/database"
	"github.com/google/uuid"
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
	jsonFile, err := os.Open("./scripts/migrate-data/data.json")
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

		var peopleNames []string
		members := jsondata.JsonData[i].Team
		for i := range members {
			if members[i] != "" {
				peopleNames = append(peopleNames, strings.TrimSpace(members[i]))
			}
		}

		var peopleIDs []uuid.UUID;
		query := database.DB.Select("uid").Model(&schema.User{})
		err := query.Where("first_name IN (?)", peopleNames).Scan(&peopleIDs)
		
		if err.Error != nil {
			panic(err)
		}
	
		if len(peopleIDs) != len(peopleNames) {
			panic("Some or all of these people do not exist")
		}
	
		// create team & return team ID
		team := new(schema.Team)
		if len(peopleIDs) > 0 {
			team.Member1 = peopleIDs[0]
		}
		if len(peopleIDs) > 1 {
			team.Member2 = peopleIDs[1]
		}
		if len(peopleIDs) > 2 {
			team.Member3 = peopleIDs[2]
		} 
		if len(peopleIDs) > 3 {
			team.Member3 = peopleIDs[3]
		}
	
		res := database.DB.Create(&team)
		if res.Error != nil {
			panic(res.Error)
		}
		project.ID = team.ID
		database.DB.Create(project)
	}
	
}