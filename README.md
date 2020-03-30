# CallsWebService
Simple service to manage call records and extract statistics written in GO and using a Postgres Database.

## Requirements
1. Have **docker**, **docker compose**, **go(local only)**  installed
- [Install Docker](https://docs.docker.com/install/)
- [Install Docker Compose](https://docs.docker.com/compose/install/)
- [Install Go](https://golang.org/doc/install)
2. Have docker running.
3. Have port 8989 (and 5432 if running locally) on your system available

## Build
```bash
cd build
#Local
./build-local.sh
#Docker
./build.sh
``` 

## Run
```bash
#Local
./run-local.sh
#Docker
./run.sh
``` 

## API
Details information about the API can be viewed in the **swagger.yaml** file and a postman Collection is also available in **CallWebService.postman_collection.json**  

## Configuration

**Server**
Example config file can be found in **conf/server.json**

| Name          | Description           | Example  |
| ------------- |:---------------------:| --------:|
| database.host | Postgres database hostname of ip address         | 127.0.0.1    |
| database.port | Postgres database port              |   5432    |
| database.user | Postgres database username          |    bob |
| database.password | Postgres database password          |    admin123 |
| server.port | Port used by server API       |    8080 |
| server.phone_number_regex | Regex to validate phone numbers      |    ^[0-9]+$ |
| server.call_cost.inbound_price_1 | Price for the 1st tier of inbound calls (as integer with 4 cent decimal places) |0|
|server.call_cost.inbound_price_2 | Price for the 2nd tier of inbound calls (as integer with 4 cent decimal places) |0| 
| server.call_cost.inbound_price_threshold | Minute Threshold for price tier of inbound calls |0 |
|server.call_cost.outbound_price_1 | Price for the 1st tier of inbound calls (as integer with 4 cent decimal places) |5|
|server.call_cost.outbound_price_2 | Price for the 2nd tier of inbound calls (as integer with 4 cent decimal places) |10| 
| server.call_cost.outbound_price_threshold | Minute Threshold for price tier of inbound calls | 5 | 

**Client**
Example config file can be found in **conf/client.json**

| Name          | Description           | Example  |
| ------------- |:---------------------:| --------:|
| server.host | Hostname of ip address of the CallWS server | 127.0.0.1    |
| server.port | Port where the CallWS server API is running |   8989    |
| server.port | CallWS server connection scheme  | http |
| simulation.wipe_on_start | Determines if all calls are deleted on start | true |
| simulation.num_call | Number of random calls generated | 1000 |
| simulation.start_date | Start of call data range | "2020-01-01T00:00:00Z" |
| simulation.end_date | End of call data range | "2020-01-03T00:00:00Z" |
