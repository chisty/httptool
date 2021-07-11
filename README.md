# httptool

A simple http tool which makes http requests and prints the address of the request along with the MD5 hash of the response.

## Features

 - The tool takes command line arguments as input request.
 - By default, this tool will run **10** concurrent goroutines to execute `HTTP GET` request. 
 - To prevent local resource exhausting, user can update the default concurrent worker value. In that case, a flag **-parallel** along with a **numeric** number (the number of concurrent request to execute) needs to be provided.
 - If number of worker is **more** than number of requests to process, then tool will create **minimum** number of worker. (ex: if **5** *request argument* provided, but *-parallel flag* is **10** then only **5** worker will be created).
 - This tool will add `http` schema by default if request `URL Schema` is empty. So, google.com will change to http://google.com 

## How to use
Build commands

    go build						

Execute with default settings. (1 worker will execute)

    ./httptool	http://google.com

Execute with custom settings.  (2 worker will execute)

    ./httptool -parallel 2 http://google.com facebook.com twitter.com http://yahoo.com

Execute with custom settings.  (4 worker will execute)

    ./httptool -parallel 10 http://google.com facebook.com twitter.com http://yahoo.com

## Misc

 - A Makefile is added in this repository to test this tool. User can run `make run` or `make runAll` command to build, run unit tests, and execute the tool with dummy data.
 - If the `request url` is not **valid**, the tool will print request along with the `error message` instead of MD5 hash.