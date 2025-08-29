<!-- go install github.com/swaggo/swag/cmd/swag@latest
export PATH=$PATH:$(go env GOPATH)/bin
source ~/.bashrc  # ou source ~/.zshrc
swag init --generalInfo cmd/main.go --output docs
Gorm -->

## Application Description
<p> This is an aplication to manager tasks</p>

## Tools

- [x]   go lang 1.23.0
- [x]   mysql
- [x]   gorm
- [x]   toolchain go1.24.6

## Entities

- [x]   users
- [x]   tasks
- [x]   roles
- [x]   permissions
- [x]   user_roles
- [x]   role_permissions

## Migrations and seeders
<p> The migration with DDL SQL and the seeders to populate the tables are generated automatically when you run the application</p>
## Diagram


## Layer Architecture

<p> The primary goal of this pattern is to separate concerns, making your codebase more organized, testable, and easier to manage as it grows. Here's a breakdown of why each layer is important:</p>

* Model: This layer represents your data structures. It defines the schema and business objects used throughout the application. By centralizing models, you ensure data consistency and provide a clear blueprint for what your application manipulates.

* Repository: This layer is responsible for data persistence. It contains the logic for interacting with your database, external APIs, or any other data source. By abstracting this logic behind an interface, you can easily swap out your database (e.g., from SQL to a NoSQL database) without affecting the other layers

* Service: This is the core of your business logic. The service layer orchestrates interactions between the repository and the handler. It performs tasks like data validation, business calculations, and calling multiple repositories to fulfill a request. It should not know about the HTTP details, making it highly reusable and easy to test.

* Handler: Also known as the controller or HTTP handler, this layer is the entry point for incoming requests. It handles HTTP-specific tasks like parsing requests, validating input data, and formatting responses. It calls the service layer to perform the business logic and then sends the response back to the client.

<p>Key Benefits of this Architecture</p>

* Testability: Since each layer has a single responsibility, it becomes much easier to write unit tests. You can test your service layer by providing mock repositories, without needing a real database connection. This leads to faster and more reliable tests.

* Maintainability: When you need to fix a bug or add a new feature, you know exactly where to go. A bug in data saving? Look at the repository. A change in business rules? The service layer is your target.

* Scalability: The clear separation of responsibilities allows different teams to work on different layers simultaneously without stepping on each other's toes. This is crucial for larger projects.

* Flexibility: The use of interfaces for the repository and service layers makes your code very flexible. You can create different implementations for different environments (e.g., a mock repository for local development and a real one for production).

<p>
Potential Downsides </br>
While this architecture is excellent, it's important to be aware of potential issues, especially for smaller projects:
</p>

* Initial Overhead: For a very simple application, this pattern can feel like overkill. It requires more files and interfaces, which might seem unnecessary at first.

* Increased Complexity: As with any pattern, if not implemented correctly, it can lead to over-engineering. It's vital to keep the layers focused on their specific roles and avoid mixing responsibilities.


## Run application
```bash

# run mysql database
docker compose up -d mysql

# run
chmod +x scripts/generateDocs.sh

# generate docs if you need
# scripts/generateDocs.sh


# run
go run cmd/main.go

# access swagger
http://localhost:${PORT}/swagger/

# Docker compose
# run
docker compose   --env-file .env  up -d

````

## Run tests
```bash
# execute
export PATH="$PATH:$(go env GOPATH)/bin"
## Run if you need
# mockgen -source=./modules/tasks/repository/task_repository.go -destination=./modules/tasks/repository/mocks/mock_task_repository.go

# acceess the folders
# cd modules/tasks/repository 
# cd modules/tasks/services
# cd modules/users/repository 
go test
go 
```