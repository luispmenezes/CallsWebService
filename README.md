# CallsWebService
Simple service to manage call records and extract statistics written in GO and using a Postgres Database.

## Requirements
1. Have **docker**, **docker compose**, **go(local only)**  installed
- [Install Docker](https://docs.docker.com/install/)
- [Install Docker Compose](https://docs.docker.com/compose/install/)
- [Install Go](https://golang.org/doc/install)
2. Have docker running.
3. Have port 8989 (and 5432 if running locally) on your system available

## Run Locally
```bash
cd build
docker-compose up
``` 

## Run Locally
```bash
cd build
./local-build.sh
./local-start.sh
``` 

## API
Details information about the API can be viewed in the **swagger.yaml** file and a postman Collection is also available in **CallWebService.postman_collection.json**  

## Configuration
Example config file can be found in **build/config.json**

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
