package database

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/wallanaq/jencli/config"
	"go.etcd.io/bbolt"
)

var db *bbolt.DB

func Open() error {

	configDir, err := config.GetConfigDir()

	if err != nil {
		return fmt.Errorf("failed to get configuration directory: %w", err)
	}

	dbPath := filepath.Join(configDir, "db")

	if db, err = bbolt.Open(dbPath, 0600, &bbolt.Options{Timeout: 1 * time.Second}); err != nil {
		return fmt.Errorf("")
	}

	return nil
}

func Close() error {

	if db == nil {
		return fmt.Errorf("database is not open")
	}

	return db.Close()

}

func Set(bucketName string, key string, value string) error {

	return db.Update(func(tx *bbolt.Tx) error {

		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))

		if err != nil {
			return fmt.Errorf("failed to create or retrieve bucket %s: %w", bucketName, err)
		}

		if err := bucket.Put([]byte(key), []byte(value)); err != nil {
			return fmt.Errorf("failed to set key %s in bucket %s: %w", key, bucketName, err)
		}

		return nil

	})

}

func Get(bucketName string, key string) (value string, err error) {

	err = db.View(func(tx *bbolt.Tx) error {

		bucket := tx.Bucket([]byte(bucketName))

		if bucket == nil {
			return fmt.Errorf("bucket %s not found", bucketName)
		}

		bValue := bucket.Get([]byte(key))

		if bValue == nil {
			return fmt.Errorf("key %s not found in bucket %s", key, bucketName)
		}

		value = string(bValue)

		return nil

	})

	if err != nil {
		return "", fmt.Errorf("failed to get key %s from bucket %s: %w", key, bucketName, err)
	}

	return
}

func Delete(bucketName string, key string) error {

	return db.Update(func(tx *bbolt.Tx) error {

		bucket := tx.Bucket([]byte(bucketName))

		if bucket == nil {
			return fmt.Errorf("bucket %s not found", bucketName)
		}

		if err := bucket.Delete([]byte(key)); err != nil {
			return fmt.Errorf("failed to delete key %s from bucket %s: %w", key, bucketName, err)
		}

		return nil

	})

}
