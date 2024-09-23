# Usage
You can parse values to overwrite the code defaults using flags:

`go run cmd/loadgen/main.go -url http://google.com -c 5 -r 50`

this for example will send 100 total requests (-r 100) to http://google.com (-url) with a concurrency level of 10 (-c 10).

# Current implementation

Sends HTTP requests concurrently
Collects and print response times
Customizable with flags