# Redis client for Golang
Supports:

- Redis 3 commands except QUIT, MONITOR, SLOWLOG and SYNC.
- [Pub/Sub](https://godoc.org/github.com/go-redis/redis#PubSub).
- [Transactions](https://godoc.org/github.com/go-redis/redis#Multi).
- [Pipeline](https://godoc.org/github.com/go-redis/redis#example-Client-Pipeline) and [TxPipeline](https://godoc.org/github.com/go-redis/redis#example-Client-TxPipeline).
- [Scripting](https://godoc.org/github.com/go-redis/redis#Script).
- [Timeouts](https://godoc.org/github.com/go-redis/redis#Options).
- [Redis Sentinel](https://godoc.org/github.com/go-redis/redis#NewFailoverClient).
- [Redis Cluster](https://godoc.org/github.com/go-redis/redis#NewClusterClient).
- [Ring](https://godoc.org/github.com/go-redis/redis#NewRing).
- [Instrumentation](https://godoc.org/github.com/go-redis/redis#ex-package--Instrumentation).
- [Cache friendly](https://github.com/go-redis/cache).
- [Rate limiting](https://github.com/go-redis/rate).
- [Distributed Locks](https://github.com/bsm/redis-lock).

API docs: https://godoc.org/github.com/go-redis/redis.
Examples: https://godoc.org/github.com/go-redis/redis#pkg-examples.

## Installation

Install:

```shell
go get -u github.com/go-redis/redis
```

Import:

```go
import "github.com/go-redis/redis"
```

## Quickstart

```go
func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func ExampleClient() {
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exists
}
```
