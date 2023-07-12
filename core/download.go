package core

import (
	"encoding/json"
	"fmt"
    "io"
    "os"
	"github.com/Manan-Prakash-Singh/leetcode-go/utils"
)

const queryDownload =`
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

func newDownloadProblemQuery(query , problem string) *GraphqlQuery {
    return &GraphqlQuery{
        Query: queryDownload,
        Variables: map[string]interface{}{
            "titleSlug" : problem,
        },
        OperationName: "questionContent",
    }
}
func _downloadProblem(problem, lang string) error {

    problem = utils.GetTitleSlug(problem)  

    queryDownload := newDownloadProblemQuery(query, problem)

    query, err := json.Marshal(queryDownload)

    if err != nil {
        return fmt.Errorf("Json Marshalling error %v", err) 
    }
    
    request, err := utils.NewNormalRequest("POST",GRAPHQL_URL,query)

    if err != nil {
        return fmt.Errorf("Request creation err: %v",err)
    }

    response, err := utils.SendRequest(request)
    
    if err != nil {
        return fmt.Errorf("Response error: %v",err)
    }

    data, err := newDownloadProblemResponse(response)

    if err != nil {
        return fmt.Errorf("Response parsing err: %v", err)
    }

    if data.Data.Question.IsPaidOnly {
        return fmt.Errorf("The selected problem is premium problem.")    
    }

    if len(data.Errors) > 0 {
        return fmt.Errorf("Couldn't find the requested problem. Try doing a search instead.")
    }

    problemHTML := data.Data.Question.Content

    if err := utils.RenderHTML(problemHTML); err != nil {
        return fmt.Errorf("HTML rendering error, did you install lynx? : %v", err)
    }
 
    return createCodeFile(problem,lang,data)
}

func createCodeFile(problem, lang string, cnt *DownloadProblemResponse) error {
    
	questionID := cnt.Data.Question.QuestionID

	fileName := problem + "_" + questionID + "." + lang

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
