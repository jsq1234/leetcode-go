package utils

import (
	"bytes"
	"io"
	"net/http"
)

func NewNormalRequest(method string, url string, body []byte) (*http.Request, error) {
    if method == "POST" {
        request, err := http.NewRequest(method,url,bytes.NewBuffer(body))
        if err != nil {
            return nil, err
        }
        request.Header.Add("content-type","application/json")
        return request, err
    }
    return http.NewRequest(method,url,nil)
}

func NewAuthRequest(method string, url string, body []byte) (*http.Request, error) {
    key, err := getCookiesFromBrowser() 

    if err != nil {
        return nil, err
    }

    request, err := NewNormalRequest(method,url,body)

    if err != nil {
        return nil, err
    }
    
    request.AddCookie(&http.Cookie{
        Name: "csrftoken",
        Value: key.Csrftoken,
    })
    request.AddCookie(&http.Cookie{
        Name: "LEETCODE_SESSION",
        Value: key.LeetcodeSession,
    })

    request.Header.Set("Referer", "https://leetcode.com/")
	request.Header.Set("User-Agent", "Mozilla/6.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	request.Header.Set("X-csrftoken", key.Csrftoken)

	return request, nil
}

func SendRequest(request *http.Request) ([]byte, error) {
    client := &http.Client{}

    response, err := client.Do(request)
    
    if err != nil {
        return nil, err
    }

    defer response.Body.Close()

    body, err := io.ReadAll(response.Body)

    if err != nil {
        return nil, err
    }

    return body, nil
}
