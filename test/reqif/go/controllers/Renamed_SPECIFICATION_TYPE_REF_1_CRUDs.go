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
var __Renamed_SPECIFICATION_TYPE_REF_1__dummysDeclaration__ models.Renamed_SPECIFICATION_TYPE_REF_1
var __Renamed_SPECIFICATION_TYPE_REF_1_time__dummyDeclaration time.Duration

var mutexRenamed_SPECIFICATION_TYPE_REF_1 sync.Mutex

// An Renamed_SPECIFICATION_TYPE_REF_1ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRenamed_SPECIFICATION_TYPE_REF_1 updateRenamed_SPECIFICATION_TYPE_REF_1 deleteRenamed_SPECIFICATION_TYPE_REF_1
type Renamed_SPECIFICATION_TYPE_REF_1ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Renamed_SPECIFICATION_TYPE_REF_1Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRenamed_SPECIFICATION_TYPE_REF_1 updateRenamed_SPECIFICATION_TYPE_REF_1
type Renamed_SPECIFICATION_TYPE_REF_1Input struct {
	// The Renamed_SPECIFICATION_TYPE_REF_1 to submit or modify
	// in: body
	Renamed_SPECIFICATION_TYPE_REF_1 *orm.Renamed_SPECIFICATION_TYPE_REF_1API
}

// GetRenamed_SPECIFICATION_TYPE_REF_1s
//
// swagger:route GET /renamed_specification_type_ref_1s renamed_specification_type_ref_1s getRenamed_SPECIFICATION_TYPE_REF_1s
//
// # Get all renamed_specification_type_ref_1s
//
// Responses:
// default: genericError
//
//	200: renamed_specification_type_ref_1DBResponse
func (controller *Controller) GetRenamed_SPECIFICATION_TYPE_REF_1s(c *gin.Context) {

	// source slice
	var renamed_specification_type_ref_1DBs []orm.Renamed_SPECIFICATION_TYPE_REF_1DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRenamed_SPECIFICATION_TYPE_REF_1s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRenamed_SPECIFICATION_TYPE_REF_1.GetDB()

	query := db.Find(&renamed_specification_type_ref_1DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	renamed_specification_type_ref_1APIs := make([]orm.Renamed_SPECIFICATION_TYPE_REF_1API, 0)

	// for each renamed_specification_type_ref_1, update fields from the database nullable fields
	for idx := range renamed_specification_type_ref_1DBs {
		renamed_specification_type_ref_1DB := &renamed_specification_type_ref_1DBs[idx]
		_ = renamed_specification_type_ref_1DB
		var renamed_specification_type_ref_1API orm.Renamed_SPECIFICATION_TYPE_REF_1API

		// insertion point for updating fields
		renamed_specification_type_ref_1API.ID = renamed_specification_type_ref_1DB.ID
		renamed_specification_type_ref_1DB.CopyBasicFieldsToRenamed_SPECIFICATION_TYPE_REF_1_WOP(&renamed_specification_type_ref_1API.Renamed_SPECIFICATION_TYPE_REF_1_WOP)
		renamed_specification_type_ref_1API.Renamed_SPECIFICATION_TYPE_REF_1PointersEncoding = renamed_specification_type_ref_1DB.Renamed_SPECIFICATION_TYPE_REF_1PointersEncoding
		renamed_specification_type_ref_1APIs = append(renamed_specification_type_ref_1APIs, renamed_specification_type_ref_1API)
	}

	c.JSON(http.StatusOK, renamed_specification_type_ref_1APIs)
}

// PostRenamed_SPECIFICATION_TYPE_REF_1
//
// swagger:route POST /renamed_specification_type_ref_1s renamed_specification_type_ref_1s postRenamed_SPECIFICATION_TYPE_REF_1
//
// Creates a renamed_specification_type_ref_1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRenamed_SPECIFICATION_TYPE_REF_1(c *gin.Context) {

	mutexRenamed_SPECIFICATION_TYPE_REF_1.Lock()
	defer mutexRenamed_SPECIFICATION_TYPE_REF_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRenamed_SPECIFICATION_TYPE_REF_1s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRenamed_SPECIFICATION_TYPE_REF_1.GetDB()

	// Validate input
	var input orm.Renamed_SPECIFICATION_TYPE_REF_1API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create renamed_specification_type_ref_1
	renamed_specification_type_ref_1DB := orm.Renamed_SPECIFICATION_TYPE_REF_1DB{}
	renamed_specification_type_ref_1DB.Renamed_SPECIFICATION_TYPE_REF_1PointersEncoding = input.Renamed_SPECIFICATION_TYPE_REF_1PointersEncoding
	renamed_specification_type_ref_1DB.CopyBasicFieldsFromRenamed_SPECIFICATION_TYPE_REF_1_WOP(&input.Renamed_SPECIFICATION_TYPE_REF_1_WOP)

	query := db.Create(&renamed_specification_type_ref_1DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRenamed_SPECIFICATION_TYPE_REF_1.CheckoutPhaseOneInstance(&renamed_specification_type_ref_1DB)
	renamed_specification_type_ref_1 := backRepo.BackRepoRenamed_SPECIFICATION_TYPE_REF_1.Map_Renamed_SPECIFICATION_TYPE_REF_1DBID_Renamed_SPECIFICATION_TYPE_REF_1Ptr[renamed_specification_type_ref_1DB.ID]

	if renamed_specification_type_ref_1 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), renamed_specification_type_ref_1)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, renamed_specification_type_ref_1DB)
}

// GetRenamed_SPECIFICATION_TYPE_REF_1
//
// swagger:route GET /renamed_specification_type_ref_1s/{ID} renamed_specification_type_ref_1s getRenamed_SPECIFICATION_TYPE_REF_1
//
// Gets the details for a renamed_specification_type_ref_1.
//
// Responses:
// default: genericError
//
//	200: renamed_specification_type_ref_1DBResponse
func (controller *Controller) GetRenamed_SPECIFICATION_TYPE_REF_1(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRenamed_SPECIFICATION_TYPE_REF_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRenamed_SPECIFICATION_TYPE_REF_1.GetDB()

	// Get renamed_specification_type_ref_1DB in DB
	var renamed_specification_type_ref_1DB orm.Renamed_SPECIFICATION_TYPE_REF_1DB
	if err := db.First(&renamed_specification_type_ref_1DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var renamed_specification_type_ref_1API orm.Renamed_SPECIFICATION_TYPE_REF_1API
	renamed_specification_type_ref_1API.ID = renamed_specification_type_ref_1DB.ID
	renamed_specification_type_ref_1API.Renamed_SPECIFICATION_TYPE_REF_1PointersEncoding = renamed_specification_type_ref_1DB.Renamed_SPECIFICATION_TYPE_REF_1PointersEncoding
	renamed_specification_type_ref_1DB.CopyBasicFieldsToRenamed_SPECIFICATION_TYPE_REF_1_WOP(&renamed_specification_type_ref_1API.Renamed_SPECIFICATION_TYPE_REF_1_WOP)

	c.JSON(http.StatusOK, renamed_specification_type_ref_1API)
}

// UpdateRenamed_SPECIFICATION_TYPE_REF_1
//
// swagger:route PATCH /renamed_specification_type_ref_1s/{ID} renamed_specification_type_ref_1s updateRenamed_SPECIFICATION_TYPE_REF_1
//
// # Update a renamed_specification_type_ref_1
//
// Responses:
// default: genericError
//
//	200: renamed_specification_type_ref_1DBResponse
func (controller *Controller) UpdateRenamed_SPECIFICATION_TYPE_REF_1(c *gin.Context) {

	mutexRenamed_SPECIFICATION_TYPE_REF_1.Lock()
	defer mutexRenamed_SPECIFICATION_TYPE_REF_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRenamed_SPECIFICATION_TYPE_REF_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRenamed_SPECIFICATION_TYPE_REF_1.GetDB()

	// Validate input
	var input orm.Renamed_SPECIFICATION_TYPE_REF_1API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var renamed_specification_type_ref_1DB orm.Renamed_SPECIFICATION_TYPE_REF_1DB

	// fetch the renamed_specification_type_ref_1
	query := db.First(&renamed_specification_type_ref_1DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	renamed_specification_type_ref_1DB.CopyBasicFieldsFromRenamed_SPECIFICATION_TYPE_REF_1_WOP(&input.Renamed_SPECIFICATION_TYPE_REF_1_WOP)
	renamed_specification_type_ref_1DB.Renamed_SPECIFICATION_TYPE_REF_1PointersEncoding = input.Renamed_SPECIFICATION_TYPE_REF_1PointersEncoding

	query = db.Model(&renamed_specification_type_ref_1DB).Updates(renamed_specification_type_ref_1DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	renamed_specification_type_ref_1New := new(models.Renamed_SPECIFICATION_TYPE_REF_1)
	renamed_specification_type_ref_1DB.CopyBasicFieldsToRenamed_SPECIFICATION_TYPE_REF_1(renamed_specification_type_ref_1New)

	// redeem pointers
	renamed_specification_type_ref_1DB.DecodePointers(backRepo, renamed_specification_type_ref_1New)

	// get stage instance from DB instance, and call callback function
	renamed_specification_type_ref_1Old := backRepo.BackRepoRenamed_SPECIFICATION_TYPE_REF_1.Map_Renamed_SPECIFICATION_TYPE_REF_1DBID_Renamed_SPECIFICATION_TYPE_REF_1Ptr[renamed_specification_type_ref_1DB.ID]
	if renamed_specification_type_ref_1Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), renamed_specification_type_ref_1Old, renamed_specification_type_ref_1New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the renamed_specification_type_ref_1DB
	c.JSON(http.StatusOK, renamed_specification_type_ref_1DB)
}

// DeleteRenamed_SPECIFICATION_TYPE_REF_1
//
// swagger:route DELETE /renamed_specification_type_ref_1s/{ID} renamed_specification_type_ref_1s deleteRenamed_SPECIFICATION_TYPE_REF_1
//
// # Delete a renamed_specification_type_ref_1
//
// default: genericError
//
//	200: renamed_specification_type_ref_1DBResponse
func (controller *Controller) DeleteRenamed_SPECIFICATION_TYPE_REF_1(c *gin.Context) {

	mutexRenamed_SPECIFICATION_TYPE_REF_1.Lock()
	defer mutexRenamed_SPECIFICATION_TYPE_REF_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRenamed_SPECIFICATION_TYPE_REF_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRenamed_SPECIFICATION_TYPE_REF_1.GetDB()

	// Get model if exist
	var renamed_specification_type_ref_1DB orm.Renamed_SPECIFICATION_TYPE_REF_1DB
	if err := db.First(&renamed_specification_type_ref_1DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&renamed_specification_type_ref_1DB)

	// get an instance (not staged) from DB instance, and call callback function
	renamed_specification_type_ref_1Deleted := new(models.Renamed_SPECIFICATION_TYPE_REF_1)
	renamed_specification_type_ref_1DB.CopyBasicFieldsToRenamed_SPECIFICATION_TYPE_REF_1(renamed_specification_type_ref_1Deleted)

	// get stage instance from DB instance, and call callback function
	renamed_specification_type_ref_1Staged := backRepo.BackRepoRenamed_SPECIFICATION_TYPE_REF_1.Map_Renamed_SPECIFICATION_TYPE_REF_1DBID_Renamed_SPECIFICATION_TYPE_REF_1Ptr[renamed_specification_type_ref_1DB.ID]
	if renamed_specification_type_ref_1Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), renamed_specification_type_ref_1Staged, renamed_specification_type_ref_1Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
