package runscope

type Step struct {
	StepType         string              `json:"step_type"`
	ID               string              `json:"id"`
	Method           string              `json:"method"`
	URL              string              `json:"url"`
	Body             string              `json:"body"`
	Auth             Auth                `json:"auth"`
	Form             map[string][]string `json:"form"`
	Assertions       []Assertion         `json:"assertions"`
	Variables        []Variable          `json:"variables"`
	Headers          map[string][]string `json:"headers"`
	Scripts          []Script            `json:"scripts"`
	Note             string              `json:"note"`
	Duration         int                 `json:"duration"`
	Comparision      string              `json:"string"`
	RightValue       string              `json:"right_value"`
	LeftValue        string              `json:"left_value"`
	Steps            []Step              `json:"steps"`
	IntegrationID    string              `json:"integration_id"`
	SuiteID          string              `json:"suite_id"`
	TestID           string              `json:"test_id"`
	IsCustomStartURL bool                `json:"is_custom_start_url"`
}

type Assertion struct {
	Source     string      `json:"source"`
	Property   string      `json:"property"`
	Comparison string      `json:"comparison"`
	Value      interface{} `json:"value"`
}

type Auth struct {
	AuthType       string `json:"auth_type"`
	Username       string `json:"username,omitempty"`
	Password       string `json:"username,omitempty"`
	AccessToken    string `json:"access_token,omitempty"`
	TokenSecret    string `json:"token_secret,omitempty"`
	ConsumerKey    string `json:"consumer_key,omitempty"`
	ConsumerSecret string `json:"consumer_secret,omitempty"`
	SignatureType  string `json:"signature_type,omitempty"`
}

type Variable struct {
	Name     string `json:"name"`
	Source   string `json:"source"`
	Property string `json:"property"`
}

type Script struct {
	Value string `json:"value"`
}
