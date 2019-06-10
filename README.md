# go-rest-mux

Simple Golang REST application with Gorilla Mux

## How to run

These are the steps to run this app:

1. Make sure Golang, DEP Package Manager, and MongoDB are installed
2. Create new project on sentry.io for your logger
3. Clone this repository to your dir, eg. ```$GOPATH/src/github.com/sangianpatrick```
4. Go to project root directory (```$GOPATH/src/github.com/sangianpatrick/go-rest-mux```)
5. Populate the env file ```./.env.example``` with your own configuration and copy to ```./.env```
6. Run command ```dep ensure``` to install the dependencies
7. Run command ```make test``` to run the unit testing
8. Run command ```make dev``` to run app on development environment or
9. Run command ```make run``` to create and run from app's executable

## Application

### Accepted Request Header

```{ Authorization: "Basic <token> }"``` or  ```{ Authorization: "Bearer <token>" }``` and ```{ Content-Type: "application/json" }```

### Error Response

```json
{
  "success": false,
  "message":"Error message",
  "code": "<HTTP_ERROR_CODE> [400,401,404,409]"
}
```

### Success Response

```json
{
  "success": true,
  "data": "<MULTI DATA TYPE: array, string, int and object>",
  "message":"Success message",
  "code": "<HTTP_SUCCESS_CODE> [200, 201]"
}
```

### Success Pagination Response

```json
{
  "success": true,
  "data": "<DATA TYPE: array>",
  "message":"Success message",
  "code": "<HTTP_SUCCESS_CODE> [200]",
  "meta": {
    "totalPage": 20,
    "page": 1,
    "totalData": 200,
    "totalDataOnPage": 10
  }
}
```

### List of Endpoint

Method       | Authorization          | URI                          | Description
------------ | ---------------------- | ---------------------------- | -------------
POST         | Basic {token}          | /api/v1/auth/signin          | User signin.
POST         | Basic {token}          | /api/v1/users/registration   | Create new user.
GET          | Bearer {token}         | /api/v1/users/profile/me     | Get my profile.
GET          | Bearer {token}         | /api/v1/users?page=1&size=10 | Get list of user depends on page and size.
GET          | Bearer {token}         | /api/v1/users/{userID}       | Get one user with ID.
PUT          | Bearer {token}         | /api/v1/users/profile/me     | Update authorized user property.
DELETE       | Bearer {token}         | /api/v1/users/profile/me     | Delete authorized user account.
POST         | Bearer {token}         | /api/v1/users/article        | Create authorized user's article.

## References

- [YouTube] (<https://www.youtube.com/playlist?list=PLMrwI6jIZn-3a4Hjn-GoYihbMBAzZ6Ae3>)
- [GitHub] (<https://github.com/bxcodec/go-clean-arch>)
