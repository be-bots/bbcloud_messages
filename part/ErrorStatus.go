package part

// ErrorStatus tell you if an error ocurred and if yes, it tells you the description of the error
type ErrorStatus struct {
	ErrorOcurred bool
	Error        string
}
