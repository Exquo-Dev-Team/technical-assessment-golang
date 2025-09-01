package cache

import (
	"testing"
)

func TestNewMemoryCache(t *testing.T) {
	cache := NewMemoryCache()
	if cache == nil {
		t.Fatal("NewMemoryCache() returned nil")
	}
}

func TestMemoryCache_Set(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		value   any
		wantErr bool
	}{
		{
			name:    "set string value",
			key:     "key1",
			value:   "value1",
			wantErr: false,
		},
		{
			name:    "set integer value",
			key:     "key2",
			value:   42,
			wantErr: false,
		},
		{
			name:    "set struct value",
			key:     "key3",
			value:   struct{ Name string }{Name: "test"},
			wantErr: false,
		},
		{
			name:    "set nil value",
			key:     "key4",
			value:   nil,
			wantErr: true,
		},
		{
			name:    "set empty key",
			key:     "",
			value:   "value",
			wantErr: true,
		},
		{
			name:    "overwrite existing key",
			key:     "key1",
			value:   "new_value",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := NewMemoryCache()
			err := cache.Set(tt.key, tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemoryCache_Get(t *testing.T) {
	cache := NewMemoryCache()

	testKey := "test_key"
	testValue := "test_value"

	cache.Set(testKey, testValue)

	tests := []struct {
		name      string
		key       string
		wantValue any
		wantFound bool
	}{
		{
			name:      "get existing key",
			key:       testKey,
			wantValue: testValue,
			wantFound: true,
		},
		{
			name:      "get non-existing key",
			key:       "non_existing",
			wantValue: nil,
			wantFound: false,
		},
		{
			name:      "get empty key",
			key:       "",
			wantValue: nil,
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := cache.Get(tt.key)
			if found != tt.wantFound {
				t.Errorf("Get() found = %v, wantFound %v", found, tt.wantFound)
			}
			if found && value != tt.wantValue {
				t.Errorf("Get() value = %v, wantValue %v", value, tt.wantValue)
			}
		})
	}
}

func TestMemoryCache_Del(t *testing.T) {
	tests := []struct {
		name        string
		setupKey    string
		setupValue  any
		deleteKey   string
		wantDeleted bool
	}{
		{
			name:        "delete existing key",
			setupKey:    "key1",
			setupValue:  "value1",
			deleteKey:   "key1",
			wantDeleted: true,
		},
		{
			name:        "delete non-existing key",
			setupKey:    "key1",
			setupValue:  "value1",
			deleteKey:   "key2",
			wantDeleted: false,
		},
		{
			name:        "delete empty key",
			setupKey:    "key1",
			setupValue:  "value1",
			deleteKey:   "",
			wantDeleted: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := NewMemoryCache()

			if tt.setupKey != "" {
				cache.Set(tt.setupKey, tt.setupValue)
			}

			deleted := cache.Del(tt.deleteKey)
			if deleted != tt.wantDeleted {
				t.Errorf("Del() = %v, want %v", deleted, tt.wantDeleted)
			}

			if deleted {
				if _, found := cache.Get(tt.deleteKey); found {
					t.Errorf("Key %s still exists after deletion", tt.deleteKey)
				}
			}
		})
	}
}
