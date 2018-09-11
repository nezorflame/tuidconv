# tuidconv [![Build Status](https://travis-ci.org/nezorflame/tuidconv.svg?branch=master)](https://travis-ci.org/nezorflame/tuidconv) [![Go Report Card](https://goreportcard.com/badge/github.com/nezorflame/tuidconv)](https://goreportcard.com/report/github.com/nezorflame/tuidconv)

Small utility for the extraction of datetime from UUID v1 and v2.

Supports Cassandra's [timeuuid](http://docs.datastax.com/en/cql/3.3/cql/cql_reference/uuid_type_r.html) type (which is essentialy a `Type 1 UUID`).

[GolangCI status report](https://golangci.com/r/github.com/nezorflame/tuidconv)

## Requirements

Go version 1.11+ (to support [modules](https://github.com/golang/go/wiki/Modules)).

## Usage

tuidconv **time-uuid** **_location_**

Location is set in [TZ format](https://www.wikiwand.com/en/List_of_tz_database_time_zones) according to the [IANA Time Zone database](https://www.iana.org/time-zones).

If skipped, location is set to `Etc/GMT`.

## Installation

```bash
go get -u github.com/nezorflame/tuidconv
# To use with Go modules, do the next steps as well:
cd $GOPATH/src/github.com/nezorflame/tuidconv
GO111MODULE=on go install
```

## Example

```bash
$ tuidconv a7b6fe0a-2afd-11e7-b9f8-005056864dbf Europe/Moscow

time: 2017-04-27 06:57:38.696961 +0300 MSK
unix: 1493265458
```

```bash
$ tuidconv 7adb9273-88b0-11e7-80a2-005056864dbf

time: 2017-08-24 09:42:01.4921331 +0000 GMT
unix: 1503567721
```
