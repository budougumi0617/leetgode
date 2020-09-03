package leetgode

type LeetCode struct {
}

type LCOp func(*LeetCode) error

func NewLeetCode(ops []LCOp) (*LeetCode, error) {
	lc := &LeetCode{}
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

//POST https://leetcode.com/graphql
//Content-Type: application/json
//Cache-Control: no-cache
//Referer: https://leetcode.com/problems/validate-binary-search-tree/description/
//
//{
//"query": "query getQuestionDetail($titleSlug: String!) {\nisCurrentUserAuthenticated\n  question(titleSlug: $titleSlug) {\n questionId\ncontent\nstats\ncodeDefinition\nsampleTestCase\nenableRunCode\nmetaData\ntranslatedContent\n}\n}",
//
//"variables": {
//"titleSlug": "validate-binary-search-tree"
//},
//"operationName": "getQuestionDetail"
//}
// TODO: getQuestionメソッドをつくる

// TODO: submitメソッドをつくる

// curl https://leetcode.com/api/problems/algorithms/
// TODO: リストメソッドをつくる
