package parser

import (
	"fmt"
	"strconv"
	"thaBeat/internal/app/thaBeat/ride"
)

// ParseData gets as parameter string arrays and for each of them a Ride struct is generated
func ParseData(input [][]string) <-chan ride.Ride {
	dataChan := make(chan ride.Ride)

	go func() {
		var locations []ride.Point
		var id string

		for row := range input {
			// parse every row in order to get the id and the locations
			currentID, currentLocation, err := parseRow(input[row])
			//check for errors during the parsing
			if err != nil {
				fmt.Printf("wrong data sended to parser; error: %v", err)
				continue
			}
			// in case the file contains data for different id_rides
			if len(locations) != 0 && id != currentID {
				r := ride.Ride{
					ID:     id,
					Points: locations,
				}
				// append the data at ride
				dataChan <- r
				// empty the locations in order to add the new data from the new id_ride
				locations = []ride.Point{}
			}
			id = currentID
			locations = append(locations, *currentLocation)
		}
		//for the last ride_id ,or in case thee file contains points for only one ride
		r := ride.Ride{
			ID:     id,
			Points: locations,
		}
		dataChan <- r
		close(dataChan)
	}()
	return dataChan
}

// parseRow parse each row into a ride.Point struct
func parseRow(row []string) (string, *ride.Point, error) {

	if len(row) != 4 {
		return "", nil, fmt.Errorf("expectd 4 elements but row hasn't: %v", row)
	}

	id := row[0]
	latitude, errLat := strconv.ParseFloat(row[1], 64)
	longitude, errLon := strconv.ParseFloat(row[2], 64)
	timestamp, errTime := strconv.ParseInt(row[3], 10, 32)

	if errLat != nil || errLon != nil || errTime != nil {
		return "", nil, fmt.Errorf("failed to parse row")
	}

	return id, &ride.Point{
		Latitude:  latitude,
		Longitude: longitude,
		Timestamp: int32(timestamp),
	}, nil
}
