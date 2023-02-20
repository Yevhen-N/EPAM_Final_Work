package app

import (
	"github.com/Yevhen-N/EPAM_Final_Work/pkg/apiv1"
	"github.com/Yevhen-N/EPAM_Final_Work/pkg/db/model"
)

func mapLog(row *model.Log) *apiv1.LogResponse {
	res := &apiv1.LogResponse{
		ID:     row.ID,
		UserID: row.UserID,
		Date:   row.Date,
		Action: row.Action,
	}
	return res
}
