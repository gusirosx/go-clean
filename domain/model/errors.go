package model

import "errors"

var ErrNotFound = errors.New("User not found")
var ErrListing = errors.New("error occured while listing user items")
var ErrNoID = errors.New("no user ID was provided")
var ErrQuery = errors.New("unable to execute the query: ")
