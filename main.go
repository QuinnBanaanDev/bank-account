package main

import (
    "fmt"
    "os"
)


var bankAccount = 12342

func withdrawAcnt(withdrawAmnt int) int { //withdraw function, return int
    bankAccount -= withdrawAmnt
    if bankAccount < 0 {
        bankAccount += withdrawAmnt //bankaccount cannot go negative
        return 1
    }
    return 0
}

func depositAcnt(depositAmnt int) {
    bankAccount += depositAmnt
    
}

func main() {
    var choice int
    var withdraw int
    var deposit int
    var exit bool = true
    for exit {
    
    fmt.Printf("%d \n 1. withdraw \n 2. deposit \n 3. exit\n", bankAccount)
    fmt.Scanln(&choice)
    switch choice {
        case 1:
            fmt.Scanln(&withdraw)
            if withdrawAcnt(withdraw) != 0 {
                fmt.Println("Not succesful: not enough money (you're broke)")
            }
        case 2:
            fmt.Scanln(&deposit)
            depositAcnt(deposit)
        case 3:
            exit = false
            os.Exit(0)
            
    }
    }
}