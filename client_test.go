package leetgode

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLeetCode_GetQuestion(t *testing.T) {
	tests := [...]struct {
		name      string
		titleSlug string
		want      *Question
	}{
		{
			name:      "GetQuestionSuccess",
			titleSlug: "add-two-numbers",
			want: &Question{
				Referer:    "https://leetcode.com/problems/add-two-numbers/description/",
				QuestionID: "2",
			}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			lc := &LeetCode{
				BaseURL:     "https://leetcode.com",
				gqlEndpoint: "https://leetcode.com/graphql",
			}
			got, err := lc.GetQuestion(context.TODO(), tt.titleSlug)
			if err != nil {
				t.Fatalf("GetQuestion faield: %v", err)
			}
			if diff := cmp.Diff(got.QuestionID, tt.want.QuestionID); diff != "" {
				t.Errorf("GetQuestion: there is diff (-got +want)\n%s", diff)
			}
		})
	}
}

func TestLeetCode_GetProblems(t *testing.T) {
	tests := [...]struct {
		name string
	}{
		{
			name: "getProblems",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &LeetCode{
				BaseURL:     "https://leetcode.com",
				gqlEndpoint: "https://leetcode.com/graphql",
			}
			got, err := lc.GetProblems(context.TODO())
			if err != nil {
				t.Fatalf("GetProblems faield: %v", err)
			}
			if len(got.StatStatusPairs) == 0 {
				t.Errorf("not find problems")
			}
		})
	}
}
