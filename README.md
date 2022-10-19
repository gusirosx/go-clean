# go-clean
GoLang clean achitecture demonstration in a REST-API


Rules of Clean Architecture by Uncle Bob:
Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
Idependent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.
More on topic can be found here.

Project Description&Structure:
REST API with custom JWT-based authentication system. Core functionality is about creating and managing bookmarks (Simple clone of Pocket).

Structure:
4 Domain layers:



Infrastructure layer
Adapters layer
Use Cases layer
Domain layer


Let’s sum this all up by creating a list of all layers and the parts of our software therein:

Domain:
Customer entity
Item entity
Order entity
Use Cases:
User entity
Use case: Add Item to Order
Use case: Get Items in Order
Use case: Admin adds Item to Order
Interfaces:
Web Services for Item/Order handling
Repositories for Use Cases and Domain entities persistence
Infrastructure:
The Database
Code that handles DB connections
The HTTP server
Go Standard Library



Infrastructure	Interfaces	Use Cases	Domain
application-agnostic	application-specific	application-specific	application-agnostic
business-agnostic	business-agnostic	business-specific	business-specific


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