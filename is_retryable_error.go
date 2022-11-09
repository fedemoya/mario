package mario

type IsRetryableError interface {
	IsRetryable() bool
}
