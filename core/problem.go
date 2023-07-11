package core

import (
    "github.com/Manan-Prakash-Singh/leetcode-go/utils"
)

const query = `
query questionOfToday {
    activeDailyCodingChallengeQuestion {
        link
        question {
            title
            difficulty
        }
    }
}`
