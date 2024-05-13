package types

type CarryPrize struct {
	PrizeId uint   `form:"prize_id" json:"prize_id"`
	Name    string `form:"name" json:"name"`
}
