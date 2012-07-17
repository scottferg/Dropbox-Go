package main

import (
    "fmt"
    "io/ioutil"
    "github.com/scottferg/Dropbox-Go/session"
    "github.com/scottferg/Dropbox-Go/account"
    "github.com/scottferg/Dropbox-Go/files"
)

func main() {
    s := session.Session{
        AppKey: "3bvxdbph6b0vtks",
        AppSecret: "01l0an50qemvz9u",
        AccessType: "app_folder",
        Token: session.AccessToken{
            Secret: "a0727z0kybebpzc",
            Key: "yvrboxjs5benha3",
        },
    }

    a, err := account.GetAccount(s)

    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println(a.ReferralLink)
        fmt.Println(a.DisplayName)
    }

    file, err := ioutil.ReadFile("./test_form.pdf")

    if err != nil {
        fmt.Println(err.Error())
    } else {
        m, err := files.UploadFile(s, file, files.RootSandbox, "NERDS/test_form.pdf")

        if err != nil {
            fmt.Println(err.Error())
        } else {
            fmt.Println(m)
        }
    }

    fmt.Println("\n==========\n")

    m, err := files.GetMetadata(s, files.RootSandbox, "NERDS/test_form.pdf")

    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println(m)
    }
}
