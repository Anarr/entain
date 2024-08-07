# entain

Entain is a microservice responsible for manage user balance across the app


### Local dev environment
If you wish to start the API server locally try these steps:
- run `make compose-up`

If you want to run tests
- run `make test`

Api Doc available on:
- http://localhost:8888/swagger/index.html

** if you want to monitor cancelled request just run the app and
look logs in cli. the cancellation occurs per minute. You can easily change
the interval from `internal/config/config.json` `task_interval` key.
````
{"time":"2023-05-01T09:30:39.062245425Z","level":"INFO","prefix":"-","file":"manager.go","line":"39","message":"take latest 0 requests for cancellation"}
````
