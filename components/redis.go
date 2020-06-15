package components
import (
	"github.com/astaxie/beego"
	"time"
	"github.com/garyburd/redigo/redis"
)

var(
	pool *redis.Pool
	redisHost = beego.AppConfig.String("redisinfo")
	redisPass = beego.AppConfig.String("redispwd")
)

//newRedisPool:创建redis连接池
func newRedisPool()*redis.Pool{
	return &redis.Pool{
		MaxIdle:500,
		MaxActive:300,
		IdleTimeout:300*time.Second,
		Dial: func() (redis.Conn,  error) {
			c,err := redis.Dial("tcp",redisHost,redis.DialPassword(redisPass))
			if err != nil {
				beego.Error("redis初始化失败：",err)
				return nil,err
			}
			//2、访问认证
			//if _, err =c.Do("AUTH",redisPass);err!=nil{
			//	c.Close()
			//	return nil,err
			//}
			beego.Info("redis初始化成功")
			return c,nil
		},
		//定时检查redis是否出状况
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t)<time.Minute{
				return nil
			}
			_,err := conn.Do("PING")
			return err
		},
	}
}

//初始化redis连接池
func InitRedis(){
	pool = newRedisPool()
}

//对外暴露连接池
func RedisPool()*redis.Pool{
	return pool
}

func Set(key,value string,expireTime int) (err error){
	rc := RedisPool().Get()
	//注意close()
	defer rc.Close()
	beego.Info(redis.PoolStats{})
	//默认不选库操作的是0
	//rc.Do("select",2)
	// 写入值10S后过期
	_, err = rc.Do("SET", key, value, "EX", expireTime)
	if err != nil {
		beego.Error(err)
	}
	return err
}
func Get(key string) (err error){
	rc := RedisPool().Get()
	//注意close()
	defer rc.Close()
	_, err = rc.Do("GET", key)
	if err != nil {
		beego.Error(err)
	}
	return err
}