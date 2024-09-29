package request

type ReviewRequest struct {
	Review string `json:"review" binding:"required"`
}
