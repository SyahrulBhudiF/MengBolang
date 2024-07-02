package main

type validationError struct {
	message string
}

func (err *validationError) Error() string {
	return err.message
}

type notFoundError struct {
	Message string
}

func (err *notFoundError) Error() string {
	return err.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"ID is required"}
	}

	if data == nil {
		return &validationError{"Data is required"}
	}

	if data != "eko" {
		return &notFoundError{"Data not found"}
	}

	// process save data
	return nil
}

func main() {
	err := SaveData("1", "eko")
	if err != nil {
		switch err.(type) {
		case *validationError:
			println("Validation Error", err.Error())
		case *notFoundError:
			println("Not Found Error", err.Error())
		default:
			println("Unknown Error", err.Error())
		}
	}
}
