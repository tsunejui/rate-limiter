## Rate Limiter

Rate Limiter is a middleware on Echo Server for limiting network traffic based on IP address

## Example of usage

You can run the following command for using this middleware on your own project:

```
go get -u github.com/tsunejui/rate-limiter
```


### With default config (Implementation of `Token Bucket`)

```
e := echo.New()
e.Use(
    rMiddleware.Limiter(&rMiddleware.MiddlewareCofig{
        Mode:     "general-mode",
        TimeRate: 1 * time.Minute,
        Max:      1000,
    }),
)
```

### With redis config (Implementation of `Sliding Window`)

```
e := echo.New()
e.Use(
    rate.Limiter(&rate.MiddlewareCofig{
        Mode:     "redis-mode",
        TimeRate: 1 * time.Minute,
        Max:      1000,
        RedisOptions: &redis.Options{
            Addr:     "localhost:6379",
            Password: "test",
            DB:       1,
        },
    }),
)
```