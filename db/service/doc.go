/*
Package service

This package contains DBService interface which defines a set of methods to interact with the database:

  - SaveUrl creates a models.Url instance from (req SaveUrlRequest) and saves it into the database
    returns an error if any

  - GetOriginalUrl returns the original url of a shortened url from the database
    returns an error if any

Each of these methods has its own request and response type. For example SaveUrl receives
SaveUrlRequest as input and returns SaveUrlResponse

Request and response structs are designed in a way that once an instance is created,
its fields cannot be changed.
To Access instance fields, use getter functions
To make instances of request and response structs you need to use the New function
related to that request or response
*/
package service
