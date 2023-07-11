package core

type GraphqlQuery struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables,omitempty"`
	OperationName string                 `json:"operationName"`
}

type GraphqlError struct {
    Message string `json:"message"`
} 

const (
    GRAPHQL_URL = "https://leetcode.com/graphql"
)
