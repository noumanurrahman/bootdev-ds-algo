package structures

import (
	"log"
	"unicode/utf8"
)

type HashMap[T any] struct {
	size   int
	list   []*T
	bucket []string
}

func NewHashMap[T any](size int) *HashMap[T] {
	hashmap := &HashMap[T]{size: size}
	hashmap.list = make([]*T, size)
	return hashmap
}

func HashToInt(hash string, size int) int {
	sum := 0
	for char := range hash {
		code, _ := utf8.DecodeRuneInString(hash[char:])
		sum += int(code)
	}
	return sum % size
}

func (hashmap *HashMap[T]) Set(key string, value T) {
	index := HashToInt(key, hashmap.size)
	hashmap.list[index] = &value
	hashmap.bucket = append(hashmap.bucket, key)
}

func (hashmap *HashMap[T]) Get(key string) T {
	isKeyInBucket := false
	for keyBucket := range hashmap.bucket {
		if key == hashmap.bucket[keyBucket] {
			isKeyInBucket = true
		}
	}
	var zero T
	if !isKeyInBucket {
		log.Fatalln("The key doesn't exist in the hashmap.")
		return zero
	}
	index := HashToInt(key, hashmap.size)
	value := *hashmap.list[index]
	return value
}
