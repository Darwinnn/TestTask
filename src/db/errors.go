package db

type errNegativeBalance struct{}
type errUUIDExists struct{}

var (
	// ErrNegativeBalance is returned when decrement value is greater than balance value (which would lead to negative balance value)
	ErrNegativeBalance errNegativeBalance
	// ErrUUIDExists is returned when record with this UUID already exists
	ErrUUIDExists errUUIDExists
)

func (errNegativeBalance) Error() string {
	return "Subtracting balance would lead to negative value"
}
func (errUUIDExists) Error() string {
	return "uuid already exists in the database"
}
