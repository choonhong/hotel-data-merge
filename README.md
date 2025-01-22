# hotel-data-merge

This project is a simple example of how to merge data from different sources. The data is related to hotels provided by different sources. The goal is to merge the data into a single data source.

It also serves endpoints to query the hotels data by hotel ids and destination.

## Data schema

The data schema is defined in the `ent/schema/hotel.go` file. The data is provided in JSON format.

This project uses the `ent` framework to define the data schema and generate the data model. It also provide code generation for the data model and migration.

After changing the schema, you need to run the following command to generate the data model and migration.

```
$ make ent
```

## Run on docker

```
$ make docker
```

## Call the endpoints

1. Get all hotels
```
$ curl "http://localhost:80/v1/hotels"
```

2. Get hotels by hotel ids
```
$ curl "http://localhost:80/v1/hotels?ids=f8c9&ids=iJhz"
```

3. Get hotels by destination
```
$ curl "http://localhost:80/v1/hotels?destination=1122"
```
