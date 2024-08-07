// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongxsd/go/models"
	"github.com/fullstack-lang/gongxsd/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __MinInclusive__dummysDeclaration__ models.MinInclusive
var __MinInclusive_time__dummyDeclaration time.Duration

var mutexMinInclusive sync.Mutex

// An MinInclusiveID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMinInclusive updateMinInclusive deleteMinInclusive
type MinInclusiveID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MinInclusiveInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMinInclusive updateMinInclusive
type MinInclusiveInput struct {
	// The MinInclusive to submit or modify
	// in: body
	MinInclusive *orm.MinInclusiveAPI
}

// GetMinInclusives
//
// swagger:route GET /mininclusives mininclusives getMinInclusives
//
// # Get all mininclusives
//
// Responses:
// default: genericError
//
//	200: mininclusiveDBResponse
func (controller *Controller) GetMinInclusives(c *gin.Context) {

	// source slice
	var mininclusiveDBs []orm.MinInclusiveDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMinInclusives", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMinInclusive.GetDB()

	query := db.Find(&mininclusiveDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	mininclusiveAPIs := make([]orm.MinInclusiveAPI, 0)

	// for each mininclusive, update fields from the database nullable fields
	for idx := range mininclusiveDBs {
		mininclusiveDB := &mininclusiveDBs[idx]
		_ = mininclusiveDB
		var mininclusiveAPI orm.MinInclusiveAPI

		// insertion point for updating fields
		mininclusiveAPI.ID = mininclusiveDB.ID
		mininclusiveDB.CopyBasicFieldsToMinInclusive_WOP(&mininclusiveAPI.MinInclusive_WOP)
		mininclusiveAPI.MinInclusivePointersEncoding = mininclusiveDB.MinInclusivePointersEncoding
		mininclusiveAPIs = append(mininclusiveAPIs, mininclusiveAPI)
	}

	c.JSON(http.StatusOK, mininclusiveAPIs)
}

// PostMinInclusive
//
// swagger:route POST /mininclusives mininclusives postMinInclusive
//
// Creates a mininclusive
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMinInclusive(c *gin.Context) {

	mutexMinInclusive.Lock()
	defer mutexMinInclusive.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMinInclusives", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMinInclusive.GetDB()

	// Validate input
	var input orm.MinInclusiveAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create mininclusive
	mininclusiveDB := orm.MinInclusiveDB{}
	mininclusiveDB.MinInclusivePointersEncoding = input.MinInclusivePointersEncoding
	mininclusiveDB.CopyBasicFieldsFromMinInclusive_WOP(&input.MinInclusive_WOP)

	query := db.Create(&mininclusiveDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMinInclusive.CheckoutPhaseOneInstance(&mininclusiveDB)
	mininclusive := backRepo.BackRepoMinInclusive.Map_MinInclusiveDBID_MinInclusivePtr[mininclusiveDB.ID]

	if mininclusive != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), mininclusive)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, mininclusiveDB)
}

// GetMinInclusive
//
// swagger:route GET /mininclusives/{ID} mininclusives getMinInclusive
//
// Gets the details for a mininclusive.
//
// Responses:
// default: genericError
//
//	200: mininclusiveDBResponse
func (controller *Controller) GetMinInclusive(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMinInclusive", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMinInclusive.GetDB()

	// Get mininclusiveDB in DB
	var mininclusiveDB orm.MinInclusiveDB
	if err := db.First(&mininclusiveDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var mininclusiveAPI orm.MinInclusiveAPI
	mininclusiveAPI.ID = mininclusiveDB.ID
	mininclusiveAPI.MinInclusivePointersEncoding = mininclusiveDB.MinInclusivePointersEncoding
	mininclusiveDB.CopyBasicFieldsToMinInclusive_WOP(&mininclusiveAPI.MinInclusive_WOP)

	c.JSON(http.StatusOK, mininclusiveAPI)
}

// UpdateMinInclusive
//
// swagger:route PATCH /mininclusives/{ID} mininclusives updateMinInclusive
//
// # Update a mininclusive
//
// Responses:
// default: genericError
//
//	200: mininclusiveDBResponse
func (controller *Controller) UpdateMinInclusive(c *gin.Context) {

	mutexMinInclusive.Lock()
	defer mutexMinInclusive.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMinInclusive", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMinInclusive.GetDB()

	// Validate input
	var input orm.MinInclusiveAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var mininclusiveDB orm.MinInclusiveDB

	// fetch the mininclusive
	query := db.First(&mininclusiveDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	mininclusiveDB.CopyBasicFieldsFromMinInclusive_WOP(&input.MinInclusive_WOP)
	mininclusiveDB.MinInclusivePointersEncoding = input.MinInclusivePointersEncoding

	query = db.Model(&mininclusiveDB).Updates(mininclusiveDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	mininclusiveNew := new(models.MinInclusive)
	mininclusiveDB.CopyBasicFieldsToMinInclusive(mininclusiveNew)

	// redeem pointers
	mininclusiveDB.DecodePointers(backRepo, mininclusiveNew)

	// get stage instance from DB instance, and call callback function
	mininclusiveOld := backRepo.BackRepoMinInclusive.Map_MinInclusiveDBID_MinInclusivePtr[mininclusiveDB.ID]
	if mininclusiveOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), mininclusiveOld, mininclusiveNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the mininclusiveDB
	c.JSON(http.StatusOK, mininclusiveDB)
}

// DeleteMinInclusive
//
// swagger:route DELETE /mininclusives/{ID} mininclusives deleteMinInclusive
//
// # Delete a mininclusive
//
// default: genericError
//
//	200: mininclusiveDBResponse
func (controller *Controller) DeleteMinInclusive(c *gin.Context) {

	mutexMinInclusive.Lock()
	defer mutexMinInclusive.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMinInclusive", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMinInclusive.GetDB()

	// Get model if exist
	var mininclusiveDB orm.MinInclusiveDB
	if err := db.First(&mininclusiveDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&mininclusiveDB)

	// get an instance (not staged) from DB instance, and call callback function
	mininclusiveDeleted := new(models.MinInclusive)
	mininclusiveDB.CopyBasicFieldsToMinInclusive(mininclusiveDeleted)

	// get stage instance from DB instance, and call callback function
	mininclusiveStaged := backRepo.BackRepoMinInclusive.Map_MinInclusiveDBID_MinInclusivePtr[mininclusiveDB.ID]
	if mininclusiveStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), mininclusiveStaged, mininclusiveDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
