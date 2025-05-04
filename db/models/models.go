package models
import(
	"time"
	jwt "github.com/dgirjalva/jwt-go"
    "github.com/prasu-baran/randomstring"
)

type User struct {
	Username, PasswordHash , Role string
}
type Tokenchains struct {
	jwt.StandardClaims
	Role string 'json:"role"'
	Csrf string 'json:"csrf"'

}

const RefreshTokenValidTime = time.Hour * 72
const AuthTokenValidTime = time.Minute * 15

func GenerateRandomString()(string , error){
return randomStrings.GenearateRadomString(32)
}