package dbtask

type DbTask struct {
	Id      int    `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func NewDbTask(id int, date, title, comment, repeat string) *DbTask {
	newDbTask := &DbTask{
		Id:      id,
		Date:    date,
		Title:   title,
		Comment: comment,
		Repeat:  repeat,
	}

	return newDbTask
}

func NewDbTaskWithoutId(date, title, comment, repeat string) *DbTask {
	newDbTask := &DbTask{
		Date:    date,
		Title:   title,
		Comment: comment,
		Repeat:  repeat,
	}

	return newDbTask
}
