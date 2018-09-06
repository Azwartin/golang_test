package readers

//ReaderInterface is interface for input data readers
type ReaderInterface interface {
	Read(from string, params ...interface{}) (string, error)
}
