package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
)

type CouponUseCase interface {
	AddCoupon(coupon domain.Coupon) (models.CouponResp, error)
	GetCoupon(int, int) ([]models.CouponResp, error)
	EditCoupon(edit domain.Coupon) (models.CouponResp, error)
}
