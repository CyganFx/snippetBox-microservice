package validator

type Errors struct {
	Errors []string
}

func (e *Errors) Add(message string) {
	e.Errors = append(e.Errors, message)
}
