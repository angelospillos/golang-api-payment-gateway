package main

import (
	"github.com/boltdb/bolt"
	"fmt"
	"time"
	"gateway/model"
	"encoding/json"
)

var rootBucket = []byte("DB")
var accountBucket = []byte("ACCOUNT")
var paymentBucket = []byte("PAYMENT")
var businessStatementBucket = []byte("BUSINESS_STATEMENT")
var personalStatementBucket = []byte("PERSONAL_STATEMENT")

func setupDB() (*bolt.DB, error) {
	db, err := bolt.Open("src/gateway/main.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists(rootBucket)
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists(accountBucket)
		if err != nil {
			return fmt.Errorf("could not create Account bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists(paymentBucket)
		if err != nil {
			return fmt.Errorf("could not create Payment bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	fmt.Println("DB Setup Done")
	return db, nil
}

func saveAccount(db *bolt.DB, account model.Account) error {

	var key = fmt.Sprintf("%v", account.Id)
	var value, err = json.Marshal(account)

	if err != nil {
		return fmt.Errorf("could not json marshal entry: %v", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket(rootBucket).Bucket(accountBucket).Put([]byte(key), []byte(value))
		if err != nil {
			return fmt.Errorf("could not insert entry: %v", err)
		}

		return nil
	})

	fmt.Println("Added Account Entry " + account.Id)

	return err
}

func getAccount(db *bolt.DB, id string) (model.Account, error) {

	var account model.Account
	var key = []byte(id)

	err := db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket(rootBucket).Bucket(accountBucket).Get(key)
		fmt.Sprintf("Found Account Entry %s", value)
		return json.Unmarshal(value, &account)
	})

	if err != nil {
		fmt.Printf("Could not get Account ID %s", id)
		return account, nil
	}

	return account, nil
}

func savePayment(db *bolt.DB, payment model.Payment) error {

	var key = fmt.Sprintf("%v", payment.Id)

	var value, err = json.Marshal(payment)

	if err != nil {
		return fmt.Errorf("could not json marshal entry: %v", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket(rootBucket).Bucket(paymentBucket).Put([]byte(key), []byte(value))
		if err != nil {
			return fmt.Errorf("could not insert payment entry: %v", err)
		}

		return nil
	})

	fmt.Println("Added Payment Entry " + key)

	return err
}

func getPayment(db *bolt.DB, id string) (model.Payment, error) {

	var key = []byte(id)
	var payment model.Payment

	err := db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket(rootBucket).Bucket(paymentBucket).Get(key)
		fmt.Sprintf("Found Payment Entry %s", value)
		return json.Unmarshal(value, &payment)
	})

	if err != nil {
		fmt.Printf("Could not get Payment ID %s", id)
		return payment, nil
	}

	return payment, nil
}
