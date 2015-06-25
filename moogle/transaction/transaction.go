package transaction

type Transaction interface {
	// Adds key with the given value to the transaction.
	// Returns true if successful,
	// false if the key already existed.
	Add(key, value string) (ok bool)

	// Set key to the given value, if the key already exists.
	// Returns true if successful.
	// If the key did not exist, it is not added and false is returned.
	Set(key, value string) (ok bool)

	// Delete key from the transaction.
	// Returns true and the value that was there, if the key existed.
	// Returns false if it did not exist.
	Del(key string) (value string, ok bool)

	// Get key from the transaction.
	// Returns true and the value that was there, if the key exists.
	// Returns false if it does not exist.
	Get(key string) (value string, ok bool)
}
