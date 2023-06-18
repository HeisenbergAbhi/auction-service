# Auction Service

The Auction Service is a web application that facilitates auctions for ad spaces. It consists of two sub-modules: the Supply Side Service, which lists available ad spaces, and the Demand Side Service, which manages bids from interested bidders.

## Technologies Used

- Go programming language
- MySQL database
- Docker
- Docker Compose

## Prerequisites

Make sure you have the following software installed on your machine:

- Docker: [Installation Guide](https://docs.docker.com/get-docker/)
- Docker Compose: [Installation Guide](https://docs.docker.com/compose/install/)

## Getting Started

Follow the steps below to set up and run the Auction Service:

1. Clone this repository to your local machine:

git clone https://github.com/HeisenbergAbhi/auction-service.git

2. Navigate to the project directory:

cd auction-service

3. Start the services using Docker Compose:

docker-compose up

4. Wait for the containers to start up. You should see logs indicating the successful startup of the services.

5. Once the services are up and running, you can access them using the following URLs:

- Supply Side Service: [http://localhost:8000](http://localhost:8000)
- Demand Side Service: [http://localhost:8001](http://localhost:8001)

## API Endpoints

The following API endpoints are available:

### Supply Side Service

- **GET /ad-spaces**: Retrieve a list of available ad spaces.
- **POST /ad-spaces**: Create a new ad space.

### Demand Side Service

- **GET /bids/ad-space/{adSpaceID}**: Retrieve bids for a specific ad space.
- **GET /bids/ad-space/{adSpaceID}/winning**: Retrieve the winning bid for a specific ad space.
- **POST /bids**: Place a bid for an ad space.

Refer to the source code for detailed information about the request and response payloads.

## Running Tests

To run the tests for the Auction Service, execute the following command in the project directory:

go test ./...

## Database Configuration

The Auction Service uses a MySQL database for data storage. The database connection details can be configured in the `pkg/database/mysql.go` file. Update the `dsn` variable with your MySQL connection details.