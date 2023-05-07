package errs

import "fmt"

// Errors with buses

var BusStorageOverflowError = fmt.Errorf("bus with such ID already exists! Probably the storage if full")
var BusNotExistsError = fmt.Errorf("bus with such ID does not exist")

// Errors with cities

var CityStorageOverflowError = fmt.Errorf("city with such ID already exists! Probably the storage if full")
var CityNotExistsError = fmt.Errorf("city with such ID does not exist")
