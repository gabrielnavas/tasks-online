package repositories

import "errors"

var ErrResourceNotFound = errors.New("resource not found")
var ErrConnectionDone = errors.New("connection is done")
