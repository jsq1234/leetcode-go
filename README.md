# Leetcode-go

Leetcode-go is a CLI written in golang for retrieval and submission of problems hosted on https://leetcode.com/. You can search problems, download them, run them against sample/your own testcases and submit them to your leetcode profile.

## Installation 

If you have golang installed, you can install the app using ``go install``. Simply write this to in your terminal.
```go install github.com/Manan-Prakash-Singh/leetcode-go@latest```

## Authentication

In order for you to run/submit test cases, you must be authenticated. This can be done by creating and exporting an environment variable with the name ``LEETCODE_SESSION_KEY``. Put your csrftoken and LEETCODE_SESSION token, seperated by ``;`` and enclosed in double quotes.  
``export LEETCODE_SESSION_KEY="csrftoken=asdXsdsa......;LEETCODE_SESSION=asdsaASDDD....."``

These tokens can be found from your browser. Open the developer tools and click the network tab. Browse your profile to the see the network traffic. In your request headers, find the ``Cookie`` field. There you will find both ``csrftoken`` and ``LEETCODE_SESSION`` token. Copy them and export the environment variable.

Note: You'll have to always export this environment variable if you wish to submit/run your code to leetcode. To prevent this hasle, put the ``export LEETCODE_SESSION_KEY="..."`` line in your ~/.bashrc or ~/.zshrc file.

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
