# Armada Management System (ArmadaMS)

## Tech used
- Go (Gin, GORM)
- PostgreSQL 17
- RabbitMQ
- Eclipse Mosquitto (MQTT broker)
- Docker & Docker Compose
- Postman

## Installation
### 1. Clone the repo

```bash
git clone https://github.com/dopamintea/armada_management_system.git
cd armada_management_system
```

### 2. Use Docker Compose
make sure docker is installed; check with
```bash
docker --version
```
if docker exist run
```bash
docker compose up --build
```

This will run;
- API service
- MQTT broker
- PostgreSQL
- RabbitMQ
- PgAdmin
- Mock Publisher (publish location every 2s)
- Worker (geofence alert)

when the system starts, a seed location is published.
you can monitor services with the terminal or docker gui for separate log.

## Use PgAdmin to access published data
- Access dockerized PgAdmin from http://localhost:5050/ or http://127.0.0.1:5050/
- You can put in `admin` for every required input 
- Expand Servers->Armada DB->Databases->armada_db->Schemas->public->Table->Right click on `vehicle_loctions` and view/edit data

## Test API using Postman
- Import the postman collection to test `/ArmadaMS.postman_collection.json`
- API Endpoint
Get last location of :vehicle_id
```bash
GET	/vehicles/:vehicle_id/location
```
Get location of :vehicle_id based on time range
```bash
GET	/vehicles/:vehicle_id/history?from=...&to=...	
```
param used; `"from" #unix timestamp` and `"to" #unix timestamp`

The project is structured with clean code practices in mind, following the SOLID principles for maintainability and scalability.

Created by @dopamintea
