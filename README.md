# Timeline
Store measurements and events in a timeline.

Measurements and events belong to a *series* and can have associated with them any number of *tag*s. Tags are indexed and can help with retrieving your data. Measurements have a name and a value associated with them (e.g. "temparature, value=10.0"). Events have a name, but no associated value (e.g., "restart").

## How to store a data point

Measurements and events must be associated with a *series*. This is a sort of grouping which the data belongs to. For example, "living-room-temperature" could be a series, which contains a number of measurements. To store something, send an http POST request with appropriate parameters using any http library to the `/api/measurement` or `/api/event` endpoints.

### Examples using curl

    curl -X POST 'timeline-service.com/api/measurement?name=living-room-temperature&value=70'
    
    OR
    
    curl -d 'name=living-room-temperature&value=70' 'timeline-service.com/api/measurement'


### API

See the [api documentation](docs/api.md) for more details

## How to run it

 * Install [InfluxDB](http://influxdb.com/ InfluxDB) (osx: `brew update && brew install influxdb`)
 * Run `influxd run -config config/influxdb/development.toml`, set up a database and user (See [InfluxDb Getting Started](http://influxdb.com/docs/v0.9/introduction/getting_started.html "InfluxDB Getting Started"))
 * `go get github.com/jeffbmartinez/timeline`
 * `cp config/timeline/influxdb.json.template config/timeline/influxdb.json`
 * Fill out appropriate values in *config/timeline/influxdb.json*
 * `go run timeline.go`
