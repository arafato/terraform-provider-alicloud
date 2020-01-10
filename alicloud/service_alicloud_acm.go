package alicloud

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var acmEndpointURL = "http://acm.aliyun.com:8080/diamond-server/"

func acmPublishConfigurations(tenant string, dataID string, group string, content string, ak string, sk string, token string) (status int, err error) {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}
	params := url.Values{}
	params.Add("dataId", dataID)
	params.Add("group", group)
	params.Add("content", content)
	params.Add("tenant", tenant)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%sbasestone.do?method=syncUpdateAll%s", acmEndpointURL, params.Encode()), nil)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	str := tenant + group + timestamp
	_acmAddHeaders(req, ak, timestamp, _acmSignString(str, sk), token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	return resp.StatusCode, nil
}

func acmGetConfiguration(tenant string, dataID string, group string, ak string, sk string, token string) (string, int, error) {

	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%sconfig.co", acmEndpointURL), nil)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	str := tenant + group + timestamp
	_acmAddHeaders(req, ak, timestamp, _acmSignString(str, sk), token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", 0, err
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}

	return string(content), resp.StatusCode, nil
}

func acmDeleteConfiguration(tenant string, dataID string, group string, ak string, sk string, token string) (status int, err error) {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%sdatum.do", acmEndpointURL), nil)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	str := tenant + group + timestamp
	_acmAddHeaders(req, ak, timestamp, _acmSignString(str, sk), token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	return resp.StatusCode, nil
}

func _acmAddHeaders(req *http.Request, ak string, timestamp string, spasSignature string, token string) {

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=GBK")
	req.Header.Add("Spas-AccessKey", ak)
	req.Header.Add("timeStamp", timestamp)
	req.Header.Add("Spas-Signature", spasSignature)
	if token != "" {
		req.Header.Add("Spas-SecurityToken", token)
	}
}

func _acmSignString(str string, sk string) string {

	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(sk))
	io.WriteString(h, str)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signedStr
}
