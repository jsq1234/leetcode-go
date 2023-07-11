package core

import (
	"encoding/json"
	"github.com/Manan-Prakash-Singh/leetcode-go/utils"
)

const query = `
query questionOfToday {
    activeDailyCodingChallengeQuestion {
        link
        question {
            title
            difficulty
        }
    }
}`

func NewProblemOfTheDayQuery() ([]byte,error){
    query := GraphqlQuery{
        Query: query,
        Variables: nil,
        OperationName: "questionOfToday",
    }
    return json.Marshal(&query)
}

func GetProblemOfTheDay() (*ProblemOfTheDayResponse,error) {

    body, err := NewProblemOfTheDayQuery()

    if err != nil {
       return nil, err
    }

    request, err := utils.NewNormalRequest("POST",GRAPHQL_URL,body)

    if err != nil {
        return nil, err
    }

    response, err := utils.SendRequest(request)

    if err != nil {
        return nil, err
    }

    return newProblemOfTheDayResponse(response)
}
