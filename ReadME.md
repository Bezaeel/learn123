# Learn123

This solution demonstrates multi-module projects using Golang, allowing the flexibility of re-usable modules across the solution

## Project Structure

```
learn123/
├── src/
│   ├── api/
│   │   ├── common/
│   │   ├── modules/
│   │   │   └── course/
│   │   │       ├── v1/
│   │   │       │   └── course.controller.go
│   │   │       └── course.mapper.go
|   |   ├── go.mod
|   |   ├── go.sum
|   |   └── main.go
│   ├── infrastructure/
│   │   └── database/
│   └── learn123.events/
│       └── CourseCreatedEvent.go
|   |
|   |
├── .env.sample
├── .gitignore
```

## Getting Started

### Prerequisites

- Go 1.17 or later
- Docker (optional, for database setup)


### Running the Application

To run the application, use the following command:
```sh
make run
```

## Modules

### API

- **Common**: Contains shared utilities and extensions.
- **Modules**: Contains feature modules like `course`.

### Infrastructure

- **Database**: Contains database-related configurations and utilities.

### Events

- **CourseCreatedEvent**: Defines the `CourseCreated` event structure.

## License

This project is licensed under the MIT License.