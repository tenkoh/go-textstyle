package textstyle

const (
	BOLD_LOWER = 0xf09d9039
	BOLD_UPPER = 0xf09d903f
	BOLD_DIGIT = 0xf09d9f5e
)

var Bold = NewTransformer(
	NewSimpleReplacer(
		BOLD_LOWER,
		BOLD_UPPER,
		BOLD_DIGIT,
	),
)
