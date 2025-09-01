package cache

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

func BenchmarkMemoryCache_Set(b *testing.B) {
	cache := NewMemoryCache()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Set(strconv.Itoa(i), i)
	}
}

func BenchmarkMemoryCache_Get(b *testing.B) {
	cache := NewMemoryCache()
	numKeys := 10000
	
	for i := 0; i < numKeys; i++ {
		cache.Set(strconv.Itoa(i), i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(strconv.Itoa(i % numKeys))
	}
}

func BenchmarkMemoryCache_Del(b *testing.B) {
	cache := NewMemoryCache()
	
	for i := 0; i < b.N; i++ {
		cache.Set(strconv.Itoa(i), i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Del(strconv.Itoa(i))
	}
}

func BenchmarkMemoryCache_Mixed(b *testing.B) {
	cache := NewMemoryCache()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		switch i % 3 {
		case 0:
			cache.Set(key, i)
		case 1:
			cache.Get(key)
		case 2:
			cache.Del(key)
		}
	}
}

func BenchmarkMemoryCache_ConcurrentSet(b *testing.B) {
	cache := NewMemoryCache()
	
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			cache.Set(strconv.Itoa(i), i)
			i++
		}
	})
}

func BenchmarkMemoryCache_ConcurrentGet(b *testing.B) {
	cache := NewMemoryCache()
	numKeys := 10000
	
	for i := 0; i < numKeys; i++ {
		cache.Set(strconv.Itoa(i), i)
	}
	
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cache.Get(strconv.Itoa(rand.Intn(numKeys)))
		}
	})
}

func BenchmarkMemoryCache_ConcurrentMixed(b *testing.B) {
	cache := NewMemoryCache()
	
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := strconv.Itoa(i)
			switch i % 3 {
			case 0:
				cache.Set(key, i)
			case 1:
				cache.Get(key)
			case 2:
				cache.Del(key)
			}
			i++
		}
	})
}

func BenchmarkMemoryCache_LargeValues(b *testing.B) {
	cache := NewMemoryCache()
	largeData := make([]byte, 1024*10)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Set(strconv.Itoa(i), largeData)
	}
}

func BenchmarkMemoryCache_HighContention(b *testing.B) {
	cache := NewMemoryCache()
	numKeys := 10
	
	for i := 0; i < numKeys; i++ {
		cache.Set(strconv.Itoa(i), i)
	}
	
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := strconv.Itoa(rand.Intn(numKeys))
			switch rand.Intn(3) {
			case 0:
				cache.Set(key, rand.Int())
			case 1:
				cache.Get(key)
			case 2:
				cache.Del(key)
			}
		}
	})
}

func BenchmarkMemoryCache_Sequential_vs_Concurrent(b *testing.B) {
	b.Run("Sequential", func(b *testing.B) {
		cache := NewMemoryCache()
		for i := 0; i < b.N; i++ {
			key := fmt.Sprintf("key_%d", i)
			cache.Set(key, i)
			cache.Get(key)
			cache.Del(key)
		}
	})
	
	b.Run("Concurrent", func(b *testing.B) {
		cache := NewMemoryCache()
		var wg sync.WaitGroup
		numWorkers := 10
		
		b.ResetTimer()
		for w := 0; w < numWorkers; w++ {
			wg.Add(1)
			go func(workerID int) {
				defer wg.Done()
				for i := 0; i < b.N/numWorkers; i++ {
					key := fmt.Sprintf("key_%d_%d", workerID, i)
					cache.Set(key, i)
					cache.Get(key)
					cache.Del(key)
				}
			}(w)
		}
		wg.Wait()
	})
}