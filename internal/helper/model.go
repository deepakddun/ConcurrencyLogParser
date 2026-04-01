package helper

type Count struct {
	FileName     string
	NumErr       int
	NumInfo      int
	NumWarn      int
	NumDebug     int
	NumLines     int
	UnknownLines int
}

type TotalCount struct {
	TotalNumErr       int
	TotalNumInfo      int
	TotalNumWarn      int
	TotalNumDebug     int
	TotalNumLines     int
	TotalUnknownLines int
}

type Result struct {
	Count Count

	Err error
}
