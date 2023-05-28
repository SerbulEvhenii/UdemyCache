Go udemy Cashe first my home work (in-memory cache).

See it in action:

## Example #1

```go
package main

import (
	"fmt"

	"github.com/SerbulEvhenii/udemyCache"
)

func main() {
	cache := udemyCache.New()
	fmt.Println(*cache)

	cache.Set("userId", 42)
	fmt.Println(*cache)

	userId := cache.Get("userId")
	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")
	fmt.Println(userId)
}

```