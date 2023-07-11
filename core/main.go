package core

import (
	"fmt"
	"os"
	"github.com/Manan-Prakash-Singh/leetcode-go/utils"
    "github.com/pterm/pterm"
	"github.com/jedib0t/go-pretty/v6/table"
)

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

func ProblemOfTheDay() {
    data, err := GetProblemOfTheDay()

    if err != nil {
        fmt.Fprintln(os.Stderr,err)
        os.Exit(1)
    }

    title := data.Data.ActiveDailyCodingChallengeQuestion.Question.Title
    difficulty := data.Data.ActiveDailyCodingChallengeQuestion.Question.Difficulty

    fmt.Printf("%v\t[%v]\n",title,utils.Color(difficulty))
    
    reply := utils.UserInput("Download problem? [y,n]: ")

    if reply == "y" {
        lang := utils.UserInput("Language? [cpp,java,python,etc.]: ")
        fmt.Println("Downloading problem of the day....")
        if err := DownloadProblem(title,lang); err != nil {
              fmt.Println(err)
        } 
    }

}

func AuthenticateUser() {
    spinnerInfo, _ := pterm.DefaultSpinner.Start("Authenticating...") 
    data, err := Authenticate()
    if err != nil {
        spinnerInfo.Fail(err)
        os.Exit(1)
    }
    spinnerInfo.Success("")

    t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)
    t.AppendRow([]interface{}{ "Username", data.UserName }) 
    t.AppendSeparator()
    t.AppendRows([]table.Row{
        { "Easy", data.EasyCount },
        { "Medium", data.MediumCount },
        { "Hard", data.HardCount },
    })
    t.SetStyle(table.StyleBold)
    t.Render()
}
