package packager

type FileHeader struct {
	Magic   string
	Version byte
	Key     []byte
	iv      []byte
}
