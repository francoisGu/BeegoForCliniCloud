package APIConnect

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	POST         = "POST"
	GET          = "GET"
	host         = "https://stg.clinicloud.com"
	basePath     = "/api/v2"
	clientId     = "X-CHALLENGE-APP"
	clientSecret = "Y2xpbmljbG91ZGNoYWxsZW5nZXNlY3JldGtleQ=="
	email        = "test@clinicloud.com"
	password     = "pass123"
)

type (
	Recording struct {
		ID        string
		SessionID string
		UserID    string
		Type      int
		Data      struct {
			ambientTemp float64
			surfaceTemp float64
		}
		Created   float64
		Updated   float64
		State     string
		TrackData struct {
			SerialMeasuringDevice   string
			FirmwareMeasuringDevice string
		}
	}

	ContentElem struct {
		ID          string
		UserId      string
		CaregiverID string
		Created     float64
		Latitude    float64
		Longitud    float64
		Location    string
		Note        string
		DeviceId    string
		DeviceModel string
		AppVersion  string
		UpdateOrder int
		Recordings  []interface{}
		AppType     string
		Sequence    int
	}

	Session struct {
		Type    string
		Content []ContentElem
	}
)

/*Public functions*/
func LoginToAPI(email string, password string,
	clientId string, clientSecret string) (succcess bool,
	jsonBody string, contentMsg map[string]interface{}) {

	url, headerMsg, bodyMsg := prepareLoginInfo(email, password, clientId, clientSecret)
	statusJson, bodyByte := httpProcess(POST, url, headerMsg, bodyMsg)

	status := strings.Fields(statusJson)
	if status[0] == "200" {
		succcess = true
	} else {
		succcess = false
	}

	fmt.Println(status)

	contentMsg = ProcessResp(bodyByte)
	return succcess, string(bodyByte), contentMsg
	// return processResp(statusJson, bodyByte)
}

//, typeMsg string, bodyMsg map[string]string
func GetUser(userId string, token string) (succcess bool,
	jsonBody string, contentMap map[string]interface{}) {

	url, headerMsg := prepareGetUserInfo(userId, token)
	statusJson, bodyByte := httpProcess(GET, url, headerMsg, nil)

	status := strings.Fields(statusJson)
	if status[0] == "200" {
		contentMap = ProcessResp(bodyByte)
		return true, string(bodyByte), contentMap
	} else {
		msg := map[string]interface{}{"type": "error"}
		return false, string(bodyByte), msg

	}

}

func GetSession(contentMsg map[string]interface{}) (succcess bool, jsonBody string, s Session) {
	url, headerMsg, bodyMsg := prepareGetSessionInfo(contentMsg)
	statusJson, bodyByte := httpProcess(POST, url, headerMsg, bodyMsg)

	status := strings.Fields(statusJson)
	if status[0] == "200" {
		succcess = true
	} else {
		succcess = false
	}
	fmt.Println(status)

	rtype, rCL := processSessionBodyMsg(bodyByte)
	session := Session{
		Type:    rtype,
		Content: rCL}
	fmt.Println(len(rCL))

	return succcess, string(bodyByte), session
}

/*Pricate functions*/
func getAuthorizationMsg(email string, password string) string {
	authMessage := email + ":" + password
	authEncoded := "Basic " + base64.StdEncoding.EncodeToString([]byte(authMessage))
	return authEncoded
}

func httpProcess(methodType string, url string,
	headerMsg map[string]string, bodyMsg []byte) (respStatus string, respBody []byte) {
	client := &http.Client{}
	req, err := http.NewRequest(methodType, url, bytes.NewBuffer(bodyMsg))

	for key, value := range headerMsg {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	// defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return resp.Status, body
}

func ProcessResp(bodyMsg []byte) (respContent map[string]interface{}) {

	var dat map[string]interface{}

	if err := json.Unmarshal(bodyMsg, &dat); err != nil {
		panic(err)
	}

	// respType = dat["Type"].(string)

	respContent = dat["Content"].(map[string]interface{})
	respContent["Type"] = dat["Type"].(string)

	return respContent
}

func processSessionBodyMsg(bodyMsg []byte) (rtype string, contentList []ContentElem) {
	var dat map[string]interface{}
	if err := json.Unmarshal(bodyMsg, &dat); err != nil {
		panic(err)
	}

	resType := dat["Type"].(string)
	contentJsonList := dat["Content"].([]interface{})
	cl := make([]ContentElem, len(contentJsonList))
	fmt.Println("Session content list length:", len(contentJsonList))

	for index, cElem := range contentJsonList {

		elemByte, _ := json.Marshal(cElem)

		structElem := ContentElem{}
		if err1 := json.Unmarshal(elemByte, &structElem); err1 != nil {
			panic(err1)
		}
		// structElem := cElem.(ContentElem)
		cl[index] = structElem
	}

	return resType, cl
}

/*login functions*/
func prepareLoginInfo(email string, password string,
	clientId string, clientSecret string) (url string,
	headerMsg map[string]string, bodyMsg []byte) {
	// url
	url = host + basePath + "/login"

	// body message
	bodyMessage := map[string]string{
		"client_id":     clientId,
		"client_secret": clientSecret}
	bodyJson, _ := json.Marshal(bodyMessage)

	// header message
	authEncoded := getAuthorizationMsg(email, password)
	headerMsg = map[string]string{
		"Content-Type":  "application/json",
		"PushToken":     "",
		"DeviceType":    "",
		"DeviceToken":   "",
		"Authorization": authEncoded}

	return url, headerMsg, bodyJson
}

/*Get user functions*/
func prepareGetUserInfo(userID string,
	access_token string) (url string, headerMsg map[string]string) {
	//url
	url = host + basePath + "/user/" + userID

	// header message
	headerMsg = map[string]string{
		"Content-Type":  "application/json",
		"PushToken":     "",
		"DeviceToken":   "",
		"userID":        userID,
		"Authorization": "Bearer " + access_token}

	return url, headerMsg

}

/*Get session*/
func prepareGetSessionInfo(contentMsg map[string]interface{}) (url string, headerMsg map[string]string, bodyMsg []byte) {
	//url
	url = host + basePath + "/sessions/get"

	//header
	access_token := contentMsg["access_token"].(string)
	headerMsg = map[string]string{
		"Content-Type":  "application/json",
		"PushToken":     "",
		"DeviceToken":   "",
		"Authorization": "Bearer " + access_token}

	//body
	userID := contentMsg["uuid"].(string)
	bodyMsg = []byte(`[{"UserID":"` + userID + `","UpdateOrder":1}]`)

	return url, headerMsg, bodyMsg
}
