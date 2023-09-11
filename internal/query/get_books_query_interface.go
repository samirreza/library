package query

type GetBooksQueryInterface interface {
	Get(GetBooksInput) ([]GetBookOutput, error)
}
