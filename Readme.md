# Ride Fare Estimation Script

Ride Fare Estimation Script implements an automated way of estimating the fare for each ride so that can flag
suspicious rides for review by the support teams.The fare estimation algorithm is capable of detecting erroneous
coordinates and removing them before attempting to evaluate the ride fare.

Accepts as arguments two files of type CSV. An input file with records in format of type
(id_ride, lat, lng, timestamp), creates the calculation for each id_ride and writes down in the second argument file the
information in format (id_ride, fare_estimate).

## Code Style

It follows the basic layout for Go application projects ,
as it's described [here](https://github.com/golang-standards/project-layout)
It's not an official standard defined by the core GO Dev Team.

## Tech

GO version go1.17 used for the implementation of this project.
You can lear more about the GO Programming language here :

 * The official tutorial [here](https://tour.golang.org/basics/1)

 * Download and install Go from [here](https://golang.org/doc/install)

Also the algorithm  makes use of the Haversine distance formula .
The Haversine Formula determines the great-circle distance between two points on a sphere given their
longitudes and latitudes.You can learn more about Haversine [here](https://en.wikipedia.org/wiki/Haversine_formula)

### Layout

```tree
├── assets
├── cmd
│   └── thabeat
│       ├── main.go
│       └── main_test.go
├── internal
│   └── app
│      └── thabeat
│          │
│          └── ride
│                ├── farecalculation
│                │   ├── farecalculation.go
│                │   └── farecalculation_test.go
│                ├── parser
│                │   ├── parser.go
│                │   └── parser_test.go
│                └── ride.go
├── pkg
│   ├── csv
│   │    ├── csv.qo
│   │    └── csv_test.go
│   └── haversine
│       ├── haversine.go
│       └── haversine_test.go
├── test
│    └── testdata
└── README.md

```
A brief description of the layout:

* `assets` contains the default input / output files.
* `cmd` contains main packages with it's unit test file.
* `internal` application and library code with it's unit test files.
* `pkg` contains code that's ok to use by external application.
* `test` holds all tests data.
* `README.md` is a detailed description of the project.

## Tests

Into the /internal folder is located the application code. For each .go file you can find the _test.go file
which has the unit tests .They test the  components by all possible means and compares the result to the
expected output. The tests written with out assertion , in order  to write the logic to test the results of the unit components.

Additional test data can be found at folder `/test/testdata`. Also at the same folder is located the `coverage report`

To run all tests:

```
go test ./...
```

## How to Use

```bash

Usage:
  `go run .cmd/thabeat/main.go  -in="inpath" -out="outpath"`


Flags:
      in   the path for the file that has the Ride data ,ex : assets/input.csv
      out  the path for the file that will have the calculated data for each Ride , ex :assets/output.csv

      to print a usage message with all defined command-line flags execute :
      `go run .cmd/thabeat/main.go  -h`
```
