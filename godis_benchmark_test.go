package godis

import "testing"

// Benchmark for SET method
func BenchmarkSET(b *testing.B) {
	db := setUp()
	b.ResetTimer()

	// run the SET method b.N times
	for n := 0; n < b.N; n++ {
		db.SET("key", "value")
	}
}

// Benchmark for GET method
func BenchmarkGET(b *testing.B) {
	db := setUp()
	db.SET("key", "value")
	b.ResetTimer()

	// run the GET method b.N times
	for n := 0; n < b.N; n++ {
		_ = db.GET("key")
	}
}

// Benchmark for EXISTS method
func BenchmarkEXISTS(b *testing.B) {
	db := setUp()
	db.SET("key", "value")
	b.ResetTimer()

	// run the EXISTS method b.N times
	for n := 0; n < b.N; n++ {
		_ = db.EXISTS("key")
	}
}

// Benchmark for DEL method
func BenchmarkDEL(b *testing.B) {
	db := setUp()
	b.ResetTimer()

	// SET a key and DEL method b.N times
	for n := 0; n < b.N; n++ {
		// Not an accurate measure as we can't DEL a key without setting it
		db.SET("key", "value")
		_ = db.DEL("key")
	}
}
