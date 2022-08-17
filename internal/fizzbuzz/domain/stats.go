package domain

// Statistics hold data object return when calling statistic API
type Statistics struct {
	Hits int32 `json:"hits"`
	Fizzbuz
}
