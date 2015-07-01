# InfluxDB Cheat Sheet

Cheatsheet for common commands using `influx` command line client.

By default, `influx` will attempt to connect to `localhost:8086`.

## Simple commands

`create database MyDatabaseName`

`show databases`

`use database MyDatabaseName`

`insert cpu,host=serverA,region=us_west value=0.64`: A point with the measurement name of **cpu** and tag **host** has now been written to the database, with the measured value of **0.64**.

`exit`

## How data in influxdb is structured

- **Databases** are organized by **Time Series**
- **Time Series** each contain **points**
- **Points** have:
  - a *timestamp*
  - a *measurement* (e.g. "cpu_load")
  - one or more *field*s
  	- e.g. "value=0.64" or "15min=0.78"
  	- not indexed
  - zero or more *tag*s
  	- e.g. "host=server01", "region=EMEA", or "datacenter=Frankfurt"
  	- indexed (so they are queryable)

### Line storage Protocol

    <measurement>[,<tag-key>=<tag-value>...] <field-key>=<field-value>[,<field2-key>=<field2-value>...] [unix-nano-timestamp]

Examples:

    cpu,host=serverA,region=us_west value=0.64 payment,device=mobile,product=Notepad,method=credit billed=33,licenses=3 1434067467100293230

    stock,symbol=AAPL bid=127.46,ask=127.48  temperature,machine=unit42,type=assembly external=25,internal=37 1434067467000000000
