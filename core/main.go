package core

import (
	"fmt"
	"os"
    "strings"
    "github.com/gookit/color"
    "text/tabwriter"
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
    if err != nil {
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
    p = *p.WithMaxHeight(count)

    selectedOptions, _ := p.WithOptions(opts).Show()

    problem := selectedOptions[:strings.Index(selectedOptions,"[")]
    problem = strings.Trim(problem, "\t \n")

    DownloadProblem(problem,"cpp")
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
    spinnerInfo.Success("")
    
    questions := data.Questions

    w := tabwriter.NewWriter(os.Stdout, 4, 4, 4, ' ', 0)

	for i, val := range questions {
		var difficulty string
		switch val.Difficulty {
		case "Hard":
			difficulty = color.Red.Sprintf(val.Difficulty)
		case "Medium":
			difficulty = color.Yellow.Sprintf(val.Difficulty)
		case "Easy":
			difficulty = color.Green.Sprintf(val.Difficulty)

		}
		fmt.Fprintf(w, "%v.\t%v\t[%v]\n", i+1, val.Title, difficulty)
	}

	w.Flush()

}
