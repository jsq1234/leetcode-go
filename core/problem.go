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

type ProblemOfTheDayResponse struct {
	Data struct {
		ActiveDailyCodingChallengeQuestion struct {
			Link     string   `json:"link"`
            Question struct {
                QuestionID         string `json:"questionId"`
                QuestionFrontEndID string `json:"questionFrontEndId"`
                IsPaidyOnly        bool   `json:"isPaidOnly"`
                Title              string `json:"title"`
                Difficulty         string `json:"difficulty"`
                TopicTags          []struct {
                    Name string `json:"name"`
                } `json:"topicTags"`
            } `json:"question"`
		} `json:"activeDailyCodingChallengeQuestion"`
	} `json:"data"`
}

func NewProblemOfTheDayResponse(data []byte) (*ProblemOfTheDayResponse, error){
    v := ProblemOfTheDayResponse{}
    err := json.Unmarshal(data,&v)
    return &v, err
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

    return NewProblemOfTheDayResponse(response)
}
