
# Live Coding Challenge

## Description
This project is a Go-based application that follows **Hexagonal Architecture** (also known as Ports and Adapters). The core business logic is decoupled from the infrastructure and external systems like the database and file reader. This architecture allows for easier testing, maintainability, and flexibility when integrating with external systems.

The application processes user data by reading it from a JSON file, storing it in a PostgreSQL database, and providing APIs to interact with the data. It leverages **Wire** for dependency injection and **Swagger** for API documentation.

## Hexagonal Architecture
The application follows the **Hexagonal Architecture**, where:
- **The core domain** contains the business logic and entities.
- **Adapters** handle external interactions such as reading/writing to a database, API calls, or file systems.
- **Ports** define the interfaces for the core domain to interact with the outside world.

### Benefits of Hexagonal Architecture:
- **Separation of concerns**: Keeps business logic isolated from infrastructure details, making it easier to test and modify.
- **Flexibility**: Allows easy replacement of external systems (e.g., switching from one database to another) without affecting the core domain.
- **Testability**: Simplifies unit testing since the core domain can be tested independently of the infrastructure.

## Installation

### Step 1: Set up PostgreSQL
This application uses PostgreSQL as the database. You can set up the database using the provided **Docker Compose** file.

```bash
docker compose up -d
```

This will start PostgreSQL on the default port `5432`.

## 2. Run Migrations

To create the necessary tables in the PostgreSQL database, run the provided SQL migration files located in the `database/migration` folder.

Make sure to run all the required migration files to set up the database schema correctly.

### Step 3: Configure `app.yaml`
After setting up PostgreSQL, you may need to adjust the database configuration in the `app.yaml` file, specifically for the database port. If you changed the port in the Docker Compose configuration or if PostgreSQL is running on a custom port, update the `app.yaml` accordingly.

Example configuration in `app.yaml`:

```yaml
database:
  host: "localhost"
  port: 5432
  username: "devuser"
  password: "devpass"
  name: "live_coding"
  
worker:
  user_worker_count: 10  # Number of workers to process user data
```

### Step 4: Generate Wire Dependencies
Wire is used for dependency injection in this project. If the `wire_gen.go` file does not exist, generate it using the following command:

```bash
make wire
```

This will generate the necessary dependency files to wire up the application.

### Step 5: Generate Swagger Documentation
The application uses **Swagger** for API documentation. To generate the Swagger documentation, run:

```bash
make swagger
```

Once generated, the Scalar UI is accessible at `http://localhost:8080/reference` (assuming you haven't changed the port in `app.yaml`).

## Running the Application

You can run the application using either of the following commands:

```bash
go run cmd/main.go http
```

or

```bash
make run
```

This will start the HTTP server, and the application will be available at `http://localhost:8080`.

## API Endpoints

## Base URL
base url for all the requests is /api/v1 which is configurable in app.yaml file

### `/reference`

- **Method**: GET
- **Description**: Provides the API documentation via Swagger UI.
- **Access**: Available at `http://localhost:8080/reference`.

### `/user/{id}`

- **Method**: GET
- **Description**: Retrieves a user by their UUID.

### `/user/ingest`

- **Method**: POST
- **Description**: Creates users from a JSON file. The file path is configured in `app.yaml`.

## Configuration

All application settings, including the database configuration and file reader count, are managed in the `app.yaml` file. You can adjust the settings here before running the application.



## Database Setup

The project uses PostgreSQL as the database. The tables for `users` and `addresses` are created using SQL migrations. Ensure that the PostgreSQL server is up and running.

## Coding Interview Code Restrictions

This code is provided as part of a coding interview and is for evaluation purposes only.

## Restrictions:
- Redistribution, modification, and commercial use are **prohibited**.
- The code is **not** intended for production use or distribution.
- You may **view** and **evaluate** the code as part of the interview process.
- For any further use or inquiries, please contact m.hesam.khorshidi@gmail.com.

By using or reviewing this code, you acknowledge and agree to these restrictions.
