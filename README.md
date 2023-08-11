# Port Service
This is a port Service implementation, really would like to improve the memmory efficiency aspects of it.
I think that I was too ambituous trying to implement most of the requirements. In this matter, first idea
comes to mind is that using more optimized 3d-party libraries to handle json parsing, like `jsoniter` is
going to help a lot. Other thing that we could do is to parse only the keys of the JSON object first,
retrieving a map with keys that we can use as transparent cache, so it will have all values empty at first,
and when a user queries the service, we can lookup the port in the file, store it in the map and then serve
it to the user.

Also added a TO-DO list in this file with things I would add to call this project complete.

## How to run locally
Easiest way to run it locally. We can also run it in a docker container as described below.

```
go run main.go
```

## How to run tests
```
go test ./...
```

## How to build image
```
docker build . -t portservice:0.0.1
```

### Run it with docker (including database)
```
docker-compose up -d
```

### Run it without building an image
```
go run ./...
```

--- 
TODO:
 * Instead of decoding the whole file, traverse it decoding one item at a time
 * Add more tests
 * Add at least one end-to-end test for the whole process, that is setting a file and then testing that rows in the DB are correct
 * Add a benchmark for the loading function
 * Add documentation
 * Add a new process that serves the ports API, with a GET endpoint that can filter port by code or by nearest location 
