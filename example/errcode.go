//go:generate  gen-const-msg -t=Code -m

package example

type Code int

const (
	// ErrParams err params
	ErrParams Code = 400
	// ErrServer Internal Server Error
	ErrServer Code = 500
	// ErrCodeHasNotExist has not exist
	ErrCodeHasNotExist Code = 501
)
