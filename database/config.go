package database

type GetCustomersRequest struct {
	Limit  int `json:"limit" form:"limit" query:"limit"`
	Offset int `json:"offset" form:"offset" query:"offset"`
}

const (
	defaultDatabaseURI = "mongodb+srv://saul:1234@cluster0-ooeaq.mongodb.net/test?retryWrites=true&w=majority"
	defaultDbName      = "eatos"
	defaultTable       = "user"
)
