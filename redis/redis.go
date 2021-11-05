package redis

import (
	"fmt"
	"strings"
	"time"

	rds "github.com/gomodule/redigo/redis"
)

var db *rds.Pool

// Open connect to redis service, host: 127.0.0.1, port: 6379, auth: , serial: 1
func Open(host string, port int, auth string, serial int) (err error) {
	db = &rds.Pool{
		Dial: func() (con rds.Conn, err error) {
			con, err = rds.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
			if err != nil {
				return
			}
			if auth != "" {
				if _, err = con.Do("AUTH", auth); err != nil {
					return
				} // OK
			}
			if _, err = con.Do("SELECT", serial); err != nil {
				return
			} // OK
			return
		},
		MaxIdle:     256,
		MaxActive:   1024,
		IdleTimeout: time.Minute,
	}
	_, err = db.Dial()
	return
}

// GetDatabase get database connect pool
func GetDatabase() *rds.Pool {
	return db
}

// SetDatabase set database connect pool
func SetDatabase(database *rds.Pool) {
	db = database
}

// Do exec redis command 执行 redis 命令
func Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return db.Get().Do(strings.ToUpper(commandName), args...)
}

// parse reply

func Int(reply interface{}, err error) (int, error) {
	return rds.Int(reply, err)
}

func Int64(reply interface{}, err error) (int64, error) {
	return rds.Int64(reply, err)
}

func Uint64(reply interface{}, err error) (uint64, error) {
	return rds.Uint64(reply, err)
}

func Float64(reply interface{}, err error) (float64, error) {
	return rds.Float64(reply, err)
}

func String(reply interface{}, err error) (string, error) {
	return rds.String(reply, err)
}

func Bytes(reply interface{}, err error) ([]byte, error) {
	return rds.Bytes(reply, err)
}

func Bool(reply interface{}, err error) (bool, error) {
	return rds.Bool(reply, err)
}

func Values(reply interface{}, err error) ([]interface{}, error) {
	return rds.Values(reply, err)
}

func Float64s(reply interface{}, err error) ([]float64, error) {
	return rds.Float64s(reply, err)
}

func Strings(reply interface{}, err error) ([]string, error) {
	return rds.Strings(reply, err)
}

func ByteSlices(reply interface{}, err error) ([][]byte, error) {
	return rds.ByteSlices(reply, err)
}

func Int64s(reply interface{}, err error) ([]int64, error) {
	return rds.Int64s(reply, err)
}

func Ints(reply interface{}, err error) ([]int, error) {
	return rds.Ints(reply, err)
}

func StringMap(result interface{}, err error) (map[string]string, error) {
	return rds.StringMap(result, err)
}

func IntMap(result interface{}, err error) (map[string]int, error) {
	return rds.IntMap(result, err)
}

func Int64Map(result interface{}, err error) (map[string]int64, error) {
	return rds.Int64Map(result, err)
}

func Positions(result interface{}, err error) ([]*[2]float64, error) {
	return rds.Positions(result, err)
}

func Uint64s(reply interface{}, err error) ([]uint64, error) {
	return rds.Uint64s(reply, err)
}

func Uint64Map(result interface{}, err error) (map[string]uint64, error) {
	return rds.Uint64Map(result, err)
}

// redis client command doc => http://redisdoc.com/index.html

// 字符串

func SET(args ...interface{}) (reply interface{}, err error) {
	return Do("SET", args...)
}

func SETNX(args ...interface{}) (reply interface{}, err error) {
	return Do("SETNX", args...)
}

func SETEX(args ...interface{}) (reply interface{}, err error) {
	return Do("SETEX", args...)
}

func PSETEX(args ...interface{}) (reply interface{}, err error) {
	return Do("PSETEX", args...)
}

func GET(args ...interface{}) (reply interface{}, err error) {
	return Do("GET", args...)
}

func GETSET(args ...interface{}) (reply interface{}, err error) {
	return Do("GETSET", args...)
}

func STRLEN(args ...interface{}) (reply interface{}, err error) {
	return Do("STRLEN", args...)
}

func APPEND(args ...interface{}) (reply interface{}, err error) {
	return Do("APPEND", args...)
}

func SETRANGE(args ...interface{}) (reply interface{}, err error) {
	return Do("SETRANGE", args...)
}

func GETRANGE(args ...interface{}) (reply interface{}, err error) {
	return Do("GETRANGE", args...)
}

func INCR(args ...interface{}) (reply interface{}, err error) {
	return Do("INCR", args...)
}

func INCRBY(args ...interface{}) (reply interface{}, err error) {
	return Do("INCRBY", args...)
}

func INCRBYFLOAT(args ...interface{}) (reply interface{}, err error) {
	return Do("INCRBYFLOAT", args...)
}

func DECR(args ...interface{}) (reply interface{}, err error) {
	return Do("DECR", args...)
}

func DECRBY(args ...interface{}) (reply interface{}, err error) {
	return Do("DECRBY", args...)
}

func MSET(args ...interface{}) (reply interface{}, err error) {
	return Do("MSET", args...)
}

func MSETNX(args ...interface{}) (reply interface{}, err error) {
	return Do("MSETNX", args...)
}

func MGET(args ...interface{}) (reply interface{}, err error) {
	return Do("MGET", args...)
}

// 哈希表

func HSET(args ...interface{}) (reply interface{}, err error) {
	return Do("HSET", args...)
}

func HSETNX(args ...interface{}) (reply interface{}, err error) {
	return Do("HSETNX", args...)
}

func HGET(args ...interface{}) (reply interface{}, err error) {
	return Do("HGET", args...)
}

func HEXISTS(args ...interface{}) (reply interface{}, err error) {
	return Do("HEXISTS", args...)
}

func HDEL(args ...interface{}) (reply interface{}, err error) {
	return Do("HDEL", args...)
}

func HLEN(args ...interface{}) (reply interface{}, err error) {
	return Do("HLEN", args...)
}

func HSTRLEN(args ...interface{}) (reply interface{}, err error) {
	return Do("HSTRLEN", args...)
}

func HINCRBY(args ...interface{}) (reply interface{}, err error) {
	return Do("HINCRBY", args...)
}

func HINCRBYFLOAT(args ...interface{}) (reply interface{}, err error) {
	return Do("HINCRBYFLOAT", args...)
}

func HMSET(args ...interface{}) (reply interface{}, err error) {
	return Do("HMSET", args...)
}

func HMGET(args ...interface{}) (reply interface{}, err error) {
	return Do("HMGET", args...)
}

func HKEYS(args ...interface{}) (reply interface{}, err error) {
	return Do("HKEYS", args...)
}

func HVALS(args ...interface{}) (reply interface{}, err error) {
	return Do("HVALS", args...)
}

func HGETALL(args ...interface{}) (reply interface{}, err error) {
	return Do("HGETALL", args...)
}

func HSCAN(args ...interface{}) (reply interface{}, err error) {
	return Do("HSCAN", args...)
}

// 列表

func LPUSH(args ...interface{}) (reply interface{}, err error) {
	return Do("LPUSH", args...)
}

func LPUSHX(args ...interface{}) (reply interface{}, err error) {
	return Do("LPUSHX", args...)
}

func RPUSH(args ...interface{}) (reply interface{}, err error) {
	return Do("RPUSH", args...)
}

func RPUSHX(args ...interface{}) (reply interface{}, err error) {
	return Do("RPUSHX", args...)
}

func LPOP(args ...interface{}) (reply interface{}, err error) {
	return Do("LPOP", args...)
}

func RPOP(args ...interface{}) (reply interface{}, err error) {
	return Do("RPOP", args...)
}

func RPOPLPUSH(args ...interface{}) (reply interface{}, err error) {
	return Do("RPOPLPUSH", args...)
}

func LREM(args ...interface{}) (reply interface{}, err error) {
	return Do("LREM", args...)
}

func LLEN(args ...interface{}) (reply interface{}, err error) {
	return Do("LLEN", args...)
}

func LINDEX(args ...interface{}) (reply interface{}, err error) {
	return Do("LINDEX", args...)
}

func LINSERT(args ...interface{}) (reply interface{}, err error) {
	return Do("LINSERT", args...)
}

func LSET(args ...interface{}) (reply interface{}, err error) {
	return Do("LSET", args...)
}

func LRANGE(args ...interface{}) (reply interface{}, err error) {
	return Do("LRANGE", args...)
}

func LTRIM(args ...interface{}) (reply interface{}, err error) {
	return Do("LTRIM", args...)
}

func BLPOP(args ...interface{}) (reply interface{}, err error) {
	return Do("BLPOP", args...)
}

func BRPOP(args ...interface{}) (reply interface{}, err error) {
	return Do("BRPOP", args...)
}

func BRPOPLPUSH(args ...interface{}) (reply interface{}, err error) {
	return Do("BRPOPLPUSH", args...)
}

// 集合

func SADD(args ...interface{}) (reply interface{}, err error) {
	return Do("SADD", args...)
}

func SISMEMBER(args ...interface{}) (reply interface{}, err error) {
	return Do("SISMEMBER", args...)
}

func SPOP(args ...interface{}) (reply interface{}, err error) {
	return Do("SPOP", args...)
}

func SRANDMEMBER(args ...interface{}) (reply interface{}, err error) {
	return Do("SRANDMEMBER", args...)
}

func SREM(args ...interface{}) (reply interface{}, err error) {
	return Do("SREM", args...)
}

func SMOVE(args ...interface{}) (reply interface{}, err error) {
	return Do("SMOVE", args...)
}

func SCARD(args ...interface{}) (reply interface{}, err error) {
	return Do("SCARD", args...)
}

func SMEMBERS(args ...interface{}) (reply interface{}, err error) {
	return Do("SMEMBERS", args...)
}

func SSCAN(args ...interface{}) (reply interface{}, err error) {
	return Do("SSCAN", args...)
}

func SINTER(args ...interface{}) (reply interface{}, err error) {
	return Do("SINTER", args...)
}

func SINTERSTORE(args ...interface{}) (reply interface{}, err error) {
	return Do("SINTERSTORE", args...)
}

func SUNION(args ...interface{}) (reply interface{}, err error) {
	return Do("SUNION", args...)
}

func SUNIONSTORE(args ...interface{}) (reply interface{}, err error) {
	return Do("SUNIONSTORE", args...)
}

func SDIFF(args ...interface{}) (reply interface{}, err error) {
	return Do("SDIFF", args...)
}

func SDIFFSTORE(args ...interface{}) (reply interface{}, err error) {
	return Do("SDIFFSTORE", args...)
}

// 有序集合

func ZADD(args ...interface{}) (reply interface{}, err error) {
	return Do("ZADD", args...)
}

func ZSCORE(args ...interface{}) (reply interface{}, err error) {
	return Do("ZSCORE", args...)
}

func ZINCRBY(args ...interface{}) (reply interface{}, err error) {
	return Do("ZINCRBY", args...)
}

func ZCARD(args ...interface{}) (reply interface{}, err error) {
	return Do("ZCARD", args...)
}

func ZCOUNT(args ...interface{}) (reply interface{}, err error) {
	return Do("ZCOUNT", args...)
}

func ZRANGE(args ...interface{}) (reply interface{}, err error) {
	return Do("ZRANGE", args...)
}

func ZREVRANGE(args ...interface{}) (reply interface{}, err error) {
	return Do("ZREVRANGE", args...)
}

func ZRANGEBYSCORE(args ...interface{}) (reply interface{}, err error) {
	return Do("ZRANGEBYSCORE", args...)
}

func ZREVRANGEBYSCORE(args ...interface{}) (reply interface{}, err error) {
	return Do("ZREVRANGEBYSCORE", args...)
}

func ZRANK(args ...interface{}) (reply interface{}, err error) {
	return Do("ZRANK", args...)
}

func ZREVRANK(args ...interface{}) (reply interface{}, err error) {
	return Do("ZREVRANK", args...)
}

func ZREM(args ...interface{}) (reply interface{}, err error) {
	return Do("ZREM", args...)
}

func ZREMRANGEBYRANK(args ...interface{}) (reply interface{}, err error) {
	return Do("ZREMRANGEBYRANK", args...)
}

func ZREMRANGEBYSCORE(args ...interface{}) (reply interface{}, err error) {
	return Do("ZREMRANGEBYSCORE", args...)
}

func ZRANGEBYLEX(args ...interface{}) (reply interface{}, err error) {
	return Do("ZRANGEBYLEX", args...)
}

func ZLEXCOUNT(args ...interface{}) (reply interface{}, err error) {
	return Do("ZLEXCOUNT", args...)
}

func ZREMRANGEBYLEX(args ...interface{}) (reply interface{}, err error) {
	return Do("ZREMRANGEBYLEX", args...)
}

func ZSCAN(args ...interface{}) (reply interface{}, err error) {
	return Do("ZSCAN", args...)
}

func ZUNIONSTORE(args ...interface{}) (reply interface{}, err error) {
	return Do("ZUNIONSTORE", args...)
}

func ZINTERSTORE(args ...interface{}) (reply interface{}, err error) {
	return Do("ZINTERSTORE", args...)
}

// 数据库

func EXISTS(args ...interface{}) (reply interface{}, err error) {
	return Do("EXISTS", args...)
}

func TYPE(args ...interface{}) (reply interface{}, err error) {
	return Do("TYPE", args...)
}

func RENAME(args ...interface{}) (reply interface{}, err error) {
	return Do("RENAME", args...)
}

func RENAMENX(args ...interface{}) (reply interface{}, err error) {
	return Do("RENAMENX", args...)
}

func MOVE(args ...interface{}) (reply interface{}, err error) {
	return Do("MOVE", args...)
}

func DEL(args ...interface{}) (reply interface{}, err error) {
	return Do("DEL", args...)
}

func RANDOMKEY(args ...interface{}) (reply interface{}, err error) {
	return Do("RANDOMKEY", args...)
}

func DBSIZE(args ...interface{}) (reply interface{}, err error) {
	return Do("DBSIZE", args...)
}

func KEYS(args ...interface{}) (reply interface{}, err error) {
	return Do("KEYS", args...)
}

func SCAN(args ...interface{}) (reply interface{}, err error) {
	return Do("SCAN", args...)
}

func SORT(args ...interface{}) (reply interface{}, err error) {
	return Do("SORT", args...)
}

func FLUSHDB(args ...interface{}) (reply interface{}, err error) {
	return Do("FLUSHDB", args...)
}

func FLUSHALL(args ...interface{}) (reply interface{}, err error) {
	return Do("FLUSHALL", args...)
}

func SELECT(args ...interface{}) (reply interface{}, err error) {
	return Do("SELECT", args...)
}

func SWAPDB(args ...interface{}) (reply interface{}, err error) {
	return Do("SWAPDB", args...)
}

// 自动过期

func EXPIRE(args ...interface{}) (reply interface{}, err error) {
	return Do("EXPIRE", args...)
}

func EXPIREAT(args ...interface{}) (reply interface{}, err error) {
	return Do("EXPIREAT", args...)
}

func TTL(args ...interface{}) (reply interface{}, err error) {
	return Do("TTL", args...)
}

func PERSIST(args ...interface{}) (reply interface{}, err error) {
	return Do("PERSIST", args...)
}

func PEXPIRE(args ...interface{}) (reply interface{}, err error) {
	return Do("PEXPIRE", args...)
}

func PEXPIREAT(args ...interface{}) (reply interface{}, err error) {
	return Do("PEXPIREAT", args...)
}

func PTTL(args ...interface{}) (reply interface{}, err error) {
	return Do("PTTL", args...)
}

// 事务

func MULTI(args ...interface{}) (reply interface{}, err error) {
	return Do("MULTI", args...)
}

func EXEC(args ...interface{}) (reply interface{}, err error) {
	return Do("EXEC", args...)
}

func DISCARD(args ...interface{}) (reply interface{}, err error) {
	return Do("DISCARD", args...)
}

func WATCH(args ...interface{}) (reply interface{}, err error) {
	return Do("WATCH", args...)
}

func UNWATCH(args ...interface{}) (reply interface{}, err error) {
	return Do("UNWATCH", args...)
}

// 持久化

func SAVE(args ...interface{}) (reply interface{}, err error) {
	return Do("SAVE", args...)
}

func BGSAVE(args ...interface{}) (reply interface{}, err error) {
	return Do("BGSAVE", args...)
}

func BGREWRITEAOF(args ...interface{}) (reply interface{}, err error) {
	return Do("BGREWRITEAOF", args...)
}

func LASTSAVE(args ...interface{}) (reply interface{}, err error) {
	return Do("LASTSAVE", args...)
}

// 发布与订阅

func PUBLISH(args ...interface{}) (reply interface{}, err error) {
	return Do("PUBLISH", args...)
}

func SUBSCRIBE(args ...interface{}) (reply interface{}, err error) {
	return Do("SUBSCRIBE", args...)
}

func PSUBSCRIBE(args ...interface{}) (reply interface{}, err error) {
	return Do("PSUBSCRIBE", args...)
}

func UNSUBSCRIBE(args ...interface{}) (reply interface{}, err error) {
	return Do("UNSUBSCRIBE", args...)
}

func PUNSUBSCRIBE(args ...interface{}) (reply interface{}, err error) {
	return Do("PUNSUBSCRIBE", args...)
}

func PUBSUB(args ...interface{}) (reply interface{}, err error) {
	return Do("PUBSUB", args...)
}

// 复制

func SLAVEOF(args ...interface{}) (reply interface{}, err error) {
	return Do("SLAVEOF", args...)
}

func ROLE(args ...interface{}) (reply interface{}, err error) {
	return Do("ROLE", args...)
}

// 客户端与服务器

func AUTH(args ...interface{}) (reply interface{}, err error) {
	return Do("AUTH", args...)
}

func QUIT(args ...interface{}) (reply interface{}, err error) {
	return Do("QUIT", args...)
}

func INFO(args ...interface{}) (reply interface{}, err error) {
	return Do("INFO", args...)
}

func SHUTDOWN(args ...interface{}) (reply interface{}, err error) {
	return Do("SHUTDOWN", args...)
}

func TIME(args ...interface{}) (reply interface{}, err error) {
	return Do("TIME", args...)
}

func CLIENT_GETNAME(args ...interface{}) (reply interface{}, err error) {
	return Do("CLIENT_GETNAME", args...)
}

func CLIENT_KILL(args ...interface{}) (reply interface{}, err error) {
	return Do("CLIENT_KILL", args...)
}

func CLIENT_LIST(args ...interface{}) (reply interface{}, err error) {
	return Do("CLIENT_LIST", args...)
}

func CLIENT_SETNAME(args ...interface{}) (reply interface{}, err error) {
	return Do("CLIENT_SETNAME", args...)
}

// 配置

func CONFIG_SET(args ...interface{}) (reply interface{}, err error) {
	return Do("CONFIG_SET", args...)
}

func CONFIG_GET(args ...interface{}) (reply interface{}, err error) {
	return Do("CONFIG_GET", args...)
}

func CONFIG_RESETSTAT(args ...interface{}) (reply interface{}, err error) {
	return Do("CONFIG_RESETSTAT", args...)
}

func CONFIG_REWRITE(args ...interface{}) (reply interface{}, err error) {
	return Do("CONFIG_REWRITE", args...)
}
