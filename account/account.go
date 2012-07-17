package account

import (
    "encoding/json"
    "github.com/scottferg/Dropbox-Go/session"
)

type Account struct {
    ReferralLink string `json:"referral_link"`
    DisplayName  string `json:"display_name"`
    Uid          int    `json:"uid"` 
    Country      string `json:"country"`
    QuotaInfo struct {
        Shared int64 `json:"shared"`
        Quota  int64 `json:"quota"`
        Normal int64 `json:"normal"`
    } `json:"quota_info"`
}

func GetAccount(s session.Session) (a Account, err error) {
    body, err := s.MakeApiRequest("account/info", session.GET)
    
    if err != nil {
        return
    }

    err = json.Unmarshal(body, &a)

    return
}
