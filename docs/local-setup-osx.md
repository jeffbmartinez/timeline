# How to run timeline locally on osx

## Prerequisites

- [homebrew](http://brew.sh)
- [golang](http://golang.org)
- `brew install git`
- `brew install influxdb`

## Get timeline code

`git clone git@github.com:jeffbmartinez/timeline.git`

## Run influxdb

`influxd -config timeline/config/influxdb/development.toml`

    To have launchd start influxdb at login:
        ln -sfv /usr/local/opt influxdb/*.plist ~/Library/LaunchAgents
    Then to load influxdb now:
        launchctl load ~/Library/LaunchAgents/homebrew.mxcl.influxdb.plist

The default config file is at:

    /usr/local/etc/influxdb.conf.default

## Run timeline

Influxdb must be running already, then:

    go run timeline.go

Or if you want to build it and run the executable:

    go build && ./timeline

## Troubleshooting

If you get something like this:

    [ERROR] Could not ping the influxDB server (Get http://localhost:8086/ping: dial tcp 127.0.0.1:8086: connection refused)
    Could not connect to influxDB
    [FATAL] InfluxDB connection test failed: Could not ping the influxDB server (Get http://localhost:8086/ping: dial tcp 127.0.0.1:8086: connection refused)
    exit status 1

You just forgot to run influxdb, it is not running on that port, or perhaps you changed the influxdb configuration so as to not allow connections from wherever you're connecting to it.
