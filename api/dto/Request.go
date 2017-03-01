package dto

type Request interface{
	FromValues(map[string][]string)
}
