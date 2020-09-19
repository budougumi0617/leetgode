package leetgode

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type LeetCode struct {
	BaseURL     string
	gqlEndpoint string
}

type LCOp func(*LeetCode) error

func NewLeetCode(ops ...LCOp) (*LeetCode, error) {
	lc := &LeetCode{
		BaseURL:     "https://leetcode.com",
		gqlEndpoint: "https://leetcode.com/graphql",
	}
	for _, op := range ops {
		if err := op(lc); err != nil {
			return nil, err
		}
	}
	return lc, nil
}

// TODO: クライアントを作る

// TODO: 共通メソッドとしてヘッダーを生成するメソッドをつくる
// TODO: オプションで外部からURLを変更できるようにする

type GetQuestionVariables struct {
	TitleSlug string `json:"titleSlug"`
}
type GetQuestionBody struct {
	Query         string               `json:"query"`
	Variables     GetQuestionVariables `json:"variables"`
	OperationName string               `json:"operationName"`
}

type GetQuestionResponseData struct {
	Question *Question `json:"question"`
}
type GetQuestionResponse struct {
	Data GetQuestionResponseData `json:"data"`
}

func (lc *LeetCode) GetQuestionByID(ctx context.Context, id int) (*Question, error) {
	ss, err := lc.GetStats(ctx)
	if err != nil {
		return nil, err
	}
	s, ok := ss[id]
	if !ok {
		return nil, err
	}
	q, err := lc.GetQuestion(ctx, s.Stat.QuestionTitleSlug)
	if err != nil {
		return nil, err
	}
	return q, nil
}

func (lc *LeetCode) GetQuestion(ctx context.Context, titleSlug string) (*Question, error) {
	query := `
query getQuestionDetail($titleSlug: String!) {
  isCurrentUserAuthenticated
  question(titleSlug: $titleSlug) {
    questionId
    content
    stats
    codeDefinition
    sampleTestCase
    enableRunCode
    metaData
    translatedContent
  }
}`
	body := GetQuestionBody{
		Query: query,
		Variables: GetQuestionVariables{
			TitleSlug: titleSlug,
		},
		OperationName: "getQuestionDetail",
	}
	jbody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, lc.gqlEndpoint, bytes.NewBuffer(jbody))
	if err != nil {
		return nil, err
	}
	referer := fmt.Sprintf("https://leetcode.com/problems/%s/description/", titleSlug)
	req.Header.Set("Referer", referer)
	// req.Header.Set("x-csrftoken", guestToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cache-Control", "no-cache")
	cli := http.Client{
		Timeout: 3 * time.Second,
	}
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GetQuestion got status %d", res.StatusCode)
	}
	var q Question
	var result GetQuestionResponse
	result.Data.Question = &q
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	q.Slug = titleSlug
	q.Referer = referer
	return &q, nil
}

// TODO: submitメソッドをつくる for exec

// curl https://leetcode.com/api/problems/algorithms/
func (lc *LeetCode) GetProblems(ctx context.Context) (*ProblemsResult, error) {
	cli := http.Client{
		Timeout: 3 * time.Second,
	}
	ep := lc.BaseURL + "/api/problems/algorithms/"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, err
	}
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GetProblems got status %d", res.StatusCode)
	}
	var pr ProblemsResult
	if err := json.NewDecoder(res.Body).Decode(&pr); err != nil {
		return nil, err
	}
	return &pr, nil
}

func (lc *LeetCode) GetStats(ctx context.Context) (map[int]*StatStatusPair, error) {
	ps, err := lc.GetProblems(ctx)
	if err != nil {
		return nil, err
	}
	ss := make(map[int]*StatStatusPair, ps.NumTotal)
	for _, sp := range ps.StatStatusPairs {
		ss[sp.Stat.QuestionID] = sp
	}
	return ss, nil
}

// TODO: testメソッドをつくる for test

//SUBCOMMANDS:
//data    Manage Cache [aliases: d]
//edit    Edit question by id [aliases: e]
//exec    Submit solution [aliases: x]
//list    List problems [aliases: l]
//pick    Pick a problem [aliases: p]
//stat    Show simple chart about submissions [aliases: s]
//test    Test question by id [aliases: t]
//help    Prints this message or the help of the given subcommand(s)
