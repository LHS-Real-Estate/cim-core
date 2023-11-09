package valueobjects

type Document struct {
	FilePath  string `validate:"required,filepath"`
	Extension string `validate:"required,lowercase,min=2"`
}
