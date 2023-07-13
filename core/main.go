package core

import (
	"fmt"
	"os"
    "strings"
    "github.com/Manan-Prakash-Singh/leetcode-go/utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/pterm/pterm"
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
    spinnerInfo, _ := pterm.DefaultSpinner.Start("Fetching question of the day...")
    data, err := GetProblemOfTheDay()
    if err != nil {
        spinnerInfo.Fail(err)
        os.Exit(1)
    }

    title := data.Data.ActiveDailyCodingChallengeQuestion.Question.Title
    difficulty := data.Data.ActiveDailyCodingChallengeQuestion.Question.Difficulty
    
    str := fmt.Sprintf("%v\t%v\n",title,utils.Color(difficulty))
    spinnerInfo.Success(str)
    
    reply := utils.UserInput("Download problem? [y,n]: ")

    if reply == "y" {
        lang := utils.UserInput("Language? [cpp,java,python,etc.]: ")
        DownloadProblem(title,lang)
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

func SearchProblem(searchKey string, count int) {

    spinnerInfo, _ := pterm.DefaultSpinner.Start("Searching") 
    ques, err := searchProblemRequest(searchKey,count)
    if err != nil || len(ques) == 0 {
        if len(ques) == 0 {
            spinnerInfo.Fail("No question matched the search term")
            os.Exit(1)
        }
        spinnerInfo.Fail(err)
        os.Exit(1)
    }
    spinnerInfo.Success("") 

    var opts []string

    maxLen := 0
    for _, val := range ques {
        if maxLen < len(val.Title) {
            maxLen = len(val.Title)
        }
    }

    for _,val := range ques {
        opts = append(opts,fmt.Sprintf("%-*s\t[%s]",maxLen,val.Title,utils.Color(val.Difficulty)))
    }

    p := pterm.DefaultInteractiveSelect
    p = *p.WithDefaultText("Select the problem")
    maxHeight := count
    if count >= 20 {
        maxHeight = 20
    }
    p = *p.WithMaxHeight(maxHeight)

    selectedOptions, _ := p.WithOptions(opts).Show()

    problem := selectedOptions[:strings.Index(selectedOptions,"[")]
    problem = strings.Trim(problem, "\t \n")

    lang := utils.UserInput("Select a language: ")

    DownloadProblem(problem,lang)
}

func DownloadProblem(problem, lang string) {
    spinnerInfo, _ := pterm.DefaultSpinner.Start("Downloading")
    err := _downloadProblem(problem,lang)
    if err != nil {
        spinnerInfo.Fail(err)
        os.Exit(1)
    }
    spinnerInfo.InfoPrinter = &pterm.PrefixPrinter{
		MessageStyle: &pterm.Style{pterm.FgGreen},
		Prefix: pterm.Prefix{
			Style: &pterm.Style{pterm.FgBlack, pterm.BgGreen},
			Text:  " DOWNLOADED ",
		},
	}
    spinnerInfo.Info("")
}

func GetTopics(topic string){
    str := "Fetcing " + topic + " problems..."
    spinnerInfo, _ := pterm.DefaultSpinner.Start(str)
    data, err := getTopic(topic)
    if err != nil {
        spinnerInfo.Fail(err)
        os.Exit(1)
    }
    ques := data.Questions
    quesLen := len(ques)
    spinnerInfo.Success(fmt.Sprintf("%d questions found.",quesLen))
    
    var opts []string

    maxLen := 0
    for _, val := range ques {
        if maxLen < len(val.Title) {
            maxLen = len(val.Title)
        }
    }

    for _,val := range ques {
        opts = append(opts,fmt.Sprintf("%-*s\t[%s]",maxLen,val.Title,utils.Color(val.Difficulty)))
    }

    p := pterm.DefaultInteractiveSelect
    p = *p.WithDefaultText("Select the problem")
    p = *p.WithMaxHeight(25)

    selectedOptions, _ := p.WithOptions(opts).Show()

    problem := selectedOptions[:strings.Index(selectedOptions,"[")]
    problem = strings.Trim(problem, "\t \n")
    
    lang := utils.UserInput("Select a language: ")

    DownloadProblem(problem,lang)
}

func RunCode(fileName string) {
    spinnerInfo, _ := pterm.DefaultSpinner.Start("Executing...")
    err := runTestCases(fileName, false)
    if err != nil {
        spinnerInfo.Fail(err)
        os.Exit(1)
    }
    spinnerInfo.Stop()
}

func SubmitCode(fileName string) {
    spinnerInfo, _ := pterm.DefaultSpinner.Start("Submitting...")
    err := submitCode(fileName) 
    if err != nil {
        spinnerInfo.Fail(err)
        os.Exit(1)
    }
    spinnerInfo.Stop()

	_, problemName, _, _ := utils.ParseFileName(fileName)        

    SuggestQuestions(problemName)
}

func SuggestQuestions(problem string) {

    data, err := suggestProblems(problem)

    if err != nil {
        pterm.Info.WithPrefix(pterm.Prefix{ 
            Style: &pterm.Style{pterm.FgBlack, pterm.BgRed,pterm.Bold},
            Text: " ERROR ",
        }).Println(err)
    }

    var opts []string

    ques := data.Data.Question.NextChallenges

    maxLen := 0
    dict := make(map[string]string)

    for _, val := range ques {
        dict[val.Title] = val.TitleSlug
        if maxLen < len(val.Title) {
            maxLen = len(val.Title)
        }
    }

    for _,val := range ques {
        opts = append(opts,fmt.Sprintf("%-*s\t[%s]",maxLen,val.Title,utils.Color(val.Difficulty)))
    }

    p := pterm.DefaultInteractiveSelect
    p = *p.WithDefaultText("Select the problem")
    p = *p.WithMaxHeight(25)

    selectedOptions, _ := p.WithOptions(opts).Show()
    newProblem := selectedOptions[:strings.Index(selectedOptions,"[")]
    newProblem = strings.Trim(newProblem,"\n\t ")

    lang := utils.UserInput("Select a language: ")

    DownloadProblem(newProblem,lang)

}
