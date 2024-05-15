# Prescription API

This API allows you to manage medical prescriptions, providing endpoints for creating, retrieving, updating, and deleting prescription records. Additionally, it includes authentication functionality.

## Endpoints

### Create a New Prescription

Endpoint: /presc

Method: POST

Description: Adds a new prescription to the system.

### Retrieve a Single Prescription

Endpoint: /presc/{id}

Method: GET

Description: Retrieves details of a specific prescription by its ID.

## Retrieve All Prescriptions for a Patient

Endpoint: /patient/{patient}

Method: GET

Description: Retrieves all prescriptions for a specified patient.

### Retrieve All Prescriptions
Endpoint: /prescs
Method: GET
Description: Retrieves all prescriptions in the system.

### Update a Prescription

Endpoint: /presc/{id}

Method: PATCH

Description: Updates the details of a specific prescription by its ID.

### Delete a Prescription

Endpoint: /presc/{id}

Method: DELETE

Description: Deletes a specific prescription by its ID.

### Authenticate

Endpoint: /authenticate

Method: POST

Description: Authenticates a user and returns a token.

## Testing

Unit and integration tests are implemented for the API. Tests can be run individually or using Docker Compose.

### Run Tests Individually

To run the tests manually, use the following command:

```go test ./...```

### Run Tests with Docker Compose

A docker-compose.yml file is provided to facilitate testing in a containerized environment. To run the tests using Docker Compose, use the following command:

```docker-compose up``` or ```make doker-tests```

This command will build the Docker images, run the tests, and then stop the containers.

## Setup Instructions

### 1. Clone the repository:

```git clone https://github.com/szmulinho/prescription-API.git```
```cd prescription```

### 2. Install dependencies:

```go mod tidy```

### 3. Run the server:

```go run main.go```

## Contribution Guidelines

Fork the repository

Create a new branch ```git checkout -b feature/your-feature-name```

Commit your changes ```git commit -m 'Add some feature'```

Push to the branch ```git push origin feature/your-feature-name```

Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

For any questions or suggestions, feel free to open an issue or contact the repository owner.



