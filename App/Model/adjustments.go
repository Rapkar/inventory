package Model

import (
	"fmt"
	"inventory/App/Boot"
	"inventory/App/Utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAdjustmentInput struct {
	UserID       uint64  `json:"user_id" binding:"required"`
	OffsetAmount float64 `json:"offset_amount" binding:"required"`
	Reason       string  `json:"reason"`
	CreatedBy    uint64  `json:"created_by" binding:"required"`
}

func CreateBalanceAdjustment(c *gin.Context) {
	var input CreateAdjustmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	adjustment := Boot.BalanceAdjustment{
		UserID:       input.UserID,
		OffsetAmount: input.OffsetAmount,
		Reason:       input.Reason,
		CreatedBy:    input.CreatedBy,
		CreatedAt:    Utility.CurrentTime(),
	}

	if err := Boot.DB().Create(&adjustment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "sucess", "data": adjustment})
}

func GetUserAdjustments(userID uint64) ([]Boot.BalanceAdjustment, error) {
	var adjustments []Boot.BalanceAdjustment

	err := Boot.DB().
		Preload("User").
		Preload("CreatedByUser").
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&adjustments).Error
	fmt.Println(userID, adjustments, "adjustments")
	return adjustments, err
}
func DeleteBalanceAdjustment(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id + "id")
	var adjustment Boot.BalanceAdjustment
	if err := Boot.DB().First(&adjustment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "مورد پیدا نشد"})
		return
	}

	if err := Boot.DB().Delete(&adjustment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "حذف با خطا مواجه شد"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "با موفقیت حذف شد"})
}
