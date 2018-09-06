package validators

//ValidatorInterface interface for all tag validators
type ValidatorInterface interface {
	Validate(html string) bool
}
