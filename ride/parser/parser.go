package parser

import (
	"fmt"
	"strconv"
	"thaBeat/ride/model"
)

func ParseData(input [][]string) (<-chan model.Ride, <-chan error) {

	ridePositions := make(chan model.Ride)
	channelError := make(chan error)

	go func() {
		var locations []model.Signal
		var rides model.Ride
		var id string

		for row := range input {
			// parse every row in order to get the id and the locations
			currentID, currentLocation, err := parsePosition(input[row])
			//check for errors during the parsing
			if err != nil {
				channelError <- fmt.Errorf("wrong data sended to parser; error: %v", err)
			}
			// in case the file contains data for different id_rides
			if len(locations) != 0 && id != currentID {
				rides := model.Ride{
					ID:              id,
					LocationSignals: locations,
				}
				// append the data at ride
				ridePositions <- rides
				// empty the locations in order to add the
				// new data from the new id_ride
				locations = []model.Signal{}
			}
			id = currentID
			locations = append(locations, *currentLocation)
		}
		//for the last ride_id ,or in case thee file contains positions for only one ride
		rides = model.Ride{
			ID:              id,
			LocationSignals: locations,
		}
		ridePositions <- rides
	}()

	return ridePositions, channelError
}

//parsePosition gets as argument a slice of strings .
//transform the strings to Signal Struct
func parsePosition(row []string) (string, *model.Signal, error) {

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
