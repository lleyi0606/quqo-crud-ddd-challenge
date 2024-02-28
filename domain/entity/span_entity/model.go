package loggerentity

type Span struct {
	FunctionName string `gorm:"size:255;"`
	Path         string `gorm:"size:255;"`
	Description  string `gorm:"size:255;"`
	Body         string
}
