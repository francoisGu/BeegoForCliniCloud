package test

import (
	"beegoProject/APIConnect"
	_ "beegoProject/routers"
	// "fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestMain is a sample to run an endpoint test
func TestMain(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestMain", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

// Test with correct password usename
func TestLoginWithCorrectInfo(t *testing.T) {
	succcessLogin, _, content := APIConnect.LoginToAPI(
		"test@clinicloud.com", "pass123", "X-CHALLENGE-APP", "Y2xpbmljbG91ZGNoYWxsZW5nZXNlY3JldGtleQ==")

	Convey("Subject: Test Login\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(succcessLogin, ShouldEqual, true)
		})
		Convey("The uuid Should exist", func() {
			So(len(content["uuid"].(string)), ShouldBeGreaterThan, 0)
		})
		Convey("The access_token Should exist", func() {
			So(len(content["access_token"].(string)), ShouldBeGreaterThan, 0)
		})
	})
}

// Test with wrong password usename
func TestLoginWithWrongInfo(t *testing.T) {
	succcessLogin, _, content := APIConnect.LoginToAPI(
		"test@clinicloud.com", "wrongpass", "X-CHALLENGE-APP", "Y2xpbmljbG91ZGNoYWxsZW5nZXNlY3JldGtleQ==")

	Convey("Subject: Test Login\n", t, func() {
		Convey("Status Code Should not Be 200", func() {
			So(succcessLogin, ShouldEqual, false)
		})
		Convey("Type Should be error", func() {
			So(content["Type"], ShouldEqual, "error")
		})
	})
}

// Test to get user with correct info
func TestGetUserWithCorrectInfo(t *testing.T) {
	_, _, loginContent := APIConnect.LoginToAPI(
		"test@clinicloud.com", "pass123", "X-CHALLENGE-APP", "Y2xpbmljbG91ZGNoYWxsZW5nZXNlY3JldGtleQ==")
	success, _, userContent := APIConnect.GetUser(loginContent["uuid"].(string),
		loginContent["access_token"].(string))

	Convey("Subject: Test Get user\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(success, ShouldEqual, true)
		})
		Convey("Type Should be user_info", func() {
			So(userContent["Type"], ShouldEqual, "user_info")
		})
	})
}

//Test get session with correct uerinfo
func TestGetSessionWithCorrectInfo(t *testing.T) {
	_, _, loginContent := APIConnect.LoginToAPI(
		"test@clinicloud.com", "pass123", "X-CHALLENGE-APP", "Y2xpbmljbG91ZGNoYWxsZW5nZXNlY3JldGtleQ==")

	success, _, sessionContent := APIConnect.GetSession(loginContent)

	Convey("Subject: Test Get session\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(success, ShouldEqual, true)
		})
		Convey("Type Should be sessions", func() {
			So(sessionContent.Type, ShouldEqual, "sessions")
		})
	})
}

// Should be more test with different parameters, like wrong userid
// wrong clientid, client secret
// Need test with more functions
