package core

import (
	"encoding/json"

	"github.com/Manan-Prakash-Singh/leetcode-go/utils"
)

var (
    tags = map[string]string{
        "dp":      "dynamic-programming",
		"sort":    "sorting",
		"bfs":     "breadth-first-search",
		"dfs":     "depth-first-search",
		"hashing": "hash-table",
		"hash":    "hash-table",
    }
    
)

func getTopic(tag string) (*TagResponse,error) {
    str := tags[tag]

    if str != "" {
        tag = str    
    }else {
        tag = utils.GetTitleSlug(tag)
    }

    url := "https://leetcode.com/problems/tag-data/question-tags/" + tag + "/"
    request ,err := utils.NewNormalRequest("GET",url,nil)

    if err != nil {
        return nil, err
    }

    response, err := utils.SendRequest(request)

    if err != nil {
        return nil, err
    }

    var res_err struct { Error string }

    err = json.Unmarshal(response,&res_err)

    if err != nil{
        return nil, err
    }

    if len(res_err.Error) != 0 {
        return nil, err
    }

    json_res := TagResponse{}

    err = json.Unmarshal(response, &json_res)

    if err != nil {
        return nil, err
    }

    return &json_res, err
}
