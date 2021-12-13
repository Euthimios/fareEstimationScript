# Ride Fare Estimation Script

Ride Fare Estimation Script implements an automated way of estimating the fare for each ride so that can flag
suspicious rides for review by the support teams.The fare estimation algorithm is capable of detecting erroneous
coordinates and removing them before attempting to evaluate the ride fare.

Ride Fare Estimation Script accepts as arguments two files of type CSV. An input file with records in format of type
(id_ride, lat, lng, timestamp), creates the calculation for each id_ride and writes down in the second argument file the
information in format (id_ride, fare_estimate).

## Code Style

It follows the basic layout for Go application projects ,
as it's described [here](https://github.com/golang-standards/project-layout)
It's not an official standard defined by thee core GO Dev Team.

## Tech

GO version go1.17 used for the implementation of this project.
You can lear more about the GO Programming language here :

 * The official tutorial [here](https://tour.golang.org/basics/1)

 * Download and install Go from [here](https://golang.org/doc/install)

Also the algorithm  makes use of the Haversine distance formula .
The haversine formula determines the great-circle distance between two points on a sphere given their
longitudes and latitudes.You can learn more about Haversine [here](https://en.wikipedia.org/wiki/Haversine_formula)
## Tests

Into the /internal folder is located the application code. For each .go file you can find the _test.go file
which has the test cases . It tests the unit component by all possible means and compares the result to the
expected output. The tests written with out assertion , in order  to write the logic to test the results of the unit components.

Additional test data can be found at folder /test/testdata. Also at the same folder is located the coverage report

To run all tests:

```
go test ./...
```

## How to Use

```bash

Usage:
  go run cmd/thabeat/main.go  -in="inpath" -out="outpath"


Flags:
      in   the path for the file that has the Ride data ,ex : resources/input.csv
      out  the path for the file that will have the calculated data for each Ride , ex :resources/output.csv
```