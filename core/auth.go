package core

import (
	"encoding/json"
	"fmt"
	"github.com/Manan-Prakash-Singh/leetcode-go/utils"
)

type userData struct{
    UserName string `json:"user_name"`
    NumSolved int `json:"num_solved"`
    EasyCount int `json:"ac_easy"`
    MediumCount int `json:"ac_medium"`
    HardCount int `json:"ac_hard"`
}

func Authenticate() (*userData,error) {
    request, err := utils.NewAuthRequest("GET", "https://leetcode.com/api/problems/all", nil)

    if err != nil {
        return nil, err
    }
    response, err := utils.SendRequest(request)

    if err != nil {
        return nil, err
    }

    data := userData{}

    err = json.Unmarshal(response,&data)

    if err != nil {
        return nil, err
    }

    if data.UserName == "" {
        return nil, fmt.Errorf("Couldn't authenticate the user. Renew your cookie")
    }
    return &data, nil 
}
