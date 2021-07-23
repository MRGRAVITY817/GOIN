package db

import (
	"fmt"
	"os"

	"github.com/MRGRAVITY817/goin/utils"
	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

const (
	dbName       = "blockchain"
	dataBucket   = "data"
	blocksBucket = "blocks"
	checkpoint   = "checkpoint"
)

func getDbName() string {
	port := os.Args[2][6:]
	return fmt.Sprintf("%s_%s.db", dbName, port)
}

// Init database
func DB() *bolt.DB {
	if db == nil {
		dbPointer, err := bolt.Open(getDbName(), 0600, nil)
		db = dbPointer
		utils.HandleErr(err)
		err = db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = t.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleErr(err)
	}
	return db
}

// DB should be closed or it would be corrupted
func Close() {
	DB().Close()
}

// Since bolt allows saving only bytes format, we should get data as bytes
func SaveBlock(hash string, data []byte) {
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

func EmptyBlocks() {
	DB().Update(func(t *bolt.Tx) error {
		utils.HandleErr(t.DeleteBucket([]byte(blocksBucket)))
		_, err := t.CreateBucket([]byte(blocksBucket))
		utils.HandleErr(err)
		return nil
	})
}

func SaveCheckpoint(data []byte) {
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})
	utils.HandleErr(err)
}

func Checkpoint() []byte {
	var data []byte
	DB().View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})
	return data
}

func Block(hash string) []byte {
	var data []byte
	DB().View(func(t *bolt.Tx) error {
		b := t.Bucket([]byte(blocksBucket))
		data = b.Get([]byte(hash))
		return nil
	})
	return data
}
