# timeline
Store data points or events in a timeline

There are two ways to store points with **timeline**

 - Send a *start* point followed by *stop* point. The response for the *start* point will include an id that must be provided to the *stop* point. Note that the *stop* point will not complain if you provide an invalid or non-existant start point id.
 - Send a *simple* data point, the difference is this has no followup *stop* point. If desired, you can always add a *duration* field to this.

The only required field for either of these is the *series* field. It will keep all of the points in the same series "grouped" together.

# Endpoints

## /api/point/simple

## /api/point/start

## /api/point/stop

# How to run it

 * install [InfluxDB](http://influxdb.com/ InfluxDB) (osx: `brew update && brew install influxdb`)
 * run influxdb, set up a database and user (See [InfluxDb Getting Started](http://influxdb.com/docs/v0.9/introduction/getting_started.html "InfluxDB Getting Started")).
 * `go get github.com/jeffbmartinez/timeline`
 * `cp config/influxdb.json.template config/influxdb.json`
 * fill out appropriate values in *config/influxdb.json*
 * `go run timeline.go`

# How to use api

## Just use http get requests

Example:

    curl 'example.com/api/point/simple?series=temperatures&location=living-room&value=70'
