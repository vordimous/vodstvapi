package sess

import (
	"fmt"
	"net/url"
	"os"

	"github.com/vodstv/core"

	"github.com/gin-gonic/contrib/sessions"
	_ "github.com/jinzhu/gorm/dialects/postgres" //import postgres
)

const (
	//sessHost ...
	sessHost = "localdocker"
	//sessUser ...
	sessUser = "esvodsapi"
	//sessPassword ...
	sessPassword = "esvodsapi"
	//sessName ...
	sessName = "esvods"
)

//Init ...
func Init() sessions.RedisStore {

	redisInfo := os.Getenv("REDIS_URL")

	fmt.Println("REDIS_URL: " + redisInfo)

	if redisInfo == "" {
		s, err := sessions.NewRedisStore(10, "tcp", sessHost+":6379", "", []byte("secret"))
		core.CheckErr(err, "Cannot connect to RedisStore")
		return s
	}

	//connect to reddis REDIS_URL
	u, err := url.Parse(redisInfo)
	core.CheckErr(err, "Cannot parse redis url")
	p, _ := u.User.Password()
	s, err := sessions.NewRedisStore(10, "tcp", u.Host, p, []byte("secret"))
	core.CheckErr(err, "Cannot connect to RedisStore")

	return s

}
