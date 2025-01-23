# hotel-data-merge

This project is a simple example of how to merge data from different sources. The data is related to hotels provided by different sources. The goal is to merge the data into a single data source.

It also serves endpoints to query the hotels data by hotel ids and destination.

## Data schema

The data schema is defined in the `ent/schema/hotel.go` file. The data is provided in JSON format.

This project uses the [entgo.io](https://entgo.io/) framework to define the data schema and generate the data model. 
It also provide code generation for CRUD API and migration code.

After makeing change in the schema, run the following command to generate the data model and migration.

```
$ make ent
```

Generated code will be placed in the `ent` directory.

## API definition

The API provided by this project is defined in the `docs/swagger.yml` file.

After making changes in the API, run the following command to generate the API handler and types code.

```
$ make api
```

Generated code will be placed in the `restapi` directory.

## Mock interface

The mock interfaces is generated using [mockery](https://github.com/vektra/mockery) tool.

After making changes in the interface, run the following command to generate the mock interfaces.

```
$ mockery
```

## Run on docker

```
$ make docker
```

## Make requests

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

## Run tests

To run tests on local
```
$ make test
```

To run tests on docker
```
$ docker compose up
```
