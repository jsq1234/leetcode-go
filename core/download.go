package core

import (
	"encoding/json"
	"fmt"
	"github.com/Manan-Prakash-Singh/leetcode-go/utils"
	"io"
	"os"
)

const queryDownload = `
query questionContent($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    content
    mysqlSchemas
    questionId
    questionFrontendId
    isPaidOnly
    exampleTestcaseList
    codeSnippets {
      lang
      langSlug
      code
    }
  }
}`

func newDownloadProblemQuery(query, problem string) *GraphqlQuery {
	return &GraphqlQuery{
		Query: queryDownload,
		Variables: map[string]interface{}{
			"titleSlug": problem,
		},
		OperationName: "questionContent",
	}
}

func downloadProblem(problem string) (*DownloadProblemResponse, error) {

	problem = utils.GetTitleSlug(problem)

	queryDownload := newDownloadProblemQuery(query, problem)

	query, err := json.Marshal(queryDownload)

	if err != nil {
		return nil, fmt.Errorf("Json Marshalling error %v", err)
	}

	request, err := utils.NewNormalRequest("POST", GRAPHQL_URL, query)

	if err != nil {
		return nil, fmt.Errorf("Request creation err: %v", err)
	}

	response, err := utils.SendRequest(request)

	if err != nil {
		return nil, fmt.Errorf("Response error: %v", err)
	}

	data, err := newDownloadProblemResponse(response)

	if err != nil {
		return nil, fmt.Errorf("Response parsing err: %v", err)
	}

	if data.Data.Question.IsPaidOnly {
		return nil, fmt.Errorf("The selected problem is premium problem.")
	}

	if len(data.Errors) > 0 {
		return nil, fmt.Errorf("Couldn't find the requested problem. Try doing a search instead.")
	}

	problemHTML := data.Data.Question.Content
	id := data.Data.Question.QuestionID

	if err := utils.RenderHTML(problem, id, problemHTML); err != nil {
		return nil, fmt.Errorf("HTML rendering error: %v", err)
	}

	return data, nil
}

func createCodeFile(problem, lang string, cnt *DownloadProblemResponse) error {

	questionID := cnt.Data.Question.QuestionID

	ext := lang

	if l, ok := languageExtension[lang]; ok {
		ext = l
	}

	fileName := problem + "_" + questionID + "." + ext

	Editorfile, err := os.Create(fileName)

	if err != nil {
		return err
	}

	var content string

	io.WriteString(Editorfile, "/*\n\n")

	for _, val := range cnt.Data.Question.ExampleTestcaseList {
		io.WriteString(Editorfile, val+";\n")
		io.WriteString(Editorfile, "\n")
	}

	io.WriteString(Editorfile, "*/\n")

	for _, val := range cnt.Data.Question.CodeSnippets {
		if val.LangSlug == lang {
			content = val.Code
		}
	}

	io.WriteString(Editorfile, content)

	return nil
}
