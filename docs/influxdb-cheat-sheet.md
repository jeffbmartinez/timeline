# InfluxDB Cheat Sheet

Cheatsheet for common commands using `influx` command line client.

By default, `influx` will attempt to connect to `localhost:8086`.

## Simple commands

`create database MyDatabaseName`

`show databases`

`use database MyDatabaseName`

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

