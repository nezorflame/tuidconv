# tuidconv
Small utility for the extraction of datetime from Cassandra's [timeuuid](http://docs.datastax.com/en/cql/3.3/cql/cql_reference/uuid_type_r.html) type (which is essentialy a Type 1 UUID).

## Usage
tuidconv **time-uuid** **_location_**

Location is set in [TZ format](https://www.wikiwand.com/en/List_of_tz_database_time_zones) according to the [IANA Time Zone database](https://www.iana.org/time-zones).

If skipped, location is set to Etc/GMT.

## Example
tuidconv a7b6fe0a-2afd-11e7-b9f8-005056864dbf Europe/Moscow
```
time: 2017-04-27 06:57:38.696961 +0300 MSK
```

tuidconv a7b6fe0a-2afd-11e7-b9f8-005056864dbf
```
time: 2017-04-27 03:57:38.696961 +0000 GMT
```
