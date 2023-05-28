package udemyCache

type Cache struct {
	userId map[string]interface{}
}

func New() *Cache {
	return &Cache{make(map[string]interface{})}
}

func (c *Cache) Set(key string, value interface{}) {
	c.userId[key] = value
}

func (c *Cache) Get(key string) interface{} {
	return c.userId[key]
}

func (c *Cache) Delete(key string) {
	delete(c.userId, key)
}
