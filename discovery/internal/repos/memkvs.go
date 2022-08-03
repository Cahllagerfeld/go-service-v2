package repos

type memkvs struct {
	kvs map[string][]byte
}

func NewMemKvs() *memkvs {
	return &memkvs{
		kvs: map[string][]byte{},
	}
}
