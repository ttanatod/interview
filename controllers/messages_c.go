package controllers

import (
	"fmt"
	"net/http"
	"time"

	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DBController struct {
	Database *gorm.DB
}

// func (db *DBController) GetAllMessage(c *gin.Context) {

// 	lastTime := c.Request.Header.Get("last_time")
// 	fmt.Printf("lastTime: %v\n", lastTime)
// 	model := db.Database.Model(&models.Messages{}).Where("update_at > ?", lastTime)
// 	pg := paginate.New(&paginate.Config{
// 		DefaultSize: 50000,
// 	})

// 	c.JSON(200, pg.Response(model, c.Request, &[]models.Messages{}))
// }

// func (db *DBController) CreateMessage(c *gin.Context) {
// 	var message models.Messages

// 	c.ShouldBind(&message)

// 	result := db.Database.Create(&message)

// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"meassage": "UUID already exist"})
// 	} else {
// 		c.JSON(http.StatusCreated, gin.H{"results": &message})
// 	}
// }

func (db *DBController) GetServerTime(c *gin.Context) {
	dt := time.Now().UTC()
	fmt.Printf("dt: %v\n", dt)
	c.JSON(200, gin.H{"results": dt})
}

func (db *DBController) RegisterMember(c *gin.Context) {
	var member models.Member

	c.ShouldBind(&member)

	result := db.Database.Create(&member)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Telephone already exist"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"results": &member})
	}
}

func (db *DBController) RegisterField(c *gin.Context) {
	var field models.Field

	c.ShouldBind(&field)

	result := db.Database.Create(&field)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Name already exist"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"results": &field})
	}
}

func (db *DBController) GetAllMember(c *gin.Context) {
	var members []models.Member
	result := db.Database.Find(&members)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Not Found"})
	} else {
		c.JSON(200, gin.H{"results": &members})
	}
}

func (db *DBController) GetAllField(c *gin.Context) {
	var fields []models.Field
	result := db.Database.Find(&fields)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Not Found"})
	} else {
		c.JSON(200, gin.H{"results": &fields})
	}
}

// not done
func (db *DBController) RentField(c *gin.Context) {
	var rent models.RentRecord

	var rentRequest RentRequest
	c.ShouldBind(&rentRequest)

	//get member
	var member models.Member
	result := db.Database.Where("name = ?", rentRequest.MemberName).First(&member)

	//get field
	var field models.Field
	result = db.Database.Where("name = ?", rentRequest.FieldName).First(&field)

	//check condition user can not book one more fields at the same time
	var r models.RentRecord
	c1 := db.Database.Find(&r).Where("start >= ? AND end <= ? AND member_id = ?", rentRequest.Start, rentRequest.Start, member.ID)
	c2 := db.Database.Find(&r).Where("start >= ? AND end <= ? AND member_id = ?", rentRequest.End, rentRequest.End, member.ID)
	if c1.RowsAffected != 0 || c2.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "You already booked another field"})
	}

	//check condition is the field alreadry booked
	c1 = db.Database.Find(&r).Where("start >= ? AND end <= ? AND field_name = ?", rentRequest.Start, rentRequest.Start, field.Name)
	c2 = db.Database.Find(&r).Where("start >= ? AND end <= ? AND field_name = ?", rentRequest.End, rentRequest.End, field.Name)
	if c1.RowsAffected != 0 || c2.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Field already booked"})
	}

	rent = models.RentRecord{
		Member: member,
		Field:  field,
		Start:  rentRequest.Start,
		End:    rentRequest.End,
	}

	result = db.Database.Create(&rent)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Can't book"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"results": &rent})
	}
}

// todo
// func (db *DBController) GetRentRecordByMemberName(c *gin.Context) {
//  get member_name from request
//	result := db.Database.Find(&members).Where("member_name = ?", member_name)
//  return the result
// }

type RentRequest struct {
	MemberName string    `json:"member_name"`
	FieldName  string    `json:"field_name"`
	Start      time.Time `json:"start_time"`
	End        time.Time `json:"end_time"`
}
