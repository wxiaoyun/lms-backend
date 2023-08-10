package worksheetview

import (
	viewmodel "lms-backend/internal/viewmodel/worksheet"
)

type WorksheetSummaryView struct {
	TotalCost           float64 `json:"total_cost"`
	TotalPrice          float64 `json:"total_price"`
	TotalProfit         float64 `json:"total_profit"`
	NegativeProfitCount int     `json:"negative_profit_count"`
	PositiveProfitCount int     `json:"positive_profit_count"`
}

func ToSummaryView(summary *viewmodel.WorksheetSummaryViewModel) *WorksheetSummaryView {
	return &WorksheetSummaryView{
		TotalCost:           summary.TotalCost,
		TotalPrice:          summary.TotalPrice,
		TotalProfit:         summary.TotalProfit,
		NegativeProfitCount: summary.NegativeProfitCount,
		PositiveProfitCount: summary.PositiveProfitCount,
	}
}
