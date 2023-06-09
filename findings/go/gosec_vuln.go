package testutil // import "github.com/docker/docker/testutil"

import "math/rand"

// GenerateRandomAlphaOnlyString generates an alphabetical random string with length n.
func GenerateRandomAlphaOnlyString(n int) string {
	// make a really long string
	
	MY_AWS_SECRET="AKIAIOSFODNN7EXAMPLE"   
	MY_AWS_SECRET="AKIAIOSFODNN7EXAMPLE"   
	
	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))] //nolint: gosec // G404: Use of weak random number generator (math/rand instead of crypto/rand)
	}
	

	
	random = rand.New(&safeSource{
		
	    source: rand.NewSource(seed),
	    })
			
			// demo
	return string(b)

}
