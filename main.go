package main

import (
    "fmt"
    "os"
    "github.com/replit/database-go"
    "strconv"
)

var bankAccount int

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
    var (
        initAcnt,
        getBalance string
        intAccount, acntBalance int
    )
    var getErr error
    getBalance, getErr = database.Get("balance")
    if getErr != nil {
        fmt.Println(getErr.Error())
        return 1
    }
    intAccount, getErr = strconv.Atoi(getBalance)
    if getErr != nil {
        fmt.Println(getErr.Error())
        return 1
    }
    if intAccount == 0 {
        fmt.Println("Is your bank account initiated? (y/n)")
        fmt.Scanln(&initAcnt)

        if initAcnt == "y" {
            fmt.Scanln("%d", acntBalance)
            return acntBalance
        } 
    }
    return 0 
}

func main() {
    var choice, withdraw, deposit, setBalance, currentBalance int
    var bl string
    var exit bool = true
    var err error
    setBalance = initBalance()
    stringBalance := strconv.Itoa(setBalance) 
    if setBalance != 1 {
        database.Set("balance", stringBalance)
    }
    //reuse variable
    bl, err = database.Get("balance")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    currentBalance, err = strconv.Atoi(bl)



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