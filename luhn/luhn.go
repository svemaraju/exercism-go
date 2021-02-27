package luhn

import (
    "fmt"
    "strconv"
    "strings"
)

// Valid checks if a string is a valid card.
func Valid(card string) bool {
    card = strings.Replace(card, " ", "", -1)
    sum := 0
    doit := false
    if len(card) == 1 {
        return false
    }
    for i := range card {
        num, err := strconv.Atoi(string(card[len(card)-1-i]))
        if err == nil {
            if doit {
                if 2*num > 9 {
                    sum += 2*num - 9
                } else {
                    sum += 2 * num
                }
                doit = false
            } else {
                sum += num
                doit = true
            }
        } else {
            // Failed because of card string contains non digit.
            return false
        }
            
    }
    fmt.Println("SUM: ",sum, "CARD: ", card)
    if sum%10 == 0 {
        return true
    }
    return false
}
