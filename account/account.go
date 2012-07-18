package account

import (
    "encoding/json"
    "github.com/scottferg/Dropbox-Go/session"
)

type Parameters struct {
    Locale string
}

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

func GetAccount(s session.Session, p *Parameters) (a Account, err error) {
    params := make(map[string]string)

    if p != nil {
        if p.Locale != "" {
            params["locale"] = p.Locale
        }
    }

    body, _, err := s.MakeApiRequest("account/info", params, session.GET)
    
    if err != nil {
        return
    }

    err = json.Unmarshal(body, &a)

    return
}
