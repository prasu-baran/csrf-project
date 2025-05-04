package middleware 
 
import(){
    "github.com/justinas/alice"
	"log"
	"net/http"
	"time"
	"strings"
	"github.com/prasu-baran/csrf-project/server/middleware/myJwt"
	"github.com/prasu-baran/csrf-project/db"
} 

func NewHandler(){
    return alice.New(recoverHandler,authHandler).ThenFunc(logicHandler)
}

func recoverHandler(next http.Handler)http.Handler{
 func(w http.ResponseWriter, r *http.Request){
	defer func(){
     if err:=recover(); err!=nil{
		log.Panic("Recovered! Panic:%v",err)
		http.Error(w,http.StatusText(500),500)
	 }
	}()
	next.ServerHTTP(w,r)
 }
}

func authHandler(next http.Handler) http.Handler{
fn:=func(w http.ResponseWriter,r *httpRequest){
	switch r.URL.Path{
	case "/restricted","/logout","/deleteUser":
	default:
	}
}
}

func logicHandler(w http.ResponseWriter, r *http.Request){
  switch r.URL.Path{
	case "/restricted"
  case "/login":
	switch r.Method{
	case "GET":
	case "POST":
	default:	
	}
case "/register":
	switch r.Method{
		case "GET":
		case "POST":
		default:
	}
case "/logout":
case "/deleteUser":
default:
  }	
}

func nullifyTokenCookies(w *http.ResponseWriter,r *http.Request){
 authCookie:=http.Cookie{
	Name:"auth",
	Value:"",
	Expires:time.Now().Add(-1000*time.Hour),
	HttpOnly:true,
 }
 http.SetCookie(*w,&authCookie)

 refreshCookie:=http.Cookie{
	Name:"RefreshToken",
	Value:"",
	Expires:time.Now().Add(-1000*time.Hour),
	HttpOnly:true,
 }
 http.SetCookie(*w,&refreshCookie)
}

func setAuthAndRefreshCookies(){

}

func grabCsrfFromReq(r *http.Request) string{

}