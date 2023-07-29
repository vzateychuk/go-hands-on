package main

// MyDataStore - потребуется хранилище данных, находящее пользователя по ID
type MyDataStore interface {
	UserNameForID(userID string) (string, bool)
}
