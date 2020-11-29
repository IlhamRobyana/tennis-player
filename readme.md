# tennis-player
Built with Golang, PostgreSQL, Echo, Bcrypt, and other third-party libraries.
## Installation
Make sure Golang and is already installed.
Clone the repo to the directory ```{user}/go/src/github.com/ilhamrobyana/tennis-player```
Install the dependendices with ```dep ensure```
## Peresquite
Before running the application, set the dot env file first using .env.example as the layout.

PostgreSQL needed to be setup first.

And then migrate the data structure to PostgreSQL

```cd migration```

```go run main.go```

## Running
Go to the root folder of the project and enter the following:
```go run auth.go```

## Documentation
The API documentation can be found in the following link:
```https://documenter.getpostman.com/view/8102951/TVmJhya8```

## Testing
There's some unit testing that has been implemented, namely login and putball tests, could be executed in pg_storage folder