# Leetcode-go

Leetcode-go is a CLI written in golang for retrieval and submission of problems hosted on https://leetcode.com/. You can search problems, download them, run them against sample/your own testcases and submit them to your leetcode profile.

## Installation 

If you have golang installed, you can install the app using ``go install``. Simply write this to in your terminal.
```
go install github.com/Manan-Prakash-Singh/leetcode-go@latest
```

## Authentication
There is no need to insert your username and password. Just make sure you are logged in into your accout in your browser. It searches through your browser cookies and finds the csrftoken and LEETCODE_SESSION token. For now, this works for firefox and chrome on Linux. To test if you are authenticated, type  
```
leetcode-go auth
```

Note : I still need to test it on MacOS and Windows.

## Downloading problem 

To download the problem, you can use ``leetcode-go problem "problem name in quotes" --lang cpp``. This will download if 
it can find the problem in cpp. But this requires you to know the full name of thee problem, which can be cumbersome,
therfore, you can use ``leetcode-go search "search term"`` instead to search for the problem and then it will
automatically ask you to language that you need to download the problem in. The problem statement is downloaded in a
``problem.txt`` file whereas the code snippet is downloaded as your chosen language.

## Submitting/Running your code

When you download the code using leetcode-go, it saves it in the form "longest-increasing-subsequence_300". This format is expected when the filename is parsed, so don't change this. Moreover, the sample test cases are written inside the program file, enclosed inside block comments /* */. Each test case is seperated by ``;``. You can add your own test cases in it but make sure to keep the block comment at the top of the file.  
```
/*

[2,7,11,15]
9;

[3,2,4]
6;

[3,3]
6;

*/
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        
    }
};
```
#### Note : The problem must be in the current directory.

For help, just type ``leetcode-go help``
```
Leetcode-go is a simple cli that can search, download, and submit problems
on leetcode through the command line

Usage:
  leetcode-go [command]

Available Commands:
  auth        Authenticate user using cookies
  help        Help about any command
  problem     Download the problem in the current directory
  qod         Get problem of the day
  run         Run your code with the test cases in your program file.
  search      search a problem
  submit      Submit your code to leetcode
  topic       Get a list of problems of a particular topic

Flags:
  -h, --help   help for leetcode-go

Use "leetcode-go [command] --help" for more information about a command.
```
