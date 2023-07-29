package main

// MyDataStore - потребуется хранилище данных, находящее пользователя по ID
type MyDataStore interface {
	UserNameForID(userID string) (string, bool)
}

type DataStoreAdapter func(userId string) (string, bool)

func (ds DataStoreAdapter) UserNameForID(userId string) (string, bool) {
	return ds(userId)
}
