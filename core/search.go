package core

import (
	"encoding/json"
	"github.com/Manan-Prakash-Singh/leetcode-go/utils"
)

const searchQuery = `
query problemsetQuestionList($categorySlug: String, $limit: Int, $skip: Int, $filters: QuestionListFilterInput) {
    problemsetQuestionList: questionList(
        categorySlug: $categorySlug
        limit: $limit
        skip: $skip
        filters: $filters
    ) {
        total: totalNum
        questions: data {
            difficulty
            title
            topicTags {
                name
            }
        }
    }
}` 

func newSearchQuery(searchKey string, limit int) ([]byte,error) {
    query := GraphqlQuery{
        Query: searchQuery,
        Variables: map[string]interface{}{
            "categorySlug" : "",
            "skip" : 0,
            "limit" : limit,
            "filters" : map[string]string{
                "searchKeywords" : searchKey,
            },
        },
        OperationName: "problemsetQuestionList",
    }
    return json.Marshal(&query)
}


func searchProblemRequest(searchKey string, count int) ([]Question, error){
   
    query, err := newSearchQuery(searchKey,count)
    if err != nil {
        return nil, err 
    }

    request, err := utils.NewNormalRequest("POST",GRAPHQL_URL,query)

    if err != nil {
        return nil, err 
    }

    response, err := utils.SendRequest(request)

    if err != nil {
        return nil, err 
    }

    data := SearchProblemResponse{}

    err = json.Unmarshal(response,&data)

    if err != nil {
        return nil, err 
    }

    return data.Data.ProblemsetQuestionList.Questions, nil 
}
