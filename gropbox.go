package main

import (
    "fmt"
)

func main() {
    s := Session{
        AppKey: "3bvxdbph6b0vtks",
        AppSecret: "01l0an50qemvz9u",
        AccessType: "app_folder",
    }

    s.Token = AccessToken{
        Secret: "a0727z0kybebpzc",
        Key: "yvrboxjs5benha3",
    }

    a, err := GetAccount(s)

    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println(a.ReferralLink)
        fmt.Println(a.DisplayName)
    }
}
