package jira

const (
	customFieldFormatOption = "option"
	customFieldFormatArray  = "array"
	customFieldFormatNumber = "number"
)

type customField map[string]interface{}

type fieldValue struct {
	Value string `json:"value"`
}

type (
	customFieldTypeOption fieldValue
	customFieldTypeArray  fieldValue
	customFieldTypeNumber float64
)
