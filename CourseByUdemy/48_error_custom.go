package main

import "fmt"

// struct for custom error
type validationError struct {
	Message string
}

// method for custom error
// wajiib ada method Error() string, karena error interface membutuhkan method Error() string
func (v *validationError) Error() string { // pointer method bukan pointer function, karena terikat dengan struct
	return v.Message
}

// struct for custom error
type notFoundError struct {
	Message string
}

// method for custom error
// wajiib ada method Error() string, karena error interface membutuhkan method Error() string
func (n *notFoundError) Error() string { // pointer method bukan pointer function
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "Validation Error, ID is required"}
	} 

	if data != "Tri" {
		return &notFoundError{Message: "Data not found"}
	}

	return nil
}

func main() {
	err := SaveData("1", nil)

	if err != nil {
		switch finalError := err.(type) {

		case *validationError:
			fmt.Println("validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("Unknown error:", finalError.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}