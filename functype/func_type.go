package functype

type FunctionType string

const (
	Unknown FunctionType = "unknown"
	Lambda  FunctionType = "lambda"
)

func Of(value string) (FunctionType, bool) {
	switch value {
	case string(Lambda):
		return Lambda, true
	default:
		return Unknown, false
	}
}
