package inMemDB

import (
    "sync"
)


type Database struct {
	PointMap map[string]int
}

var database *Database

var lock = &sync.Mutex{}

func GetInstance() *Database {
    if database == nil {
        lock.Lock()
        defer lock.Unlock()
        if database == nil {
            database = &Database{make(map[string]int)}
        }
    }

    return database
}