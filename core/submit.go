package core

import (
	"encoding/json"
	"fmt"
	"github.com/Manan-Prakash-Singh/leetcode-go/utils"
	"github.com/gookit/color"
	"github.com/pterm/pterm"
	"io"
	"net/http"
	"os"
)

func getSubmissionId(fileName string, submit bool) (*RunTestCaseResponse, error) {

	questionID, problemName, lang, err := utils.ParseFileName(fileName)

	if err != nil {
		return nil, err
	}

	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	fileContent, err := io.ReadAll(file)

	fileContentStr := string(fileContent)

	if err != nil {
		return nil, err
	}

	testCases, err := utils.ParseTestCases(fileContentStr)

	if err != nil {
		return nil, err
	}

	jsonReq := map[string]string{
		"data_input":  testCases,
		"lang":        lang,
		"question_id": questionID,
		"typed_code":  fileContentStr,
	}

	requestBody, err := json.Marshal(jsonReq)

	if err != nil {
		return nil, err
	}

	runUrl := "https://leetcode.com/problems/" + problemName + "/interpret_solution/"
	submitUrl := "https://leetcode.com/problems/" + problemName + "/submit/"

	var request *http.Request

	if submit {
		request, err = utils.NewAuthRequest("POST", submitUrl, requestBody)
	} else {
		request, err = utils.NewAuthRequest("POST", runUrl, requestBody)
	}
	if err != nil {
		return nil, err
	}

	body, err := utils.SendRequest(request)

	if err != nil {
		return nil, err
	}

	var response RunTestCaseResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func execute(id string) (*SubmissionResponse, error) {

	submissionResult := &SubmissionResponse{}

	url := "https://leetcode.com/submissions/detail/" + id + "/check/"

	request, err := utils.NewAuthRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	response, err := utils.SendRequest(request)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &submissionResult)

	if err != nil {
		return nil, err
	}

	return submissionResult, nil
}

func runTestCases(fileName string, submit bool) (bool, error) {

	testCaseResponse, err := getSubmissionId(fileName, submit)

	if err != nil {
		return false, err
	}

	var id string

	if submit {
		id = fmt.Sprint(testCaseResponse.SubmissionId)
	} else {
		id = testCaseResponse.InterpretId
	}

	Result := &SubmissionResponse{}

	for {
		Result, err = execute(id)
		if err != nil {
			return false, err
		}
		if Result.State == "SUCCESS" {
			break
		}
	}

	res := OutputResult(testCaseResponse, Result)
	return res, nil
}

func OutputResult(testCaseResponse *RunTestCaseResponse, Result *SubmissionResponse) bool {

	statusMsg := Result.StatusMsg

	passed := false

	switch statusMsg {

	case "Compile Error":

		pterm.Info.WithPrefix(pterm.Prefix{
			Style: &pterm.Style{pterm.FgBlack, pterm.BgRed, pterm.Bold},
			Text:  " COMPILE ERROR ",
		}).Println()

		color.Redln(Result.FullCompileError)

		passed = false

	case "Accepted":

		if testCaseResponse.InterpretId == Result.SubmissionID {
			for i, testCase := range utils.TestCaseList {

				if Result.CodeAnswer[i] != Result.ExpectedCodeAnswer[i] {

					pterm.Info.WithPrefix(pterm.Prefix{
						Style: &pterm.Style{pterm.FgBlack, pterm.BgRed, pterm.Bold},
						Text:  " WRONG ANSWER ",
					}).Println()

				} else {
					pterm.Info.WithPrefix(pterm.Prefix{
						Style: &pterm.Style{pterm.FgBlack, pterm.BgGreen, pterm.Bold},
						Text:  " CORRECT ANSWER ",
					}).Println()
				}

				fmt.Println("Input")
				fmt.Println(testCase)
				fmt.Println("Output")
				fmt.Println(Result.CodeAnswer[i])
				fmt.Println("Expected")
				fmt.Println(Result.ExpectedCodeAnswer[i])

				color.Yellowln("----------------------------------------------------")
			}
		}

		if fmt.Sprint(testCaseResponse.SubmissionId) == Result.SubmissionID {

			pterm.Info.WithPrefix(pterm.Prefix{
				Style: &pterm.Style{pterm.FgBlack, pterm.BgGreen, pterm.Bold},
				Text:  " ACCEPTED ",
			}).Println()

			fmt.Printf("Test cases passed: %d/%d\n", Result.TotalCorrect, Result.TotalTestcases)
			fmt.Printf("Runtime : %v [Beats : %0.2f%%]\n", Result.StatusRuntime, Result.RuntimePercentile)
			fmt.Printf("Memory : %v [Beats : %0.2f%%]\n", Result.StatusMemory, Result.MemoryPercentile)

			passed = true
		}

	case "Wrong Answer":

		pterm.Info.WithPrefix(pterm.Prefix{
			Style: &pterm.Style{pterm.FgBlack, pterm.BgRed, pterm.Bold},
			Text:  " WRONG ANSWER ",
		}).Println()

		fmt.Printf("Test cases passed: %d/%d\n", Result.TotalCorrect, Result.TotalTestcases)
		fmt.Println("Last test case executed: ")
		fmt.Println(Result.LastTestcase)
		fmt.Println("Expected Output:")
		fmt.Println(Result.ExpectedOutput)
		fmt.Println("Your Output:")
		fmt.Println(Result.CodeOutput.(string))

	case "Time Limit Exceeded":

		if testCaseResponse.InterpretId == Result.SubmissionID {
			pterm.Info.WithPrefix(pterm.Prefix{
				Style: &pterm.Style{pterm.FgBlack, pterm.BgRed, pterm.Bold},
				Text:  " TIME LIMIT EXCEEDED ",
			}).Println()
		}
		if fmt.Sprint(testCaseResponse.SubmissionId) == Result.SubmissionID {

			pterm.Info.WithPrefix(pterm.Prefix{
				Style: &pterm.Style{pterm.FgBlack, pterm.BgRed, pterm.Bold},
				Text:  " TIME LIMIT EXCEEDED ",
			}).Println()

			fmt.Printf("Test cases passed: %d/%d\n", Result.TotalCorrect, Result.TotalTestcases)
			fmt.Println("Last test case executed: ")
			fmt.Println(Result.LastTestcase)
			fmt.Println("Expected Output:")
			fmt.Println(Result.ExpectedOutput)
			fmt.Println("Your Output:")
			fmt.Println(Result.CodeOutput.(string))

		}
	}

	return passed
}
func submitCode(fileName string) (bool, error) {
	res, err := runTestCases(fileName, true)
	if err != nil {
		return false, err
	}
	return res, nil
}
