package test

import (
	"os"
	"testing"

	"github.com/deep-project/kingdee"
	"github.com/deep-project/kingdee/adapters"
	"github.com/deep-project/kingdee/client"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func TestKingdeeLoginBySign(t *testing.T) {
	logKDSVCSessionId(t, getLoginBySignAdapter())
}

func TestKingdeeLoginByAppSecret(t *testing.T) {
	logKDSVCSessionId(t, getLoginByAppSecretAdapter())
}

func TestKingdeeLoginByValidateUser(t *testing.T) {
	logKDSVCSessionId(t, getLoginByValidateUserAdapter())
}

func getLoginBySignAdapter() *adapters.LoginBySign {
	return &adapters.LoginBySign{
		AccountID:  os.Getenv("ACCOUNT_ID"),
		Username:   os.Getenv("USER_NAME"),
		AppID:      os.Getenv("APP_ID"),
		AppSecret:  os.Getenv("APP_SECRET"),
		LanguageID: os.Getenv("LANGUAGE_ID"),
	}
}

func getLoginByAppSecretAdapter() *adapters.LoginByAppSecret {
	return &adapters.LoginByAppSecret{
		AccountID:  os.Getenv("ACCOUNT_ID"),
		Username:   os.Getenv("USER_NAME"),
		AppID:      os.Getenv("APP_ID"),
		AppSecret:  os.Getenv("APP_SECRET"),
		LanguageID: os.Getenv("LANGUAGE_ID"),
	}
}

func getLoginByValidateUserAdapter() *adapters.LoginByValidateUser {
	return &adapters.LoginByValidateUser{
		AccountID:  os.Getenv("ACCOUNT_ID"),
		Username:   os.Getenv("USER_NAME"),
		Password:   os.Getenv("USER_PASSWORD"),
		LanguageID: os.Getenv("LANGUAGE_ID"),
	}
}

func logKDSVCSessionId(t *testing.T, login client.Login) {
	cli, err := kingdee.New(client.NewAPI(os.Getenv("BASE_URL")), login, &client.Options{})
	if err != nil {
		t.Error(err)
	}
	t.Log(cli.Handler.API.KDSVCSessionId)
}
