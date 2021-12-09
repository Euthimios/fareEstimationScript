package parser

import (
	"fmt"
	"strconv"
	"thaBeat/ride/model"
)

func ParseData(input [][]string) ([]model.Ride, error) {

	var locations []model.Signal
	var ride []model.Ride
	var id string

	for row := range input {

		// parse every row in order to get the id and the locations
		currentID, currentLocation, err := parseRow(input[row])

		//check for errors during the parsing
		if err != nil {
			return nil, fmt.Errorf("wrong data sended to parser; error: %v", err)
		}

		// in case the file contains data for different id_rides
		if len(locations) != 0 && id != currentID {
			rides := model.Ride{
				ID:              id,
				LocationSignals: locations,
			}
			// append the data at ride
			ride = append(ride, rides)
			// empty the locations in order to add the
			// new data from the new id_ride
			locations = []model.Signal{}
		}
		id = currentID
		locations = append(locations, *currentLocation)

	}

	rides := model.Ride{
		ID:              id,
		LocationSignals: locations,
	}

	ride = append(ride, rides)

	return ride, nil
}

func parseRow(row []string) (string, *model.Signal, error) {

	id := row[0]
	latitude, errLat := strconv.ParseFloat(row[1], 64)
	longitude, errLon := strconv.ParseFloat(row[2], 64)
	timestamp, errTime := strconv.ParseInt(row[3], 10, 32)

	if errLat != nil || errLon != nil || errTime != nil {
		return "", nil, fmt.Errorf("failed to parse row")
	}

	return id, &model.Signal{
		Latitude:  latitude,
		Longitude: longitude,
		Timestamp: int32(timestamp),
	}, nil
}
