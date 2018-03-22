# sbank

sbank is a package for accessing the [Bank](https://api.sbanken.no/Bank/swagger) and [Customer](https://api.sbanken.no/Customers/swagger) APIs from sbanken.

## Getting started
Before using this package, you will need a developer account at [sbanken](https://sbanken.no/bruke/utviklerportalen/)

After registering, you will need to create an application and generate a password which will be used when authenticating with the API.
  
## Installation

```bash
> go get -u github.com/torie/sbank
```

## Usage

### Authentication
Authentication can either be handled by the package:

```go
c := sbank.New("client-id", "client-secret")
```

Or an `http.Client` which handles authentication can be provided. 

```go
client := MyHTTPClient()
c := sbank.NewWithClient(client)
```

### Examples
Examples can be found in the [examples folder](examples)
