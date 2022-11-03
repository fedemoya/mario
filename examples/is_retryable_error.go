package examples

type IsRetryableError interface {
	IsRetryable() bool
}
