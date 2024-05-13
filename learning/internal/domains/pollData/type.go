package pollData

type AllData struct {
	StatusCode        int    `json:"status_code"`
	StatusDescription string `json:"status_description"`
	Data              Data   `json:"data"`
	Meta              Meta   `json:"meta"`
}

type Sports struct {
	MatchCount  string `json:"match_count"`
	SBinomen    string `json:"s_binomen"`
	SportID     string `json:"sport_id"`
	SportName   string `json:"sport_name"`
	SportTypeID string `json:"sport_type_id"`
	Created     string `json:"created"`
	Modified    string `json:"modified"`
}

type Countries struct {
	CountryID   string `json:"country_id"`
	CountryName string `json:"country_name"`
	CtBinomen   string `json:"ct_binomen"`
	MatchCount  string `json:"match_count"`
}

type Competitions struct {
	CBinomen        string `json:"c_binomen"`
	CPriority       string `json:"c_priority"`
	CompetitionID   string `json:"competition_id"`
	CompetitionName string `json:"competition_name"`
	CountryID       string `json:"country_id"`
	CountryName     string `json:"country_name"`
	CtBinomen       string `json:"ct_binomen"`
	MatchCount      string `json:"match_count"`
	SBinomen        string `json:"s_binomen"`
	SportID         string `json:"sport_id"`
	SportName       string `json:"sport_name"`
}

type Data struct {
	Sports       []Sports       `json:"sports"`
	Countries    []Countries    `json:"countries"`
	Competitions []Competitions `json:"competitions"`
	Live         string         `json:"live"`
	LiveCount    string         `json:"live_count"`
}

type Meta struct {
	Live      string `json:"live"`
	Src       string `json:"src"`
	CountryID string `json:"country_id"`
}
