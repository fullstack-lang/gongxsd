// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongxsd/test/reqif/go/models"
	"github.com/fullstack-lang/gongxsd/test/reqif/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Renamed_ATTRIBUTE_VALUE_BOOLEAN_1__dummysDeclaration__ models.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1
var __Renamed_ATTRIBUTE_VALUE_BOOLEAN_1_time__dummyDeclaration time.Duration

var mutexRenamed_ATTRIBUTE_VALUE_BOOLEAN_1 sync.Mutex

// An Renamed_ATTRIBUTE_VALUE_BOOLEAN_1ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRenamed_ATTRIBUTE_VALUE_BOOLEAN_1 updateRenamed_ATTRIBUTE_VALUE_BOOLEAN_1 deleteRenamed_ATTRIBUTE_VALUE_BOOLEAN_1
type Renamed_ATTRIBUTE_VALUE_BOOLEAN_1ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Renamed_ATTRIBUTE_VALUE_BOOLEAN_1Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRenamed_ATTRIBUTE_VALUE_BOOLEAN_1 updateRenamed_ATTRIBUTE_VALUE_BOOLEAN_1
type Renamed_ATTRIBUTE_VALUE_BOOLEAN_1Input struct {
	// The Renamed_ATTRIBUTE_VALUE_BOOLEAN_1 to submit or modify
	// in: body
	Renamed_ATTRIBUTE_VALUE_BOOLEAN_1 *orm.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1API
}

// GetRenamed_ATTRIBUTE_VALUE_BOOLEAN_1s
//
// swagger:route GET /renamed_attribute_value_boolean_1s renamed_attribute_value_boolean_1s getRenamed_ATTRIBUTE_VALUE_BOOLEAN_1s
//
// # Get all renamed_attribute_value_boolean_1s
//
// Responses:
// default: genericError
//
//	200: renamed_attribute_value_boolean_1DBResponse
func (controller *Controller) GetRenamed_ATTRIBUTE_VALUE_BOOLEAN_1s(c *gin.Context) {

	// source slice
	var renamed_attribute_value_boolean_1DBs []orm.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRenamed_ATTRIBUTE_VALUE_BOOLEAN_1s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRenamed_ATTRIBUTE_VALUE_BOOLEAN_1.GetDB()

	query := db.Find(&renamed_attribute_value_boolean_1DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	renamed_attribute_value_boolean_1APIs := make([]orm.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1API, 0)

	// for each renamed_attribute_value_boolean_1, update fields from the database nullable fields
	for idx := range renamed_attribute_value_boolean_1DBs {
		renamed_attribute_value_boolean_1DB := &renamed_attribute_value_boolean_1DBs[idx]
		_ = renamed_attribute_value_boolean_1DB
		var renamed_attribute_value_boolean_1API orm.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1API

		// insertion point for updating fields
		renamed_attribute_value_boolean_1API.ID = renamed_attribute_value_boolean_1DB.ID
		renamed_attribute_value_boolean_1DB.CopyBasicFieldsToRenamed_ATTRIBUTE_VALUE_BOOLEAN_1_WOP(&renamed_attribute_value_boolean_1API.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1_WOP)
		renamed_attribute_value_boolean_1API.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1PointersEncoding = renamed_attribute_value_boolean_1DB.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1PointersEncoding
		renamed_attribute_value_boolean_1APIs = append(renamed_attribute_value_boolean_1APIs, renamed_attribute_value_boolean_1API)
	}

	c.JSON(http.StatusOK, renamed_attribute_value_boolean_1APIs)
}

// PostRenamed_ATTRIBUTE_VALUE_BOOLEAN_1
//
// swagger:route POST /renamed_attribute_value_boolean_1s renamed_attribute_value_boolean_1s postRenamed_ATTRIBUTE_VALUE_BOOLEAN_1
//
// Creates a renamed_attribute_value_boolean_1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRenamed_ATTRIBUTE_VALUE_BOOLEAN_1(c *gin.Context) {

	mutexRenamed_ATTRIBUTE_VALUE_BOOLEAN_1.Lock()
	defer mutexRenamed_ATTRIBUTE_VALUE_BOOLEAN_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRenamed_ATTRIBUTE_VALUE_BOOLEAN_1s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRenamed_ATTRIBUTE_VALUE_BOOLEAN_1.GetDB()

	// Validate input
	var input orm.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create renamed_attribute_value_boolean_1
	renamed_attribute_value_boolean_1DB := orm.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1DB{}
	renamed_attribute_value_boolean_1DB.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1PointersEncoding = input.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1PointersEncoding
	renamed_attribute_value_boolean_1DB.CopyBasicFieldsFromRenamed_ATTRIBUTE_VALUE_BOOLEAN_1_WOP(&input.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1_WOP)

	query := db.Create(&renamed_attribute_value_boolean_1DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRenamed_ATTRIBUTE_VALUE_BOOLEAN_1.CheckoutPhaseOneInstance(&renamed_attribute_value_boolean_1DB)
	renamed_attribute_value_boolean_1 := backRepo.BackRepoRenamed_ATTRIBUTE_VALUE_BOOLEAN_1.Map_Renamed_ATTRIBUTE_VALUE_BOOLEAN_1DBID_Renamed_ATTRIBUTE_VALUE_BOOLEAN_1Ptr[renamed_attribute_value_boolean_1DB.ID]

	if renamed_attribute_value_boolean_1 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), renamed_attribute_value_boolean_1)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, renamed_attribute_value_boolean_1DB)
}

// GetRenamed_ATTRIBUTE_VALUE_BOOLEAN_1
//
// swagger:route GET /renamed_attribute_value_boolean_1s/{ID} renamed_attribute_value_boolean_1s getRenamed_ATTRIBUTE_VALUE_BOOLEAN_1
//
// Gets the details for a renamed_attribute_value_boolean_1.
//
// Responses:
// default: genericError
//
//	200: renamed_attribute_value_boolean_1DBResponse
func (controller *Controller) GetRenamed_ATTRIBUTE_VALUE_BOOLEAN_1(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRenamed_ATTRIBUTE_VALUE_BOOLEAN_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRenamed_ATTRIBUTE_VALUE_BOOLEAN_1.GetDB()

	// Get renamed_attribute_value_boolean_1DB in DB
	var renamed_attribute_value_boolean_1DB orm.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1DB
	if err := db.First(&renamed_attribute_value_boolean_1DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var renamed_attribute_value_boolean_1API orm.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1API
	renamed_attribute_value_boolean_1API.ID = renamed_attribute_value_boolean_1DB.ID
	renamed_attribute_value_boolean_1API.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1PointersEncoding = renamed_attribute_value_boolean_1DB.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1PointersEncoding
	renamed_attribute_value_boolean_1DB.CopyBasicFieldsToRenamed_ATTRIBUTE_VALUE_BOOLEAN_1_WOP(&renamed_attribute_value_boolean_1API.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1_WOP)

	c.JSON(http.StatusOK, renamed_attribute_value_boolean_1API)
}

// UpdateRenamed_ATTRIBUTE_VALUE_BOOLEAN_1
//
// swagger:route PATCH /renamed_attribute_value_boolean_1s/{ID} renamed_attribute_value_boolean_1s updateRenamed_ATTRIBUTE_VALUE_BOOLEAN_1
//
// # Update a renamed_attribute_value_boolean_1
//
// Responses:
// default: genericError
//
//	200: renamed_attribute_value_boolean_1DBResponse
func (controller *Controller) UpdateRenamed_ATTRIBUTE_VALUE_BOOLEAN_1(c *gin.Context) {

	mutexRenamed_ATTRIBUTE_VALUE_BOOLEAN_1.Lock()
	defer mutexRenamed_ATTRIBUTE_VALUE_BOOLEAN_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRenamed_ATTRIBUTE_VALUE_BOOLEAN_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRenamed_ATTRIBUTE_VALUE_BOOLEAN_1.GetDB()

	// Validate input
	var input orm.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var renamed_attribute_value_boolean_1DB orm.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1DB

	// fetch the renamed_attribute_value_boolean_1
	query := db.First(&renamed_attribute_value_boolean_1DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	renamed_attribute_value_boolean_1DB.CopyBasicFieldsFromRenamed_ATTRIBUTE_VALUE_BOOLEAN_1_WOP(&input.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1_WOP)
	renamed_attribute_value_boolean_1DB.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1PointersEncoding = input.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1PointersEncoding

	query = db.Model(&renamed_attribute_value_boolean_1DB).Updates(renamed_attribute_value_boolean_1DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	renamed_attribute_value_boolean_1New := new(models.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1)
	renamed_attribute_value_boolean_1DB.CopyBasicFieldsToRenamed_ATTRIBUTE_VALUE_BOOLEAN_1(renamed_attribute_value_boolean_1New)

	// redeem pointers
	renamed_attribute_value_boolean_1DB.DecodePointers(backRepo, renamed_attribute_value_boolean_1New)

	// get stage instance from DB instance, and call callback function
	renamed_attribute_value_boolean_1Old := backRepo.BackRepoRenamed_ATTRIBUTE_VALUE_BOOLEAN_1.Map_Renamed_ATTRIBUTE_VALUE_BOOLEAN_1DBID_Renamed_ATTRIBUTE_VALUE_BOOLEAN_1Ptr[renamed_attribute_value_boolean_1DB.ID]
	if renamed_attribute_value_boolean_1Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), renamed_attribute_value_boolean_1Old, renamed_attribute_value_boolean_1New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the renamed_attribute_value_boolean_1DB
	c.JSON(http.StatusOK, renamed_attribute_value_boolean_1DB)
}

// DeleteRenamed_ATTRIBUTE_VALUE_BOOLEAN_1
//
// swagger:route DELETE /renamed_attribute_value_boolean_1s/{ID} renamed_attribute_value_boolean_1s deleteRenamed_ATTRIBUTE_VALUE_BOOLEAN_1
//
// # Delete a renamed_attribute_value_boolean_1
//
// default: genericError
//
//	200: renamed_attribute_value_boolean_1DBResponse
func (controller *Controller) DeleteRenamed_ATTRIBUTE_VALUE_BOOLEAN_1(c *gin.Context) {

	mutexRenamed_ATTRIBUTE_VALUE_BOOLEAN_1.Lock()
	defer mutexRenamed_ATTRIBUTE_VALUE_BOOLEAN_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRenamed_ATTRIBUTE_VALUE_BOOLEAN_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRenamed_ATTRIBUTE_VALUE_BOOLEAN_1.GetDB()

	// Get model if exist
	var renamed_attribute_value_boolean_1DB orm.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1DB
	if err := db.First(&renamed_attribute_value_boolean_1DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&renamed_attribute_value_boolean_1DB)

	// get an instance (not staged) from DB instance, and call callback function
	renamed_attribute_value_boolean_1Deleted := new(models.Renamed_ATTRIBUTE_VALUE_BOOLEAN_1)
	renamed_attribute_value_boolean_1DB.CopyBasicFieldsToRenamed_ATTRIBUTE_VALUE_BOOLEAN_1(renamed_attribute_value_boolean_1Deleted)

	// get stage instance from DB instance, and call callback function
	renamed_attribute_value_boolean_1Staged := backRepo.BackRepoRenamed_ATTRIBUTE_VALUE_BOOLEAN_1.Map_Renamed_ATTRIBUTE_VALUE_BOOLEAN_1DBID_Renamed_ATTRIBUTE_VALUE_BOOLEAN_1Ptr[renamed_attribute_value_boolean_1DB.ID]
	if renamed_attribute_value_boolean_1Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), renamed_attribute_value_boolean_1Staged, renamed_attribute_value_boolean_1Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
