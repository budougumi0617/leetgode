package leetgode

type Stat struct {
	QuestionID          int         `json:"question_id"`
	QuestionArticleLive interface{} `json:"question__article__live"`
	QuestionArticleSlug interface{} `json:"question__article__slug"`
	QuestionTitle       string      `json:"question__title"`
	QuestionTitleSlug   string      `json:"question__title_slug"`
	QuestionHide        bool        `json:"question__hide"`
	TotalAcs            int         `json:"total_acs"`
	TotalSubmitted      int         `json:"total_submitted"`
	FrontendQuestionID  int         `json:"frontend_question_id"`
	IsNewQuestion       bool        `json:"is_new_question"`
}

type Difficulty struct {
	Level int `json:"level"`
}

type StatStatusPair struct {
	Stat       Stat        `json:"stat"`
	Status     interface{} `json:"status"`
	Difficulty Difficulty  `json:"difficulty"`
	PaidOnly   bool        `json:"paid_only"`
	IsFavor    bool        `json:"is_favor"`
	Frequency  int         `json:"frequency"`
	Progress   int         `json:"progress"`
}

type DescriptionResult struct {
	UserName        string            `json:"user_name"`
	NumSolved       int               `json:"num_solved"`
	NumTotal        int               `json:"num_total"`
	AcEasy          int               `json:"ac_easy"`
	AcMedium        int               `json:"ac_medium"`
	AcHard          int               `json:"ac_hard"`
	StatStatusPairs []*StatStatusPair `json:"stat_status_pairs"`
	FrequencyHigh   int               `json:"frequency_high"`
	FrequencyMid    int               `json:"frequency_mid"`
	CategorySlug    string            `json:"category_slug"`
}
