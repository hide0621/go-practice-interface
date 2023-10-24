package counterutil

type ByteCounter int

func (bc *ByteCounter) Write(p []byte) (n int, err error) {
	*bc = ByteCounter(len(p))
	return len(p), err
}
