package main

import "fmt"

type EthProof map[string][]byte

// Has returns if the key exists in db.
// Error is always nil.
func (p EthProof) Has(key []byte) (bool, error) {
	_, ok := p[string(key)]
	return ok, nil
}

// Get retrieves the value for the given key.
// If key is not found, return invalid key error.
func (p EthProof) Get(key []byte) ([]byte, error) {
	if entry, ok := p[string(key)]; ok {
		return CopyBytes(entry), nil
	}
	return nil, fmt.Errorf("invalid key: %v", key)
}

// Put inserts the given key-value pair.
// Existing key/value is overwritten.
func (p EthProof) Put(key, value []byte) error {
	p[string(key)] = CopyBytes(value)
	return nil
}

// Delete removes the key from the key-value store.
// If the key doesn't exist, it's a no-op without error.
func (p EthProof) Delete(key []byte) error {
	delete(p, string(key))
	return nil
}

// CopyBytes returns an exact copy of the provided bytes.
func CopyBytes(b []byte) []byte {
	if b == nil {
		return nil
	}
	copiedBytes := make([]byte, len(b))
	copy(copiedBytes, b)
	return copiedBytes
}
