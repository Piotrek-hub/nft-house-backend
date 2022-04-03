package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func GetHouses(params ...string) struct {
	Houses []House `json:"houses"`
} {
	content, err := ioutil.ReadFile("db/db.json")
	if err != nil {
		log.Fatal(err)
	}

	payload := struct {
		Houses []House `json:"houses"`
	}{}

	if err := json.Unmarshal(content, &payload); err != nil {
		log.Fatal(err)
	}

	if len(params) > 0 {
		for _, house := range payload.Houses {
			if params[0] == house.Id {
				return struct {
					Houses []House `json:"houses"`
				}{Houses: []House{house}}
			}
		}
	}

	return payload
}
func AddHouse() {}
