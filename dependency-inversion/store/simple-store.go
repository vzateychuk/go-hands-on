package store

// SimpleDataStore - простое хранилище данных
type SimpleDataStore struct {
	userData map[string]string
}

func (sds *SimpleDataStore) UserNameForID(userId string) (string, bool) {
	name, ok := sds.userData[userId]
	return name, ok
}

func NewSimpleDataStore() *SimpleDataStore {
	sds := SimpleDataStore{
		userData: map[string]string{
			"fred": "Fred Kirk",
			"mary": "Mary Anderson",
			"vlad": "Vladimir Zateychuk",
		},
	}
	return &sds
}
