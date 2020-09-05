package leetgode

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLeetCode_GetQuestion(t *testing.T) {
	tests := []struct {
		name      string
		titleSlug string
		want      *Question
	}{
		{
			name:      "GetQuestionSuccess",
			titleSlug: "add-two-numbers",
			want: &Question{
				QuestionID:        "",
				Content:           "",
				Stats:             "",
				CodeDefinition:    nil,
				SampleTestCase:    "",
				EnableRunCode:     false,
				MetaData:          "",
				TranslatedContent: nil,
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
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GetQuestion: there is diff (-got +want)\n%s", diff)
			}
		})
	}
}
