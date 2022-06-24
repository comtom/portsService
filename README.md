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

## How to deploy
```
docker build . -t portservice:v0.1
```

### Run it with docker
```
docker run portservice:v0.1
```

--- 
TODO:
 * Improve memory efficiency of loading function
 * Save data to a database, chose postgres and set it up in a docker compose
 * Complete unit tests for main package
 * Add at least one end-to-end test for the whole process, that is setting a file and then testing that rows in the DB are correct
 * Add a benchmark for the loading function
 * Add a GET endpoint to serve ports with pagination and the ability to retrieve any port by its "unlocs" (I believe it corresponds to UN/LOCODE Code)
 * Add a flag/ environment variable to be able to change the file-name
 * Add a volume in the dockerfile, to be able to change the input file easily  
 * Add more documentation
 * Add kubernetes yaml to be able to deploy to a real production environment
