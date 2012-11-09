package dropbox

import (
	"encoding/json"
)

type Account struct {
	ReferralLink string `json:"referral_link"`
	DisplayName  string `json:"display_name"`
	Uid          int    `json:"uid"`
	Country      string `json:"country"`
	QuotaInfo    struct {
		Shared int64 `json:"shared"`
		Quota  int64 `json:"quota"`
		Normal int64 `json:"normal"`
	} `json:"quota_info"`
}

func GetAccount(s Session, p *Parameters) (a Account, err error) {
	params := make(map[string]string)

	if p != nil && p.Locale != "" {
		params["locale"] = p.Locale
	}

	body, _, err := s.MakeApiRequest("account/info", params, GET)

	if err != nil {
		return
	}

	err = json.Unmarshal(body, &a)

	return
}
