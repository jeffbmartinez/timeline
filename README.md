# timeline
Store data points or events in a timeline.

Data points belong to a *time series* and can have associated with them any number of *tag*s. Tags are indexed and can help with retrieving your data later on.

## How to store a data point

Data points must be associated with a *time series*. This is a sort of grouping which the data point belongs to. For example, "living-room-temperature" could be a time series. To store a point, send an http GET request with appropriate parameters using any http library to the `/api/point/simple` endpoint.

### Example using curl

    curl 'timeline-service.com/api/point/simple?series=living-room-temperature&value=70'


The only required field for either of these is the *series* field, without it the request will fail. You can add any number of arbitrary fields in the request.

## How to run it

 * Install [InfluxDB](http://influxdb.com/ InfluxDB) (osx: `brew update && brew install influxdb`)
 * Run `influxd run -config config/influxdb/development.toml`, set up a database and user (See [InfluxDb Getting Started](http://influxdb.com/docs/v0.9/introduction/getting_started.html "InfluxDB Getting Started"))
 * `go get github.com/jeffbmartinez/timeline`
 * `cp config/timeline/influxdb.json.template config/timeline/influxdb.json`
 * Fill out appropriate values in *config/timeline/influxdb.json*
 * `go run timeline.go`
