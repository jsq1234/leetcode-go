package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gookit/color"
	"golang.org/x/net/html"
)

var TestCaseList []string

type SessionKey struct {
    Csrftoken string
    LeetcodeSession string
}

func getSessionKey() (*SessionKey, error) {

    session_key := os.Getenv("LEETCODE_SESSION_KEY")

    if session_key == "" {
        return nil, fmt.Errorf("Couldn't find LEETCODE_SESSION_KEY env")
    }

    if !strings.Contains(session_key, ";LEETCODE_SESSION=") {
        return nil, fmt.Errorf("Invalid, did you put LEETCODE_SESSION=?")
    }

    if !strings.Contains(session_key, "csrftoken=") {
        return nil, fmt.Errorf("Invalid key, did you put csrftoken=****; ?")
    }

    idx := strings.Index(session_key, ";")

    if idx == -1 {
        return nil, fmt.Errorf("Error in parsing session key")
    }

    csrftoken := session_key[10:idx]
    leetcode_session := session_key[idx+18 : len(session_key)-1]

    return &SessionKey{ Csrftoken: csrftoken, LeetcodeSession: leetcode_session }, nil
}

func RenderHTML(html_content string) error {
/*
	cmd := exec.Command("lynx", "-stdin", "-dump")
	cmd.Stdin = strings.NewReader(html_content)

	file, err := os.Create("problem.txt")

	if err != nil {
		return err
	}

	defer file.Close()

	cmd.Stderr = os.Stdout
	cmd.Stdout = file

	err = cmd.Run()

	if err != nil {
		return err
	}

    return nil */
    
    doc, err := html.Parse(strings.NewReader(html_content))

    if err != nil {
        return err
    }

    var output string
    var traverse func(*html.Node) 
    
    traverse = func(n* html.Node){ 
        if n.Type == html.TextNode {
            if strings.HasPrefix(n.Data,"\n") {
                n.Data = "\n"
            }
            output += n.Data
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
           traverse(c)
        }
    }

    traverse(doc)

    file, err := os.Create("problem.txt")

    if err != nil {
        return err
    }

    defer file.Close()

    _, err = io.WriteString(file,output)

    if err != nil {
        return err
    }

    return nil
}

func GetTitleSlug(keyword string) string {
	keyword = strings.ToLower(keyword)
	list := strings.Split(keyword, " ")
	return strings.Join(list, "-")
}

func ParseFileName(fileName string) (questionID string, problemName string, lang string, err error) {

    if i := strings.LastIndex(fileName,"/"); i != -1 {
        fileName = fileName[i+1:]
    }
    
    fmt.Println(fileName)

	if !strings.Contains(fileName, ".") {
		err = fmt.Errorf("No extension found in the given file. Please write the extension, ex. .cpp, .js")
		return
	}

	if !strings.Contains(fileName, "_") {
		err = fmt.Errorf("No question ID found. Write the questionID after problem name with an underscore.\nEx: two-sum_ID.cpp")
	}

	idx := strings.Index(fileName, ".")

	lang = fileName[idx+1:]

	name := fileName[:idx]

	idx = strings.Index(name, "_")

	problemName = name[:idx]

	questionID = name[idx+1 : strings.Index(fileName, ".")]

	err = nil

	return
}

func ParseTestCases(content string) (string, error) {

	if content[:2] != "/*" {
		return "", fmt.Errorf(`Couldn't find /* at the beginning of the file.
		Make sure to enclose test cases in /* */ comment block at the beginning`)
	}

	var testCases string

	TestCaseList = make([]string,0,3)

	idx := 0

	content = strings.TrimLeft(content[2:],"\t\n ")

	for content[:2] != "*/" {
		idx = strings.Index(content, ";")
		if idx == -1 {
			return "", fmt.Errorf("Couldn't fine ';'. Delimit your test cases with ';'")
		}

		testCases += content[:idx] + "\n"
		TestCaseList = append(TestCaseList, content[:idx] + "\n")
		content = strings.TrimLeft(content[idx+1:], "\t\n ")
	}

	testCases = strings.TrimRight(testCases,"\n\t ")

	return testCases, nil
}

func Color(difficulty string) string {
    if difficulty == "Easy" {
        return color.Green.Sprint(difficulty)
    }
    if difficulty == "Medium" {
        return color.Yellow.Sprintf(difficulty)
    }
    if difficulty == "Hard" {
        return color.Red.Sprintf(difficulty)
    }
    return ""
}

func UserInput(prompt string) string {
    fmt.Printf("%v",prompt)
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    return scanner.Text()
}
