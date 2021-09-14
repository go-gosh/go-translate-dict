package api

type Option func(param map[string]string) error

type LanguageType = string

// WithFrom set from language
// TODO check from.
func WithFrom(from LanguageType) Option {
	return func(param map[string]string) error {
		param[FromLangField] = from
		return nil
	}
}

// WithTo set to language
// TODO check to.
func WithTo(to LanguageType) Option {
	return func(param map[string]string) error {
		param[ToLangField] = to
		return nil
	}
}
