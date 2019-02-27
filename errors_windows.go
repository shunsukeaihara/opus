package opus

//Error is type of opus error
type Error int

var errorStrings = []string{
	"success",
	"invalid argument",
	"buffer too small",
	"internal error",
	"corrupted stream",
	"request not implemented",
	"invalid state",
	"memory allocation failed",
}

func (e Error) Error() string {
	if int(e) > 0 || int(e) < -7 {
		return "unknown error"
	}
	return errorStrings[-int(e)]
}
