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
var __MaxLength__dummysDeclaration__ models.MaxLength
var __MaxLength_time__dummyDeclaration time.Duration

var mutexMaxLength sync.Mutex

// An MaxLengthID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMaxLength updateMaxLength deleteMaxLength
type MaxLengthID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MaxLengthInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMaxLength updateMaxLength
type MaxLengthInput struct {
	// The MaxLength to submit or modify
	// in: body
	MaxLength *orm.MaxLengthAPI
}

// GetMaxLengths
//
// swagger:route GET /maxlengths maxlengths getMaxLengths
//
// # Get all maxlengths
//
// Responses:
// default: genericError
//
//	200: maxlengthDBResponse
func (controller *Controller) GetMaxLengths(c *gin.Context) {

	// source slice
	var maxlengthDBs []orm.MaxLengthDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMaxLengths", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMaxLength.GetDB()

	_, err := db.Find(&maxlengthDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	maxlengthAPIs := make([]orm.MaxLengthAPI, 0)

	// for each maxlength, update fields from the database nullable fields
	for idx := range maxlengthDBs {
		maxlengthDB := &maxlengthDBs[idx]
		_ = maxlengthDB
		var maxlengthAPI orm.MaxLengthAPI

		// insertion point for updating fields
		maxlengthAPI.ID = maxlengthDB.ID
		maxlengthDB.CopyBasicFieldsToMaxLength_WOP(&maxlengthAPI.MaxLength_WOP)
		maxlengthAPI.MaxLengthPointersEncoding = maxlengthDB.MaxLengthPointersEncoding
		maxlengthAPIs = append(maxlengthAPIs, maxlengthAPI)
	}

	c.JSON(http.StatusOK, maxlengthAPIs)
}

// PostMaxLength
//
// swagger:route POST /maxlengths maxlengths postMaxLength
//
// Creates a maxlength
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMaxLength(c *gin.Context) {

	mutexMaxLength.Lock()
	defer mutexMaxLength.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMaxLengths", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMaxLength.GetDB()

	// Validate input
	var input orm.MaxLengthAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create maxlength
	maxlengthDB := orm.MaxLengthDB{}
	maxlengthDB.MaxLengthPointersEncoding = input.MaxLengthPointersEncoding
	maxlengthDB.CopyBasicFieldsFromMaxLength_WOP(&input.MaxLength_WOP)

	_, err = db.Create(&maxlengthDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMaxLength.CheckoutPhaseOneInstance(&maxlengthDB)
	maxlength := backRepo.BackRepoMaxLength.Map_MaxLengthDBID_MaxLengthPtr[maxlengthDB.ID]

	if maxlength != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), maxlength)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, maxlengthDB)
}

// GetMaxLength
//
// swagger:route GET /maxlengths/{ID} maxlengths getMaxLength
//
// Gets the details for a maxlength.
//
// Responses:
// default: genericError
//
//	200: maxlengthDBResponse
func (controller *Controller) GetMaxLength(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMaxLength", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMaxLength.GetDB()

	// Get maxlengthDB in DB
	var maxlengthDB orm.MaxLengthDB
	if _, err := db.First(&maxlengthDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var maxlengthAPI orm.MaxLengthAPI
	maxlengthAPI.ID = maxlengthDB.ID
	maxlengthAPI.MaxLengthPointersEncoding = maxlengthDB.MaxLengthPointersEncoding
	maxlengthDB.CopyBasicFieldsToMaxLength_WOP(&maxlengthAPI.MaxLength_WOP)

	c.JSON(http.StatusOK, maxlengthAPI)
}

// UpdateMaxLength
//
// swagger:route PATCH /maxlengths/{ID} maxlengths updateMaxLength
//
// # Update a maxlength
//
// Responses:
// default: genericError
//
//	200: maxlengthDBResponse
func (controller *Controller) UpdateMaxLength(c *gin.Context) {

	mutexMaxLength.Lock()
	defer mutexMaxLength.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMaxLength", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMaxLength.GetDB()

	// Validate input
	var input orm.MaxLengthAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var maxlengthDB orm.MaxLengthDB

	// fetch the maxlength
	_, err := db.First(&maxlengthDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	maxlengthDB.CopyBasicFieldsFromMaxLength_WOP(&input.MaxLength_WOP)
	maxlengthDB.MaxLengthPointersEncoding = input.MaxLengthPointersEncoding

	db, _ = db.Model(&maxlengthDB)
	_, err = db.Updates(&maxlengthDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	maxlengthNew := new(models.MaxLength)
	maxlengthDB.CopyBasicFieldsToMaxLength(maxlengthNew)

	// redeem pointers
	maxlengthDB.DecodePointers(backRepo, maxlengthNew)

	// get stage instance from DB instance, and call callback function
	maxlengthOld := backRepo.BackRepoMaxLength.Map_MaxLengthDBID_MaxLengthPtr[maxlengthDB.ID]
	if maxlengthOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), maxlengthOld, maxlengthNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the maxlengthDB
	c.JSON(http.StatusOK, maxlengthDB)
}

// DeleteMaxLength
//
// swagger:route DELETE /maxlengths/{ID} maxlengths deleteMaxLength
//
// # Delete a maxlength
//
// default: genericError
//
//	200: maxlengthDBResponse
func (controller *Controller) DeleteMaxLength(c *gin.Context) {

	mutexMaxLength.Lock()
	defer mutexMaxLength.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMaxLength", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMaxLength.GetDB()

	// Get model if exist
	var maxlengthDB orm.MaxLengthDB
	if _, err := db.First(&maxlengthDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&maxlengthDB)

	// get an instance (not staged) from DB instance, and call callback function
	maxlengthDeleted := new(models.MaxLength)
	maxlengthDB.CopyBasicFieldsToMaxLength(maxlengthDeleted)

	// get stage instance from DB instance, and call callback function
	maxlengthStaged := backRepo.BackRepoMaxLength.Map_MaxLengthDBID_MaxLengthPtr[maxlengthDB.ID]
	if maxlengthStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), maxlengthStaged, maxlengthDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
