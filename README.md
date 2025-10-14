# Pet Management API

Go REST API for managing pets and their medical events with MySQL.

## Quick Start


### Prerequisites
- Docker and Docker Compose
- Go 1.19+ (for local development)

### Installation

1. **Clone the repository**
```bash
# Clone and setup
git clone https://github.com/om00/menagerie.git
cd menagerie
 ```  #

2. **Setup environment**  # ← Proper step number
   ```bash
   cp .env.example .env
   ```
 update the credientials as mention env.example

# Start with Docker
docker-compose up -d --build

After this, the Go app will be running on APP_PORT and MySQL can be accessed on DB_PORT as configured in your .env file.

## API Examples


### Get All Pets List
```bash
curl --location 'http://localhost:8087/pets'
```  
Repace 8087 with app_port

### Create Pet Api
```bash
curl --location 'http://localhost:8087/pets' \
--header 'Content-Type: application/json' \
--data '{
  "name": "Coco",
  "owner": "Robert Taylor",
  "species": "Rabbit",
  "birth": "2023-01-10T00:00:00Z"
}'
```
### Update Pet Api
```bash
curl --location --request PUT 'http://localhost:8087/pets/1' \
--header 'Content-Type: application/json' \
--data '{
  "name": "BuddyOne"
}'
```

### Delete Pet Api
```bash
curl --location --request DELETE 'http://localhost:8087/pets/4'
```

### Create Pet Event Api
```bash
curl --location 'http://localhost:8087/pets/4' \
--header 'Content-Type: application/json' \
--data '{
  "date": "2024-01-20T14:30:00Z",
  "type": "vaccination",
  "remark": "ISO format with Zulu time"
}'
```
### Get Pet Events Api
```bash
curl --location 'http://localhost:8087/pets/4'
```