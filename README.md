# Cache Service Implementation Task

## Your Task

Implement a simple in-memory cache service similar to Redis with basic operations:
- `Set(key, value)` - Store a key-value pair
- `Get(key)` - Retrieve a value by key
- `Del(key)` - Delete a key

## Requirements

1. **Thread-safe**: Must handle concurrent access
2. **Error handling**: Proper error handling for edge cases
3. **No nil values**: Reject nil values with appropriate error

## Running Your Solution

```bash
go run ./cmd/memorycache

# Run tests
go test -v -race

# Run with race detector
go test -race
