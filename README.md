# Carijajan Backend (Go + Fiber)

This is the backend service for the [carijajan](https://carijajan.vercel.app/), built using Go and the Fiber web framework.

## Features
- RESTful API endpoints
- Database integration
- Modular architecture (controllers, services, routes, etc.)
- Environment configuration support

## Prerequisites
- Go 1.20 or later
- PostgreSQL database

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/carijajan-backend-go.git
   cd carijajan-backend-go
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up the environment variables:
   - Copy the `example.env` file to `.env`:
     ```bash
     cp example.env .env
     ```
   - Update the `.env` file with your database credentials and other configurations.

4. Run the application:
   ```bash
   go run main.go
   ```

## Project Structure
```
carijajan-backend-go/
├── controllers.go   # Handles HTTP requests
├── database.go      # Database connection setup
├── example.env      # Example environment variables
├── go.mod           # Go module file
├── go.sum           # Go dependencies checksum
├── main.go          # Entry point of the application
├── models.go        # Data models
├── routes.go        # API route definitions
├── services.go      # Business logic
├── tmp/             # Temporary files (ignored by Git)
```

## API Documentation

### Example Endpoint
- **GET** `/example`
  - Description: Example endpoint
  - Response:
    ```json
    {
      "message": "Hello, World!"
    }
    ```

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request.

## License
This project is licensed under the MIT License.