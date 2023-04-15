package domain

import "github.com/hernanymotoso/finance-app-reports/dto"

type TransactionRepository interface {
	Search(reportID string, accountID string, initDate string, endDate string) (dto.SearchResponse, error)
}
