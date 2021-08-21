# Building Web Services using REST APIs
- REST -> Representational State Transfer
- uses HTTP for communicating between the client and the service.
- Several HTTP request methods
    - GET: To retrieve a resource from the service
    - PUT: To create a new resource or update an existing one on the service
    - POST: To create a new resource on the service
    - DELETE: To remove a resource from the service
- `Idempotent methods` (GET, PUT, DELETE) achieve the same result, regardless of how many times the request is repeated.
- can return any type of data, JSON, XML, YAML, or any format that the client can understand.

# HTTP Requests
- HTTP Messages are made up of 
    - `Header`: Contains metadata, such as encoding information, HTTP methods, and so on. The header can contain only plain text. So, you cannot include non-ASCII characters in the header.
    - `Body`: Data to transmit over the network. The body can contain data in any format. The format is specified in the Content-Type field, such as ContentType: `application/json`.

# HTTP URLs
- identifies a resource
- Example - http://www.domain.com/api/v1/cars/model001
    - Domain - www.domain.com
    - API - api
    - Version - v1
    - Service - cars
    - Resource - model001

# REST Response
- When the REST API has received your request, it needs to perform the action that it’s asked to perform
- HTTP response codes
    - `201 Created`: The request was successful and a resource was created. This response code is used to confirm the success of a PUT or POST request.
    - `400 Bad Request`: The request was malformed. This happens especially with POST and PUT requests, when the data doesn’t pass validation or is in the wrong format.
    - `404 Not Found`: The required resource couldn’t be found. This response code is generally returned to all requests that point to a URL with no corresponding resource.
    - `401 Unauthorized`: You need to perform authentication before accessing the resource.
    - `405 Method Not Allowed`: The HTTP method used is not supported for this resource.
    - `409 Conflict`: A conflict has occurred — for example, you’re using a POST request to create the same resource twice.
    - `500 Internal Server Error`: Generally, a 500 response is used when processing fails due to unanticipated circumstances on the server side, causing the server to error out.

# REST API in Go
- package to be used - `gorilla/mux`
- `mux` - HTTP request multiplexer
- The `gorilla/mux` package implements a request router and dispatcher for matching incoming requests to their respective handlers.
- Functions takes two arguments 
    - A `ResponseWriter` object (from the `http` package) so that we can use it to construct an HTTP response to be sent back to the client
    - A pointer to a `Request` object (also from the `http` package) that represents an HTTP request sent by the client and received by the service
- `mux.Vars()` function returns the path variable for the current request. Returns a value of type `map[string]string`
- `r.URL.Query()` function returns the key/value pairs in the query string as a map object.