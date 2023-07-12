package core

import "encoding/json"

type DownloadProblemResponse struct {
    Data struct {
        Question struct {
			QuestionID          string   `json:"questionId"`
			ExampleTestcaseList []string `json:"exampleTestcaseList"`
            Content string `json:"content"`
            IsPaidOnly  bool `json:"isPaidOnly"`
			CodeSnippets        []struct {
				Lang     string `json:"lang"`
				LangSlug string `json:"langSlug"`
				Code     string `json:"code"`
			} `json:"codeSnippets"`
        }
    }
    Errors []GraphqlError
}

type Question struct {
	QuestionID         string `json:"questionId"`
	QuestionFrontEndID string `json:"questionFrontEndId"`
	IsPaidyOnly        bool   `json:"isPaidOnly"`
	Title              string `json:"title"`
	Difficulty         string `json:"difficulty"`
	TopicTags          []struct {
		Name string `json:"name"`
	} `json:"topicTags"`
}

type TagResponse struct {
	Questions []struct {
		QuestionID         string `json:"questionId"`
		QuestionFrontEndID string `json:"questionFrontEndId"`
		IsPaidyOnly        bool   `json:"isPaidOnly"`
		Title              string `json:"title"`
		Difficulty         string `json:"difficulty"`
		TopicTags          []struct {
			Name string `json:"name"`
		} `json:"topicTags"`
	} `json:"questions"`
}

type ProblemOfTheDayResponse struct {
	Data struct {
		ActiveDailyCodingChallengeQuestion struct {
			Link     string   `json:"link"`
			Question Question `json:"question"`
		} `json:"activeDailyCodingChallengeQuestion"`
	} `json:"data"`
}

type SearchProblemResponse struct {
	Data struct {
		ProblemsetQuestionList struct {
			Total     int        `json:"total"`
			Questions []Question `json:"questions"`
		} `json:"problemsetQuestionList"`
	} `json:"data"`
}

type UserResponse struct {
	Data struct {
		MatchedUser struct {
			ContestBadge string
			Username     string
			GithubUrl    string
			LinkedinUrl  string
			Profile      Profile
		}
	}
}

type Profile struct {
	Ranking   int
	RealName  string
	AboutMe   string
	SkillTags []string
}

type GenericUserData struct {
	Data struct {
		UserStatus struct {
			UserId      int
			IsSignedInt bool
			IsVerified  bool
			Username    string
		}
	}
}

type UserProblemsSolved struct {
	Data struct {
		MatchedUser struct {
			SubmitStatsGlobal struct {
				AcSubmissionNum []struct {
					Difficulty string
					Count      int
				}
			}
		}
	}
}

type RunTestCaseResponse struct {
	InterpretId  string `json:"interpret_id"`
	SubmissionId int    `json:"submission_id"`
	TestCases    string `json:"test_case"`
}

type SubmissionResponse struct {
	StatusCode             int         `json:"status_code"`
	Lang                   string      `json:"lang"`
	RunSuccess             bool        `json:"run_success"`
	CompileError           string      `json:"compile_error"`
	FullCompileError       string      `json:"full_compile_error"`
	StatusRuntime          string      `json:"status_runtime"`
	Memory                 int         `json:"memory"`
	CodeAnswer             []string    `json:"code_answer"`
	CodeOutput             interface{} `json:"code_output"`
	StdOutputList          []string    `json:"std_output_list"`
	ElapsedTime            int         `json:"elapsed_time"`
	ExpectedStatusCode     int         `json:"expected_status_code"`
	ExpectedLang           string      `json:"expected_lang"`
	ExpectedRunSuccess     bool        `json:"expected_run_success"`
	ExpectedStatusRuntime  string      `json:"expected_status_runtime"`
	ExpectedMemory         int         `json:"expected_memory"`
	ExpectedCodeAnswer     []string    `json:"expected_code_answer"`
	ExpectedCodeOutput     []string    `json:"expected_code_output"`
	ExpectedStdOutputList  []string    `json:"expected_std_output_list"`
	ExpectedElapsedTime    int         `json:"expected_elapsed_time"`
	ExpectedTaskFinishTime int64       `json:"expected_task_finish_time"`
	ExpectedTaskName       string      `json:"expected_task_name"`
	TaskFinishTime         int64       `json:"task_finish_time"`
	TaskName               string      `json:"task_name"`
	CorrectAnswer          bool        `json:"correct_answer"`
	TotalCorrect           int         `json:"total_correct"`
	TotalTestcases         int         `json:"total_testcases"`
	RuntimePercentile      float32     `json:"runtime_percentile"`
	StatusMemory           string      `json:"status_memory"`
	MemoryPercentile       float32     `json:"memory_percentile"`
	PrettyLang             string      `json:"pretty_lang"`
	SubmissionID           string      `json:"submission_id"`
	Input                  string      `json:"input"`
	StatusMsg              string      `json:"status_msg"`
	State                  string      `json:"state"`
	QuestionID             string      `json:"question_id"`
	CompareResult          string      `json:"compare_result"`
	StdOutput              string      `json:"std_output"`
	LastTestcase           string      `json:"last_testcase"`
	ExpectedOutput         string      `json:"expected_output"`
	Finished               bool        `json:"finished"`
}

func newProblemOfTheDayResponse(data []byte) (*ProblemOfTheDayResponse, error) {
    v := ProblemOfTheDayResponse{}
    err := json.Unmarshal(data,&v)
    return &v, err
}

func newDownloadProblemResponse(data []byte) (*DownloadProblemResponse, error){
    v := DownloadProblemResponse{}
    err := json.Unmarshal(data,&v)
    return &v, err
}
