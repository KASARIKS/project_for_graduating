package taskhandlers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	dbtask "github.com/kasariks/project_for_graduating/internal/db/dbEntites/dbTask"
	"github.com/kasariks/project_for_graduating/internal/nextdate"
)

func getIdentifier(r *http.Request) (int, error) {
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		return 0, errors.New("no identifier")
	}

	intId, err := strconv.Atoi(id)

	return intId, err
}

func validateTask(dbTask dbtask.DbTask) error {
	_, err := strconv.Atoi(dbTask.Id)
	if err != nil {
		return err
	}

	if len(dbTask.Title) == 0 {
		return err
	}

	_, err = time.Parse(nextdate.DateFormat, dbTask.Date)
	if err != nil {
		return err
	}

	return nil
}
