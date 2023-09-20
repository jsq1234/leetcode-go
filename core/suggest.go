package core

import (
	"encoding/json"
	"github.com/Manan-Prakash-Singh/leetcode-go/utils"
)

const suggestQuery = `
query nextChallengePairs($questionSlug: String!) {
    question(titleSlug: $questionSlug) {
        nextChallenges {
            difficulty
            title
            titleSlug
            questionFrontendId
        }
    }
}
`

func newSuggestionQuery(problem string) ([]byte, error) {
	query := GraphqlQuery{
		Query: suggestQuery,
		Variables: map[string]interface{}{
			"questionSlug": problem,
		},
		OperationName: "nextChallengePairs",
	}

	return json.Marshal(query)
}

func suggestProblems(problem string) (*SuggestResponse, error) {

	query, err := newSuggestionQuery(problem)

	if err != nil {
		return nil, err
	}

	request, err := utils.NewAuthRequest("POST", GRAPHQL_URL, query)

	if err != nil {
		return nil, err
	}

	response, err := utils.SendRequest(request)

	if err != nil {
		return nil, err
	}

	var data SuggestResponse

	err = json.Unmarshal(response, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
