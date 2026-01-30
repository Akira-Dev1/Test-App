package domain

type Question struct {
	Title					*string		`json:"title,omitempty"`
	Content					*string		`json:"content,omitempty"`
	AnswerOptions			[]string	`json:"options,omitempty"`
	CorrectAnswerOption		*int		`json:"correct_option,omitempty"`
	Version					*int		`json:"version,omitempty"`
	ID						*int		`json:"id,omitempty"`
	AuthorID				*string		`json:"author_id,omitempty"`
}
