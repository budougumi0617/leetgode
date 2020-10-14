package leetgode

import (
	"encoding/json"
	"strconv"
)

type GetQuestionDetailResult struct {
	Data QdrData `json:"data"`
}

type QdrData struct {
	IsCurrentUserAuthenticated bool     `json:"isCurrentUserAuthenticated"`
	Question                   Question `json:"question"`
}

type Question struct {
	Slug               string      `json:"-"`
	Referer            string      `json:"-"`
	FrontendQuestionID int         `json:"-"`
	QuestionID         string      `json:"questionId"`
	Content            string      `json:"content"`
	Stats              string      `json:"stats"`
	CodeDefinition     Codes       `json:"codeDefinition"`
	SampleTestCase     string      `json:"sampleTestCase"`
	EnableRunCode      bool        `json:"enableRunCode"`
	MetaData           string      `json:"metaData"`
	TranslatedContent  interface{} `json:"translatedContent"`
}

// Code the struct of leetcode codes.
type Code struct {
	Text        string `json:"text"`
	Value       string `json:"value"`
	DefaultCode string `json:"defaultCode"`
}

// Codes the slice of Code
type Codes []*Code

func (c *Codes) UnmarshalJSON(b []byte) error {
	var cs []*Code
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(s), &cs); err != nil {
		return err
	}
	*c = append(*c, cs...)
	return nil
}
