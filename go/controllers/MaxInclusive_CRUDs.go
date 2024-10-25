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
var __MaxInclusive__dummysDeclaration__ models.MaxInclusive
var __MaxInclusive_time__dummyDeclaration time.Duration

var mutexMaxInclusive sync.Mutex

// An MaxInclusiveID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMaxInclusive updateMaxInclusive deleteMaxInclusive
type MaxInclusiveID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MaxInclusiveInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMaxInclusive updateMaxInclusive
type MaxInclusiveInput struct {
	// The MaxInclusive to submit or modify
	// in: body
	MaxInclusive *orm.MaxInclusiveAPI
}

// GetMaxInclusives
//
// swagger:route GET /maxinclusives maxinclusives getMaxInclusives
//
// # Get all maxinclusives
//
// Responses:
// default: genericError
//
//	200: maxinclusiveDBResponse
func (controller *Controller) GetMaxInclusives(c *gin.Context) {

	// source slice
	var maxinclusiveDBs []orm.MaxInclusiveDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMaxInclusives", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMaxInclusive.GetDB()

	_, err := db.Find(&maxinclusiveDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	maxinclusiveAPIs := make([]orm.MaxInclusiveAPI, 0)

	// for each maxinclusive, update fields from the database nullable fields
	for idx := range maxinclusiveDBs {
		maxinclusiveDB := &maxinclusiveDBs[idx]
		_ = maxinclusiveDB
		var maxinclusiveAPI orm.MaxInclusiveAPI

		// insertion point for updating fields
		maxinclusiveAPI.ID = maxinclusiveDB.ID
		maxinclusiveDB.CopyBasicFieldsToMaxInclusive_WOP(&maxinclusiveAPI.MaxInclusive_WOP)
		maxinclusiveAPI.MaxInclusivePointersEncoding = maxinclusiveDB.MaxInclusivePointersEncoding
		maxinclusiveAPIs = append(maxinclusiveAPIs, maxinclusiveAPI)
	}

	c.JSON(http.StatusOK, maxinclusiveAPIs)
}

// PostMaxInclusive
//
// swagger:route POST /maxinclusives maxinclusives postMaxInclusive
//
// Creates a maxinclusive
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMaxInclusive(c *gin.Context) {

	mutexMaxInclusive.Lock()
	defer mutexMaxInclusive.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMaxInclusives", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMaxInclusive.GetDB()

	// Validate input
	var input orm.MaxInclusiveAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create maxinclusive
	maxinclusiveDB := orm.MaxInclusiveDB{}
	maxinclusiveDB.MaxInclusivePointersEncoding = input.MaxInclusivePointersEncoding
	maxinclusiveDB.CopyBasicFieldsFromMaxInclusive_WOP(&input.MaxInclusive_WOP)

	_, err = db.Create(&maxinclusiveDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMaxInclusive.CheckoutPhaseOneInstance(&maxinclusiveDB)
	maxinclusive := backRepo.BackRepoMaxInclusive.Map_MaxInclusiveDBID_MaxInclusivePtr[maxinclusiveDB.ID]

	if maxinclusive != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), maxinclusive)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, maxinclusiveDB)
}

// GetMaxInclusive
//
// swagger:route GET /maxinclusives/{ID} maxinclusives getMaxInclusive
//
// Gets the details for a maxinclusive.
//
// Responses:
// default: genericError
//
//	200: maxinclusiveDBResponse
func (controller *Controller) GetMaxInclusive(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMaxInclusive", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMaxInclusive.GetDB()

	// Get maxinclusiveDB in DB
	var maxinclusiveDB orm.MaxInclusiveDB
	if _, err := db.First(&maxinclusiveDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var maxinclusiveAPI orm.MaxInclusiveAPI
	maxinclusiveAPI.ID = maxinclusiveDB.ID
	maxinclusiveAPI.MaxInclusivePointersEncoding = maxinclusiveDB.MaxInclusivePointersEncoding
	maxinclusiveDB.CopyBasicFieldsToMaxInclusive_WOP(&maxinclusiveAPI.MaxInclusive_WOP)

	c.JSON(http.StatusOK, maxinclusiveAPI)
}

// UpdateMaxInclusive
//
// swagger:route PATCH /maxinclusives/{ID} maxinclusives updateMaxInclusive
//
// # Update a maxinclusive
//
// Responses:
// default: genericError
//
//	200: maxinclusiveDBResponse
func (controller *Controller) UpdateMaxInclusive(c *gin.Context) {

	mutexMaxInclusive.Lock()
	defer mutexMaxInclusive.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMaxInclusive", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMaxInclusive.GetDB()

	// Validate input
	var input orm.MaxInclusiveAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var maxinclusiveDB orm.MaxInclusiveDB

	// fetch the maxinclusive
	_, err := db.First(&maxinclusiveDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	maxinclusiveDB.CopyBasicFieldsFromMaxInclusive_WOP(&input.MaxInclusive_WOP)
	maxinclusiveDB.MaxInclusivePointersEncoding = input.MaxInclusivePointersEncoding

	db, _ = db.Model(&maxinclusiveDB)
	_, err = db.Updates(&maxinclusiveDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	maxinclusiveNew := new(models.MaxInclusive)
	maxinclusiveDB.CopyBasicFieldsToMaxInclusive(maxinclusiveNew)

	// redeem pointers
	maxinclusiveDB.DecodePointers(backRepo, maxinclusiveNew)

	// get stage instance from DB instance, and call callback function
	maxinclusiveOld := backRepo.BackRepoMaxInclusive.Map_MaxInclusiveDBID_MaxInclusivePtr[maxinclusiveDB.ID]
	if maxinclusiveOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), maxinclusiveOld, maxinclusiveNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the maxinclusiveDB
	c.JSON(http.StatusOK, maxinclusiveDB)
}

// DeleteMaxInclusive
//
// swagger:route DELETE /maxinclusives/{ID} maxinclusives deleteMaxInclusive
//
// # Delete a maxinclusive
//
// default: genericError
//
//	200: maxinclusiveDBResponse
func (controller *Controller) DeleteMaxInclusive(c *gin.Context) {

	mutexMaxInclusive.Lock()
	defer mutexMaxInclusive.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMaxInclusive", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMaxInclusive.GetDB()

	// Get model if exist
	var maxinclusiveDB orm.MaxInclusiveDB
	if _, err := db.First(&maxinclusiveDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&maxinclusiveDB)

	// get an instance (not staged) from DB instance, and call callback function
	maxinclusiveDeleted := new(models.MaxInclusive)
	maxinclusiveDB.CopyBasicFieldsToMaxInclusive(maxinclusiveDeleted)

	// get stage instance from DB instance, and call callback function
	maxinclusiveStaged := backRepo.BackRepoMaxInclusive.Map_MaxInclusiveDBID_MaxInclusivePtr[maxinclusiveDB.ID]
	if maxinclusiveStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), maxinclusiveStaged, maxinclusiveDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
