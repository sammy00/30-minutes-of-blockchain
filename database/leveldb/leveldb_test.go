package leveldb_test

import (
	"bytes"
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

func mockUpDB() *leveldb.DB {
	// db is thread-safe
	db, err := leveldb.OpenFile("db", nil)
	if nil != err {
		return nil
	}

	k1 := []byte("key1")
	v1 := []byte("value1")

	k2 := []byte("key2")
	v2 := []byte("value2")

	// insert (k,v)
	if err := db.Put(k1, v1, nil); nil != err {
		return nil
	}
	if err := db.Put(k2, v2, nil); nil != err {
		return nil
	}

	return db
}

func TestReadUpdate(t *testing.T) {
	db := mockUpDB()
	if nil == db {
		t.Fatal("failed to populate db")
	}
	defer db.Close()

	k1, v1 := []byte("key1"), []byte("value1")

	// query v by k: the return v should be read only
	if v, err := db.Get(k1, nil); nil != err {
		t.Fatal(err)
	} else if !bytes.Equal(v1, v) {
		t.Fatalf("invalid value: want %x, got %x", v1, v)
	}

	// iteration
	itr := db.NewIterator(nil, nil)
	for itr.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		t.Logf("(%x,%x)\n", itr.Key(), itr.Value())
	}
	itr.Release()
	if err := itr.Error(); nil != err {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	db := mockUpDB()
	if nil == db {
		t.Fatal("fail to populate db")
	}

	k1 := []byte("key1")
	if err := db.Delete(k1, nil); nil != err {
		t.Fatal(err)
	}
	t.Log("entries bound to k1 is removed")
	if _, err := db.Get(k1, nil); nil == err {
		t.Fatal("no value should be found")
	}
}
