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
var __Renamed_RELATION_GROUP_TYPE_REF_1__dummysDeclaration__ models.Renamed_RELATION_GROUP_TYPE_REF_1
var __Renamed_RELATION_GROUP_TYPE_REF_1_time__dummyDeclaration time.Duration

var mutexRenamed_RELATION_GROUP_TYPE_REF_1 sync.Mutex

// An Renamed_RELATION_GROUP_TYPE_REF_1ID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRenamed_RELATION_GROUP_TYPE_REF_1 updateRenamed_RELATION_GROUP_TYPE_REF_1 deleteRenamed_RELATION_GROUP_TYPE_REF_1
type Renamed_RELATION_GROUP_TYPE_REF_1ID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Renamed_RELATION_GROUP_TYPE_REF_1Input is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRenamed_RELATION_GROUP_TYPE_REF_1 updateRenamed_RELATION_GROUP_TYPE_REF_1
type Renamed_RELATION_GROUP_TYPE_REF_1Input struct {
	// The Renamed_RELATION_GROUP_TYPE_REF_1 to submit or modify
	// in: body
	Renamed_RELATION_GROUP_TYPE_REF_1 *orm.Renamed_RELATION_GROUP_TYPE_REF_1API
}

// GetRenamed_RELATION_GROUP_TYPE_REF_1s
//
// swagger:route GET /renamed_relation_group_type_ref_1s renamed_relation_group_type_ref_1s getRenamed_RELATION_GROUP_TYPE_REF_1s
//
// # Get all renamed_relation_group_type_ref_1s
//
// Responses:
// default: genericError
//
//	200: renamed_relation_group_type_ref_1DBResponse
func (controller *Controller) GetRenamed_RELATION_GROUP_TYPE_REF_1s(c *gin.Context) {

	// source slice
	var renamed_relation_group_type_ref_1DBs []orm.Renamed_RELATION_GROUP_TYPE_REF_1DB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRenamed_RELATION_GROUP_TYPE_REF_1s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRenamed_RELATION_GROUP_TYPE_REF_1.GetDB()

	query := db.Find(&renamed_relation_group_type_ref_1DBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	renamed_relation_group_type_ref_1APIs := make([]orm.Renamed_RELATION_GROUP_TYPE_REF_1API, 0)

	// for each renamed_relation_group_type_ref_1, update fields from the database nullable fields
	for idx := range renamed_relation_group_type_ref_1DBs {
		renamed_relation_group_type_ref_1DB := &renamed_relation_group_type_ref_1DBs[idx]
		_ = renamed_relation_group_type_ref_1DB
		var renamed_relation_group_type_ref_1API orm.Renamed_RELATION_GROUP_TYPE_REF_1API

		// insertion point for updating fields
		renamed_relation_group_type_ref_1API.ID = renamed_relation_group_type_ref_1DB.ID
		renamed_relation_group_type_ref_1DB.CopyBasicFieldsToRenamed_RELATION_GROUP_TYPE_REF_1_WOP(&renamed_relation_group_type_ref_1API.Renamed_RELATION_GROUP_TYPE_REF_1_WOP)
		renamed_relation_group_type_ref_1API.Renamed_RELATION_GROUP_TYPE_REF_1PointersEncoding = renamed_relation_group_type_ref_1DB.Renamed_RELATION_GROUP_TYPE_REF_1PointersEncoding
		renamed_relation_group_type_ref_1APIs = append(renamed_relation_group_type_ref_1APIs, renamed_relation_group_type_ref_1API)
	}

	c.JSON(http.StatusOK, renamed_relation_group_type_ref_1APIs)
}

// PostRenamed_RELATION_GROUP_TYPE_REF_1
//
// swagger:route POST /renamed_relation_group_type_ref_1s renamed_relation_group_type_ref_1s postRenamed_RELATION_GROUP_TYPE_REF_1
//
// Creates a renamed_relation_group_type_ref_1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRenamed_RELATION_GROUP_TYPE_REF_1(c *gin.Context) {

	mutexRenamed_RELATION_GROUP_TYPE_REF_1.Lock()
	defer mutexRenamed_RELATION_GROUP_TYPE_REF_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRenamed_RELATION_GROUP_TYPE_REF_1s", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRenamed_RELATION_GROUP_TYPE_REF_1.GetDB()

	// Validate input
	var input orm.Renamed_RELATION_GROUP_TYPE_REF_1API

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create renamed_relation_group_type_ref_1
	renamed_relation_group_type_ref_1DB := orm.Renamed_RELATION_GROUP_TYPE_REF_1DB{}
	renamed_relation_group_type_ref_1DB.Renamed_RELATION_GROUP_TYPE_REF_1PointersEncoding = input.Renamed_RELATION_GROUP_TYPE_REF_1PointersEncoding
	renamed_relation_group_type_ref_1DB.CopyBasicFieldsFromRenamed_RELATION_GROUP_TYPE_REF_1_WOP(&input.Renamed_RELATION_GROUP_TYPE_REF_1_WOP)

	query := db.Create(&renamed_relation_group_type_ref_1DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRenamed_RELATION_GROUP_TYPE_REF_1.CheckoutPhaseOneInstance(&renamed_relation_group_type_ref_1DB)
	renamed_relation_group_type_ref_1 := backRepo.BackRepoRenamed_RELATION_GROUP_TYPE_REF_1.Map_Renamed_RELATION_GROUP_TYPE_REF_1DBID_Renamed_RELATION_GROUP_TYPE_REF_1Ptr[renamed_relation_group_type_ref_1DB.ID]

	if renamed_relation_group_type_ref_1 != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), renamed_relation_group_type_ref_1)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, renamed_relation_group_type_ref_1DB)
}

// GetRenamed_RELATION_GROUP_TYPE_REF_1
//
// swagger:route GET /renamed_relation_group_type_ref_1s/{ID} renamed_relation_group_type_ref_1s getRenamed_RELATION_GROUP_TYPE_REF_1
//
// Gets the details for a renamed_relation_group_type_ref_1.
//
// Responses:
// default: genericError
//
//	200: renamed_relation_group_type_ref_1DBResponse
func (controller *Controller) GetRenamed_RELATION_GROUP_TYPE_REF_1(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRenamed_RELATION_GROUP_TYPE_REF_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRenamed_RELATION_GROUP_TYPE_REF_1.GetDB()

	// Get renamed_relation_group_type_ref_1DB in DB
	var renamed_relation_group_type_ref_1DB orm.Renamed_RELATION_GROUP_TYPE_REF_1DB
	if err := db.First(&renamed_relation_group_type_ref_1DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var renamed_relation_group_type_ref_1API orm.Renamed_RELATION_GROUP_TYPE_REF_1API
	renamed_relation_group_type_ref_1API.ID = renamed_relation_group_type_ref_1DB.ID
	renamed_relation_group_type_ref_1API.Renamed_RELATION_GROUP_TYPE_REF_1PointersEncoding = renamed_relation_group_type_ref_1DB.Renamed_RELATION_GROUP_TYPE_REF_1PointersEncoding
	renamed_relation_group_type_ref_1DB.CopyBasicFieldsToRenamed_RELATION_GROUP_TYPE_REF_1_WOP(&renamed_relation_group_type_ref_1API.Renamed_RELATION_GROUP_TYPE_REF_1_WOP)

	c.JSON(http.StatusOK, renamed_relation_group_type_ref_1API)
}

// UpdateRenamed_RELATION_GROUP_TYPE_REF_1
//
// swagger:route PATCH /renamed_relation_group_type_ref_1s/{ID} renamed_relation_group_type_ref_1s updateRenamed_RELATION_GROUP_TYPE_REF_1
//
// # Update a renamed_relation_group_type_ref_1
//
// Responses:
// default: genericError
//
//	200: renamed_relation_group_type_ref_1DBResponse
func (controller *Controller) UpdateRenamed_RELATION_GROUP_TYPE_REF_1(c *gin.Context) {

	mutexRenamed_RELATION_GROUP_TYPE_REF_1.Lock()
	defer mutexRenamed_RELATION_GROUP_TYPE_REF_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRenamed_RELATION_GROUP_TYPE_REF_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRenamed_RELATION_GROUP_TYPE_REF_1.GetDB()

	// Validate input
	var input orm.Renamed_RELATION_GROUP_TYPE_REF_1API
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var renamed_relation_group_type_ref_1DB orm.Renamed_RELATION_GROUP_TYPE_REF_1DB

	// fetch the renamed_relation_group_type_ref_1
	query := db.First(&renamed_relation_group_type_ref_1DB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	renamed_relation_group_type_ref_1DB.CopyBasicFieldsFromRenamed_RELATION_GROUP_TYPE_REF_1_WOP(&input.Renamed_RELATION_GROUP_TYPE_REF_1_WOP)
	renamed_relation_group_type_ref_1DB.Renamed_RELATION_GROUP_TYPE_REF_1PointersEncoding = input.Renamed_RELATION_GROUP_TYPE_REF_1PointersEncoding

	query = db.Model(&renamed_relation_group_type_ref_1DB).Updates(renamed_relation_group_type_ref_1DB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	renamed_relation_group_type_ref_1New := new(models.Renamed_RELATION_GROUP_TYPE_REF_1)
	renamed_relation_group_type_ref_1DB.CopyBasicFieldsToRenamed_RELATION_GROUP_TYPE_REF_1(renamed_relation_group_type_ref_1New)

	// redeem pointers
	renamed_relation_group_type_ref_1DB.DecodePointers(backRepo, renamed_relation_group_type_ref_1New)

	// get stage instance from DB instance, and call callback function
	renamed_relation_group_type_ref_1Old := backRepo.BackRepoRenamed_RELATION_GROUP_TYPE_REF_1.Map_Renamed_RELATION_GROUP_TYPE_REF_1DBID_Renamed_RELATION_GROUP_TYPE_REF_1Ptr[renamed_relation_group_type_ref_1DB.ID]
	if renamed_relation_group_type_ref_1Old != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), renamed_relation_group_type_ref_1Old, renamed_relation_group_type_ref_1New)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the renamed_relation_group_type_ref_1DB
	c.JSON(http.StatusOK, renamed_relation_group_type_ref_1DB)
}

// DeleteRenamed_RELATION_GROUP_TYPE_REF_1
//
// swagger:route DELETE /renamed_relation_group_type_ref_1s/{ID} renamed_relation_group_type_ref_1s deleteRenamed_RELATION_GROUP_TYPE_REF_1
//
// # Delete a renamed_relation_group_type_ref_1
//
// default: genericError
//
//	200: renamed_relation_group_type_ref_1DBResponse
func (controller *Controller) DeleteRenamed_RELATION_GROUP_TYPE_REF_1(c *gin.Context) {

	mutexRenamed_RELATION_GROUP_TYPE_REF_1.Lock()
	defer mutexRenamed_RELATION_GROUP_TYPE_REF_1.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRenamed_RELATION_GROUP_TYPE_REF_1", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRenamed_RELATION_GROUP_TYPE_REF_1.GetDB()

	// Get model if exist
	var renamed_relation_group_type_ref_1DB orm.Renamed_RELATION_GROUP_TYPE_REF_1DB
	if err := db.First(&renamed_relation_group_type_ref_1DB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&renamed_relation_group_type_ref_1DB)

	// get an instance (not staged) from DB instance, and call callback function
	renamed_relation_group_type_ref_1Deleted := new(models.Renamed_RELATION_GROUP_TYPE_REF_1)
	renamed_relation_group_type_ref_1DB.CopyBasicFieldsToRenamed_RELATION_GROUP_TYPE_REF_1(renamed_relation_group_type_ref_1Deleted)

	// get stage instance from DB instance, and call callback function
	renamed_relation_group_type_ref_1Staged := backRepo.BackRepoRenamed_RELATION_GROUP_TYPE_REF_1.Map_Renamed_RELATION_GROUP_TYPE_REF_1DBID_Renamed_RELATION_GROUP_TYPE_REF_1Ptr[renamed_relation_group_type_ref_1DB.ID]
	if renamed_relation_group_type_ref_1Staged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), renamed_relation_group_type_ref_1Staged, renamed_relation_group_type_ref_1Deleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
