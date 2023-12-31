package utils

import (
	"bufio"
	"fmt"
	"github.com/gookit/color"
	"github.com/zellyn/kooky"
	_ "github.com/zellyn/kooky/browser/all"
	"golang.org/x/net/html"
	"io"
	"os"
	"strings"
)

var TestCaseList []string

type SessionKey struct {
	Csrftoken       string
	LeetcodeSession string
}

func getCookiesFromBrowser() (*SessionKey, error) {

	token := SessionKey{}

	cookies := kooky.ReadCookies(
		kooky.Valid,
		kooky.DomainHasSuffix(`leetcode.com`),
		kooky.Name("LEETCODE_SESSION"),
	)

	if len(cookies) == 0 {
		return nil, fmt.Errorf("Failed to find LEETCODE_SESSION cookie in any browser. Login into your leetcode account and try again. For now, the program only works if you logic through Chrome/Firefox/Safari browsers.")
	}

	token.LeetcodeSession = cookies[0].Value

	cookies = kooky.ReadCookies(
		kooky.Valid,
		kooky.DomainHasPrefix("leetcode.com"),
		kooky.Name("csrftoken"),
	)

	if len(cookies) == 0 {
		return nil, fmt.Errorf("Failed to find LEETCODE_SESSION cookie in any browser. Login into your leetcode account and try again. For now, the program only works if you logic through Chrome/Firefox/Safari browsers.")
	}

	token.Csrftoken = cookies[0].Value

	return &token, nil
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

	return &SessionKey{Csrftoken: csrftoken, LeetcodeSession: leetcode_session}, nil
}

func RenderHTML(problemName, problemID, html_content string) error {
	doc, err := html.Parse(strings.NewReader(html_content))

	if err != nil {
		return err
	}

	var output string
	var traverse func(*html.Node)

	traverse = func(n *html.Node) {
		if n.Type == html.TextNode {
			if strings.HasPrefix(n.Data, "\n") {
				n.Data = "\n"
			}
			output += n.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(doc)

	fileName := problemName + "_" + problemID + ".txt"
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.WriteString(file, output)

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

	if i := strings.LastIndex(fileName, "/"); i != -1 {
		fileName = fileName[i+1:]
	}

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

	TestCaseList = make([]string, 0, 3)

	idx := 0

	content = strings.TrimLeft(content[2:], "\t\n ")

	for content[:2] != "*/" {
		idx = strings.Index(content, ";")
		if idx == -1 {
			return "", fmt.Errorf("Couldn't fine ';'. Delimit your test cases with ';'")
		}

		testCases += content[:idx] + "\n"
		TestCaseList = append(TestCaseList, content[:idx]+"\n")
		content = strings.TrimLeft(content[idx+1:], "\t\n ")
	}

	testCases = strings.TrimRight(testCases, "\n\t ")
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
	fmt.Printf("%v", prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func ParseTestCaseFile(fileName string) (string, error) {

	file, err := os.Open(fileName)

	if err != nil {
		return "", nil
	}

	byte, err := io.ReadAll(file)

	if err != nil {
		return "", nil
	}

	testCases := string(byte)
	testCases = strings.TrimSpace(testCases)

	return testCases, nil
}
