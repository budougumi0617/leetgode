package leetgode

// SolutionRequest is the parameters for below URL.
// https://leetcode.com/problems/${title_slug}/interpret_solution/
type SolutionRequest struct {
	Lang       string `json:"lang"` // golang
	QuestionID string `json:"question_id"`
	TestMode   string `json:"test_mode"` // false
	Name       string `json:"name"`
	DataInput  string `json:"data_input"` // ex: "[2,7,11,15]\n9", load from Question.SampleTestCase
	TypedCode  string `json:"typed_code"` // load from file
}

// SolutionResult is the response from below URL.
// https://leetcode.com/problems/${title_slug}/interpret_solution/
type SolutionResult struct {
	InterpretID string `json:"interpret_id"`
	TestCase    string `json:"test_case"`
}

// CheckResult is the response from below URL.
// https://leetcode.com/submissions/detail/${id}/check/
type CheckResult struct {
	StatusCode             int           `json:"status_code"`
	Lang                   string        `json:"lang"`
	CompileError           string        `json:"compile_error"`
	FullCompileError       string        `json:"full_compile_error"`
	RunSuccess             bool          `json:"run_success"`
	StatusRuntime          string        `json:"status_runtime"`
	Memory                 int           `json:"memory"`
	CodeAnswer             []string      `json:"code_answer"`
	CodeOutput             interface{}   `json:"code_output"` // stringのときと[]stringのときがありうる
	ElapsedTime            int           `json:"elapsed_time"`
	TaskFinishTime         int64         `json:"task_finish_time"`
	ExpectedStatusCode     int           `json:"expected_status_code"`
	ExpectedLang           string        `json:"expected_lang"`
	ExpectedRunSuccess     bool          `json:"expected_run_success"`
	ExpectedStatusRuntime  string        `json:"expected_status_runtime"`
	ExpectedMemory         int           `json:"expected_memory"`
	ExpectedCodeAnswer     []string      `json:"expected_code_answer"`
	ExpectedCodeOutput     []interface{} `json:"expected_code_output"`
	ExpectedElapsedTime    int           `json:"expected_elapsed_time"`
	ExpectedTaskFinishTime int64         `json:"expected_task_finish_time"`
	CorrectAnswer          bool          `json:"correct_answer"`
	TotalCorrect           int           `json:"total_correct"`
	TotalTestcases         int           `json:"total_testcases"`
	RuntimePercentile      interface{}   `json:"runtime_percentile"`
	StatusMemory           string        `json:"status_memory"`
	MemoryPercentile       interface{}   `json:"memory_percentile"`
	PrettyLang             string        `json:"pretty_lang"`
	SubmissionID           string        `json:"submission_id"`
	StatusMsg              string        `json:"status_msg"`
	State                  string        `json:"state"`
	QuestionID             string        `json:"question_id"`
	CompareResult          string        `json:"compare_result"`
	StdOutput              string        `json:"std_output"`
	LastTestcase           string        `json:"last_testcase"`
}

// SubmitRequest is the request parameters for below URL.
// https://leetcode.com/problems/${slug}/submit/
type SubmitRequest struct {
	Lang       string `json:"lang"`
	QuestionID string `json:"question_id"`
	TestMode   string `json:"test_mode"`
	Name       string `json:"name"`
	JudgeType  string `json:"judge_type"`
	TypedCode  string `json:"typed_code"`
}

// SubmitResult is the response from below URL.
// https://leetcode.com/problems/${slug}/submit/
type SubmitResult struct {
	SubmissionID int `json:"submission_id"`
}
