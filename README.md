# tip-redis


Tip Redis is a library to connect to redis. Made to avoid repeating code (i.e. creating a new struct for every project that needs a redis connection).
Please see 'cmd/main/main.go' for usage.

##Note!
Avoid using close unless you actually want to close connection. According to the underlying go-redis library (github.com/go-redis/redis), connections are intended to be long-lived and to be used concurrently.
