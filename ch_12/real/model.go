package main

type Currency struct {
	Result               string             `json:"result"`
	Provider             string             `json:"provider"`
	Documentation        string             `json:"documentation"`
	Time_last_update_utc string             `json:"time_last_update_utc"`
	Time_next_update_utc string             `json:"time_next_update_utc"`
	Base_code            string             `json:"base_code"`
	Rates                map[string]float64 `json:"rates"`
}

type Fact struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Source    string `json:"source"`
	Language  string `json:"language"`
	SourceURL string `json:"source_url"`
}

type Joke struct {
	Category string `json:"category"`
	Type     string `json:"type"`

	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`

	Safe bool   `json:"safe"`
	ID   int64  `json:"id"`
	Lang string `json:"lang"`
}

type Result struct {
	Currency *Currency
	Fact     *Fact
	Joke     *Joke
	Err      error
}
