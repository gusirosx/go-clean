# go clean
GoLang clean achitecture demonstration in a REST-API

Project Description&Structure:
REST API with custom JWT-based authentication system. Core functionality is about creating and managing bookmarks (Simple clone of Pocket).

Structure:
4 Domain layers:
- Infrastructure layer
- Adapters layer
- Use Cases layer
- Domain layer


Let’s sum this all up by creating a list of all layers and the parts of our software therein:

- Domain:
  - user model
Use Cases:
User entity
Use case: Add Item to Order
Use case: Get Items in Order
Use case: Admin adds Item to Order
Interfaces:
Web Services for Item/Order handling
Repositories for Use Cases and Domain entities persistence
- Infrastructure:
  - The Database
  - Code that handles DB connections
  - The HTTP server

One last thought before we dive into the code. If we look at how we separated our application, there’s a pattern here. If you look at the several layers and align them along the dimensions of how application-specific the contained code is and how business-specific it is, the pattern becomes apparent:

Infrastructure | Interfaces| Use Cases | Domain 
--- | --- | --- | --- 
application-agnostic | application-specific | application-specific | application-agnostic 
business-agnostic | business-agnostic | business-specific | business-specific 

```ssh
.
├── adapters
│   ├── controller.go
│   ├── presenter.go
│   └── repository.go
├── domain
│   └── models
│       ├── errors.go
│       └── user.go
├── infrastructure
│   ├── database
│   │   └── postgres.go
│   └── router
│       └── router.go
├── registry
│   ├── registry.go
│   └── userRegistry.go
├── usecase
│   ├── interactor.go
│   ├── presenter.go
│   └── repository.go
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```


# Simple API Clean Architecture With Gin and Gorm Framework
CRUD simple RESTFUL API with routing in using gin framework and ORM database with Gorm auto migrate. Database using mysql. Check config/config.go for configurate database. just create a database with the name simple_api in MySQL and then run this API program 

## Clean architecture divides our code into 4 layers
- Model/entity: encapsulate enterprise wide business rules. An entity in Go is a set of data structures and functions.
- Use Cases: the software in this layer contains application specific business rules. It encapsulates and implements all of the use cases of the system.
- Repository: the software in this layer is a set of adapters that convert data from the format most convenient for the use cases and entities, to the format most convenient for some external agency such as the Database or the Web
- Handler/Driver: this layer is generally composed of frameworks and tools such as the Database, the Web Framework, etc.


Domain
Entities encapsulate the most general and high-level business rules.
They are the least likely to change when something external changes.


Use Case:
The use case layer contains application specific business rules.

The use case layer only depends on the entities. change in the entities will require a change in the use case layer, but changes to other layers won’t.


Interface Adapters:
Interface adapters layer contains a set of adapters that convert data from the format most convenient for the use cases and entities, to the format most convenient for some external agency such as the Database or the Web.


Frameworks and Adapters:
This layer is where all the details go: the Web is a detail and so is the database.

In the Frameworks and Adapters layer we have the Database, UI, etc.
This is the most external level, and is bound to change more frequently than other circles. However, changes here should not affect inner circles.