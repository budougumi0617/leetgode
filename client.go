package leetgode

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type LeetCode struct {
	BaseURL        string
	gqlEndpoint    string
	session, token string
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

func fillAuth(session, token string) func(lc *LeetCode) error {
	return func(lc *LeetCode) error {
		lc.session = session
		lc.token = token
		return nil
	}
}

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
	var pair *StatStatusPair
	for _, p := range ss {
		if p.Stat.QuestionID == id {
			pair = p
		}
	}
	if pair == nil {
		return nil, fmt.Errorf("cannot find problem")
	}

	q, err := lc.GetQuestion(ctx, pair.Stat.QuestionTitleSlug)
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

func (lc *LeetCode) GetStats(ctx context.Context) ([]*StatStatusPair, error) {
	ps, err := lc.GetProblems(ctx)
	if err != nil {
		return nil, err
	}

	return ps.StatStatusPairs, nil
}

func (lc *LeetCode) fill(req *http.Request, q *Question) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("X-CsrfToken", lc.token)
	req.Header.Set("Cookie", fmt.Sprintf("LEETCODE_SESSION=%s; csrftoken=%s", lc.session, lc.token))
	// Need to set referer
}

func (lc *LeetCode) Test(ctx context.Context, q *Question, ans string) (string, error) {
	sr := &SolutionRequest{
		Lang:       "golang",
		QuestionID: q.QuestionID,
		TestMode:   "false",
		Name:       q.Slug,
		DataInput:  q.SampleTestCase,
		TypedCode:  ans,
	}
	b, err := json.Marshal(sr)
	if err != nil {
		log.Printf("failed Marshal %+v\n", err)
		return "", err
	}
	surl := lc.BaseURL + fmt.Sprintf("/problems/%s/interpret_solution/", q.Slug)
	log.Printf("send to %q", surl)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, surl, bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	lc.fill(req, q)
	req.Header.Set("Referer", surl)
	cli := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := cli.Do(req)
	if err != nil {
		log.Printf("failed Do %+v\n", err)
		return "", err
	}
	var res SolutionResult
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}
	log.Printf("result %+v\n", res)
	return res.InterpretID, nil
}

func (lc *LeetCode) Check(ctx context.Context, q *Question, id string) (*CheckResult, error) {
	curl := lc.BaseURL + fmt.Sprintf("/submissions/detail/%s/check/", id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, curl, nil)
	if err != nil {
		return nil, err
	}
	lc.fill(req, q)
	req.Header.Set("Referer", curl)
	cli := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := cli.Do(req)
	if err != nil {
		log.Printf("failed Do %+v\n", err)
		return nil, err
	}
	var res CheckResult
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	log.Printf("result %+v\n", res)
	return &res, nil
}

func (lc *LeetCode) Submit(ctx context.Context, q *Question, ans string) (string, error) {
	sr := &SubmitRequest{
		Lang:       "golang",
		QuestionID: q.QuestionID,
		TestMode:   "false",
		Name:       q.Slug,
		TypedCode:  ans,
	}
	b, err := json.Marshal(sr)
	if err != nil {
		log.Printf("failed Marshal %+v\n", err)
		return "", err
	}
	surl := lc.BaseURL + fmt.Sprintf("/problems/%s/submit/", q.Slug)
	log.Printf("send to %q", surl)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, surl, bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	lc.fill(req, q)
	req.Header.Set("Referer", surl)
	cli := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := cli.Do(req)
	if err != nil {
		log.Printf("failed Do %+v\n", err)
		return "", err
	}
	var res SubmitResult
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}
	log.Printf("result %+v\n", res)
	return strconv.Itoa(res.SubmissionID), nil
}

//SUBCOMMANDS:
//data    Manage Cache [aliases: d]
//edit    Edit question by id [aliases: e]
//exec    Submit solution [aliases: x]
//list    List problems [aliases: l]
//pick    Pick a problem [aliases: p]
//stat    Show simple chart about submissions [aliases: s]
//test    Test question by id [aliases: t]
//help    Prints this message or the help of the given subcommand(s)
