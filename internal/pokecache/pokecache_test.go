package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 30)
	if cache.mapCache == nil {
		t.Error("Cache not created")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 30)

	cases := []struct {
		inputKey string
		inputVal []byte
	} {
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, c := range cases {
		cache.Add(c.inputKey, c.inputVal)
		actual, ok := cache.Get(c.inputKey)
		if !ok {
			t.Errorf("%s not found", c.inputKey)
		}
		if string(actual) != string(c.inputVal) {
			t.Errorf("%s does not match %s", string(actual), string(c.inputVal))
		}
	}
}

func TestReap(t *testing.T) {
	interval := 10 * time.Millisecond
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))
	time.Sleep(interval + 5 * time.Millisecond)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s should have been reaped", keyOne)
		return
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond *10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("val1"))
	time.Sleep(interval / 2)

	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("%s should not have been reaped", keyOne)
		return
	}

}
