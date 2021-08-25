package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {

	// Создаем клиента для подключения к memcache (ожидается на 127.0.0.1:11211)
	MemcachedAddresses := []string{"127.0.0.1:11211"}
	memcacheClient := memcache.New(MemcachedAddresses...)

	// Upsert записи в cache
	mkey := "key"
	memcacheClient.Set(&memcache.Item{
		Key:        mkey,
		Value:      []byte("1"),
		Expiration: 3, // expiration time, in seconds
	})

	// ???
	memcacheClient.Increment(mkey, 1)

	// Получить из memcache значение по ключу
	item, err := memcacheClient.Get(mkey)
	// Признак что запись не найдена в err = memcache.ErrCacheMiss
	if err != nil && err != memcache.ErrCacheMiss {
		fmt.Println("MC error", err)
	}

	fmt.Printf("mc value %#v\n", item)

	// Удаление из кэша
	memcacheClient.Delete(mkey)
	item, err = memcacheClient.Get(mkey)
	if err == memcache.ErrCacheMiss {
		fmt.Println("record not found in MC")
	}
}
