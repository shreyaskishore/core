package provider

type FakeStorage struct{}

func (storage *FakeStorage) GetSignedUri(blobKey string, method string) (string, error) {
	return "http://fakestorage.local", nil
}
