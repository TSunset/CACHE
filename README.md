# In-Memory Cache (Go)

Простой in-memory кеш для языка Go с поддержкой базовых операций: `Set`, `Get`, `Delete`.

---

## Пример использования

```go
package main

import (
	"fmt"
	"your-project/cache"
)

func main() {
	c := cache.New()

	c.Set("userId", 42)

	userId := c.Get("userId")
	fmt.Println(userId) // 42

	c.Delete("userId")

	userId = c.Get("userId")
	fmt.Println(userId) // <nil>
}
