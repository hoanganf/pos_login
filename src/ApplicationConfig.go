package src

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hoanganf/pos_domain/entity"
	"github.com/hoanganf/pos_domain/service"
	"github.com/hoanganf/pos_login/src/application"
	"github.com/hoanganf/pos_login/src/infrastructure/persistence"
	"gopkg.in/gorp.v1"
	"os"
	"strings"
)

type Bean struct {
	LoginService *application.LoginService
	UserService  *application.UserService
	DbMap        *gorp.DbMap
}

func (bean *Bean) DestroyBean() {
	bean.DbMap.Db.Close()
}

func InitBean() (*Bean, error) {
	user := getEnvWithDefault("DB_USER", "root")
	password := getEnvWithDefault("DB_PASSWORD", "")
	//	host := getEnvWithDefault("DB_HOST", "127.0.0.1")
	//	port := getEnvWithDefault("DB_PORT", "3306")
	dbName := getEnvWithDefault("DB_NAME", "anit_pos_server_new")
	//	dsn := fmt.Sprintf("%s:%s@unix(%s:%s)/%s?parseTime=true", user, password, host, port,dbName)
	dsn := fmt.Sprintf("%s:%s@unix(/Applications/XAMPP/xamppfiles/var/mysql/mysql.sock)/%s?parseTime=true", user, password, dbName)
	fmt.Printf("dns: %s", dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	userRepository := persistence.NewUserRepository(dbMap)
	userService := application.NewUserService(service.NewUserService(
		getEnvWithDefault("POS_JWT_KEY", "b3BlbnNzaC1rZXktdjEAAAAACmFl"),
		userRepository), entity.NewUserFactory())

	loginService := application.NewLoginService(
		getEnvWithDefault("POS_HOME_PAGE", "http://pos.server.vn/pos/pos-portal"),
		strings.Split(getEnvWithDefault("POS_DOMAINS", "localhost"), ","),
		getEnvWithDefault("POS_LOGIN_TOKEN", "pos_access_token"),
		userService.UserService)
	return &Bean{LoginService: loginService, UserService: userService, DbMap: dbMap}, nil
}

func getEnvWithDefault(name, def string) string {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env
	}
	return def
}

func getEnvRequired(name string) (string, error) {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env, nil
	}
	return "", errors.New("not found env: " + name)
}
