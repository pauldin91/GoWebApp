package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/pauldin91/GoWebApp/internal/cfg"
	"github.com/pauldin91/GoWebApp/internal/models"
)

var session *scs.SessionManager
var testApp cfg.AppConfig

func TestMain(m *testing.M) {
	gob.Register(models.Reservation{})
	testApp.InProduction = false
	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (m *myWriter) Header() http.Header {
	return http.Header{}
}
func (m *myWriter) WriteHeader(i int) {

}
func (m *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}
