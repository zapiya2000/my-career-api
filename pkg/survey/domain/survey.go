package domain

// SurveyWithQuestions
type SurveyWithQuestions struct {
	Id          int                   `json:"survey_id"`
	Name        string                `json:"survey_name"`
	Description interface{}           `json:"description,omitempty"`
	Questions   []QuestionWithAnswers `json:"questions"`
}

// QuestionWithAnswers
type QuestionWithAnswers struct {
	Id       int      `json:"question_id"`
	Question string   `json:"question"`
	Type     string   `json:"type"`
	Answers  []Answer `json:"answers,omitempty"`
}

// Answer
type Answer struct {
	Id    interface{} `json:"id"`
	Value interface{} `json:"value"`
}

// Survey
type Survey struct {
	Id          int         `json:"survey_id"`
	Name        string      `json:"survey_name"`
	Description interface{} `json:"description,omitempty"`
	IsActive    bool        `json:"is_active"`
}

// UserAnswer is used with the table user_answer
type UserAnswer struct {
	Email        string      `json:"email"`
	DocumentType string      `json:"document_type"`
	Document     string      `json:"document"`
	Survey       int         `json:"survey"`
	Question     int         `json:"question"`
	Answer       interface{} `json:"answer"`
}
