# go-rest-mux

Simple Golang REST application with Gorilla Mux

## How to run

These are the steps to run this app:

1. Make sure Golang, DEP Package Manager, and MongoDB are installed
2. Clone this repository to your dir, eg. ```$GOPATH/src/gitlab.com/patricksangian```
3. Go to project root directory (```$GOPATH/src/github.com/patricksangian/go-rest-mux```)
4. Populate the env file ```./env.example``` with your own configuration and copy to ```./env```
5. Run command ```dep ensure``` to install the dependencies
6. Start the app with command ```go run main.go``` or ```make dev```

## Application

>The request header should contain:
```Content-Type: "application/json"```
>The error response should be:

```json
{
  "success": false,
  "message":"Error message",
  "success": false
}
```

>The success response should be:

```json
{
  "success": true,
  "data": "<MULTI DATA TYPE: array, stirng and object>",
  "message":"Success message",
  "code": "<HTTP_SUCCESS_CODE> [200, 201]"
}
```

These are the list of endpoint:

Method       | URI              | Description
------------ | ---------------- | -------------
POST         | /users           | Create new user.
GET          | /users/<:userID> | Get user by ID.
GET          | /users           | Get list of user.

## References

- [YouTube] (<https://www.youtube.com/playlist?list=PLMrwI6jIZn-3a4Hjn-GoYihbMBAzZ6Ae3>)
- [GitHub] (<https://github.com/bxcodec/go-clean-arch>)
