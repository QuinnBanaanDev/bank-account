package main

import (
    "fmt"
    "os"
    "github.com/replit/database-go"
    "strconv"
)



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

func initBalance() int {
    var initAcnt string
    var acntBalance int
    if database.Get("balance") == 0 {
        fmt.Println("Is your bank account initiated? (y/n)")
        fmt.Scanln(&initAcnt)
        if initAcnt == "y" {
            fmt.Scanf(&acntBalance)
            return acntBalance
        } 
    } 
    return 0
    
}

func main() {
    var choice, withdraw, deposit, setBalance, currentBalance int
    var exit bool = true
    setBalance = initBalance()
    if setBalance != 0 {
        database.Set("balance", strconv.Itoa(setBalance))
    } 
    currentBalance = strconv.Atoi(database.Get("balance"))




    
    
    //set bank account balance
    //  fmt.Println("current balance?")
    //  fmt.Scanln(&bankAccount)
    // = strconv.Itoa(bankAccount)
    
    

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