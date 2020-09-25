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
			},
		},
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

func TestLeetCode_GetQuestionByID(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		want    *Question
		wantErr bool
	}{
		{
			name: "GetQuestionByIDSuccess",
			id:   2,
			want: &Question{
				Referer:    "https://leetcode.com/problems/add-two-numbers/description/",
				QuestionID: "2",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			lc := &LeetCode{
				BaseURL:     "https://leetcode.com",
				gqlEndpoint: "https://leetcode.com/graphql",
			}
			got, err := lc.GetQuestionByID(context.TODO(), tt.id)
			if err != nil {
				t.Fatalf("GetQuestionByID faield: %v", err)
			}
			if diff := cmp.Diff(got.Referer, tt.want.Referer); diff != "" {
				t.Errorf("GetQuestionByID: there is diff (-got +want)\n%s", diff)
			}
		})
	}
}

func Test(t *testing.T) {
	t.Parallel()
	type args struct {
		q   *Question
		ans string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "動作確認用ケース",
			args: args{
				q: &Question{
					Slug:           "two-sum",
					QuestionID:     "1",
					SampleTestCase: "[2,7,11,15]\n9",
				},
				ans: "package main\n\n/*\n* @lc app=leetcode id=1 lang=golang\n*\n* [1] Two Sum\n*/\n// @lc code=start\nfunc twoSum(nums []int, target int) []int {\nl := make(map[int]int, len(nums))\n// answer - x = y\nfor i, n := range nums {\nif j, ok := l[n]; ok {\nreturn []int{j, i}\n}\nl[target-n] = i\n}\nreturn []int{}\n}\n\n// @lc code=end",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			lc := &LeetCode{
				BaseURL:     "https://leetcode.com",
				gqlEndpoint: "https://leetcode.com/graphql",
			}
			got, err := lc.Test(context.TODO(), tt.args.q, tt.args.ans)
			if err != nil {
				t.Fatalf("Test() error = %v", err)
			}
			if got != tt.want {
				t.Errorf("Test() got = %v, want %v", got, tt.want)
			}
		})
	}
}
