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
var __TotalDigit__dummysDeclaration__ models.TotalDigit
var __TotalDigit_time__dummyDeclaration time.Duration

var mutexTotalDigit sync.Mutex

// An TotalDigitID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTotalDigit updateTotalDigit deleteTotalDigit
type TotalDigitID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TotalDigitInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTotalDigit updateTotalDigit
type TotalDigitInput struct {
	// The TotalDigit to submit or modify
	// in: body
	TotalDigit *orm.TotalDigitAPI
}

// GetTotalDigits
//
// swagger:route GET /totaldigits totaldigits getTotalDigits
//
// # Get all totaldigits
//
// Responses:
// default: genericError
//
//	200: totaldigitDBResponse
func (controller *Controller) GetTotalDigits(c *gin.Context) {

	// source slice
	var totaldigitDBs []orm.TotalDigitDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTotalDigits", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTotalDigit.GetDB()

	_, err := db.Find(&totaldigitDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	totaldigitAPIs := make([]orm.TotalDigitAPI, 0)

	// for each totaldigit, update fields from the database nullable fields
	for idx := range totaldigitDBs {
		totaldigitDB := &totaldigitDBs[idx]
		_ = totaldigitDB
		var totaldigitAPI orm.TotalDigitAPI

		// insertion point for updating fields
		totaldigitAPI.ID = totaldigitDB.ID
		totaldigitDB.CopyBasicFieldsToTotalDigit_WOP(&totaldigitAPI.TotalDigit_WOP)
		totaldigitAPI.TotalDigitPointersEncoding = totaldigitDB.TotalDigitPointersEncoding
		totaldigitAPIs = append(totaldigitAPIs, totaldigitAPI)
	}

	c.JSON(http.StatusOK, totaldigitAPIs)
}

// PostTotalDigit
//
// swagger:route POST /totaldigits totaldigits postTotalDigit
//
// Creates a totaldigit
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTotalDigit(c *gin.Context) {

	mutexTotalDigit.Lock()
	defer mutexTotalDigit.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTotalDigits", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTotalDigit.GetDB()

	// Validate input
	var input orm.TotalDigitAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create totaldigit
	totaldigitDB := orm.TotalDigitDB{}
	totaldigitDB.TotalDigitPointersEncoding = input.TotalDigitPointersEncoding
	totaldigitDB.CopyBasicFieldsFromTotalDigit_WOP(&input.TotalDigit_WOP)

	_, err = db.Create(&totaldigitDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTotalDigit.CheckoutPhaseOneInstance(&totaldigitDB)
	totaldigit := backRepo.BackRepoTotalDigit.Map_TotalDigitDBID_TotalDigitPtr[totaldigitDB.ID]

	if totaldigit != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), totaldigit)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, totaldigitDB)
}

// GetTotalDigit
//
// swagger:route GET /totaldigits/{ID} totaldigits getTotalDigit
//
// Gets the details for a totaldigit.
//
// Responses:
// default: genericError
//
//	200: totaldigitDBResponse
func (controller *Controller) GetTotalDigit(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTotalDigit", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTotalDigit.GetDB()

	// Get totaldigitDB in DB
	var totaldigitDB orm.TotalDigitDB
	if _, err := db.First(&totaldigitDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var totaldigitAPI orm.TotalDigitAPI
	totaldigitAPI.ID = totaldigitDB.ID
	totaldigitAPI.TotalDigitPointersEncoding = totaldigitDB.TotalDigitPointersEncoding
	totaldigitDB.CopyBasicFieldsToTotalDigit_WOP(&totaldigitAPI.TotalDigit_WOP)

	c.JSON(http.StatusOK, totaldigitAPI)
}

// UpdateTotalDigit
//
// swagger:route PATCH /totaldigits/{ID} totaldigits updateTotalDigit
//
// # Update a totaldigit
//
// Responses:
// default: genericError
//
//	200: totaldigitDBResponse
func (controller *Controller) UpdateTotalDigit(c *gin.Context) {

	mutexTotalDigit.Lock()
	defer mutexTotalDigit.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTotalDigit", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTotalDigit.GetDB()

	// Validate input
	var input orm.TotalDigitAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var totaldigitDB orm.TotalDigitDB

	// fetch the totaldigit
	_, err := db.First(&totaldigitDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	totaldigitDB.CopyBasicFieldsFromTotalDigit_WOP(&input.TotalDigit_WOP)
	totaldigitDB.TotalDigitPointersEncoding = input.TotalDigitPointersEncoding

	db, _ = db.Model(&totaldigitDB)
	_, err = db.Updates(&totaldigitDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	totaldigitNew := new(models.TotalDigit)
	totaldigitDB.CopyBasicFieldsToTotalDigit(totaldigitNew)

	// redeem pointers
	totaldigitDB.DecodePointers(backRepo, totaldigitNew)

	// get stage instance from DB instance, and call callback function
	totaldigitOld := backRepo.BackRepoTotalDigit.Map_TotalDigitDBID_TotalDigitPtr[totaldigitDB.ID]
	if totaldigitOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), totaldigitOld, totaldigitNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the totaldigitDB
	c.JSON(http.StatusOK, totaldigitDB)
}

// DeleteTotalDigit
//
// swagger:route DELETE /totaldigits/{ID} totaldigits deleteTotalDigit
//
// # Delete a totaldigit
//
// default: genericError
//
//	200: totaldigitDBResponse
func (controller *Controller) DeleteTotalDigit(c *gin.Context) {

	mutexTotalDigit.Lock()
	defer mutexTotalDigit.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTotalDigit", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTotalDigit.GetDB()

	// Get model if exist
	var totaldigitDB orm.TotalDigitDB
	if _, err := db.First(&totaldigitDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&totaldigitDB)

	// get an instance (not staged) from DB instance, and call callback function
	totaldigitDeleted := new(models.TotalDigit)
	totaldigitDB.CopyBasicFieldsToTotalDigit(totaldigitDeleted)

	// get stage instance from DB instance, and call callback function
	totaldigitStaged := backRepo.BackRepoTotalDigit.Map_TotalDigitDBID_TotalDigitPtr[totaldigitDB.ID]
	if totaldigitStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), totaldigitStaged, totaldigitDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
