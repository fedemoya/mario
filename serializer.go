package mario

type Serializer interface {
	Serialize() ([]byte, error)
}
