package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// thi pkg will manage the env vars whether loading them or accessing them

const (
	DbUser     EnvKey = "dbuser"
	DbPassword EnvKey = "dbpassword"
	DbHost     EnvKey = "dbhost"
	DbName     EnvKey = "dbname"
	DbPort     EnvKey = "dbport"

	StripeKey EnvKey = "stripekey"

	WebhookSecret EnvKey = "webhooksecret"

	QrcodePath EnvKey = "qrcodepath"

	ServerPort EnvKey = "serverport"

	StripeSuccessUrl EnvKey = "stripesuccessurl"

	StripeCancelUrl EnvKey = "stripecancelurl"

	HomePageUrl EnvKey = "homepageurl"
)

type EnvKey string

func (key EnvKey) GetValue() string {

	return os.Getenv(string(key))
}

func LoadEnv() {

	fmt.Println("we are in the dev env")

	godotenv.Load()

}
