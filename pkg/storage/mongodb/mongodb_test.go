package mongodb

import "testing"

func TestStorage_MaxID(t *testing.T) {
	db, err := New("mongodb://localhost:27017/")
	if err != nil {
		t.Error(err)
		return
	}
	res, err := db.maxID()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}
