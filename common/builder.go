package common

type Builder[T any] interface {
	Build() *T
	WithPort(port int) Builder[T]
	WithUsername(username string) Builder[T]
	WithPassword(password string) Builder[T]
	WithImage(image string) Builder[T]
}
