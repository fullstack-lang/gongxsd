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
var __MinLength__dummysDeclaration__ models.MinLength
var __MinLength_time__dummyDeclaration time.Duration

var mutexMinLength sync.Mutex

// An MinLengthID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMinLength updateMinLength deleteMinLength
type MinLengthID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MinLengthInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMinLength updateMinLength
type MinLengthInput struct {
	// The MinLength to submit or modify
	// in: body
	MinLength *orm.MinLengthAPI
}

// GetMinLengths
//
// swagger:route GET /minlengths minlengths getMinLengths
//
// # Get all minlengths
//
// Responses:
// default: genericError
//
//	200: minlengthDBResponse
func (controller *Controller) GetMinLengths(c *gin.Context) {

	// source slice
	var minlengthDBs []orm.MinLengthDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMinLengths", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gongxsd/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMinLength.GetDB()

	_, err := db.Find(&minlengthDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	minlengthAPIs := make([]orm.MinLengthAPI, 0)

	// for each minlength, update fields from the database nullable fields
	for idx := range minlengthDBs {
		minlengthDB := &minlengthDBs[idx]
		_ = minlengthDB
		var minlengthAPI orm.MinLengthAPI

		// insertion point for updating fields
		minlengthAPI.ID = minlengthDB.ID
		minlengthDB.CopyBasicFieldsToMinLength_WOP(&minlengthAPI.MinLength_WOP)
		minlengthAPI.MinLengthPointersEncoding = minlengthDB.MinLengthPointersEncoding
		minlengthAPIs = append(minlengthAPIs, minlengthAPI)
	}

	c.JSON(http.StatusOK, minlengthAPIs)
}

// PostMinLength
//
// swagger:route POST /minlengths minlengths postMinLength
//
// Creates a minlength
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMinLength(c *gin.Context) {

	mutexMinLength.Lock()
	defer mutexMinLength.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMinLengths", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gongxsd/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMinLength.GetDB()

	// Validate input
	var input orm.MinLengthAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create minlength
	minlengthDB := orm.MinLengthDB{}
	minlengthDB.MinLengthPointersEncoding = input.MinLengthPointersEncoding
	minlengthDB.CopyBasicFieldsFromMinLength_WOP(&input.MinLength_WOP)

	_, err = db.Create(&minlengthDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMinLength.CheckoutPhaseOneInstance(&minlengthDB)
	minlength := backRepo.BackRepoMinLength.Map_MinLengthDBID_MinLengthPtr[minlengthDB.ID]

	if minlength != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), minlength)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, minlengthDB)
}

// GetMinLength
//
// swagger:route GET /minlengths/{ID} minlengths getMinLength
//
// Gets the details for a minlength.
//
// Responses:
// default: genericError
//
//	200: minlengthDBResponse
func (controller *Controller) GetMinLength(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMinLength", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gongxsd/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMinLength.GetDB()

	// Get minlengthDB in DB
	var minlengthDB orm.MinLengthDB
	if _, err := db.First(&minlengthDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var minlengthAPI orm.MinLengthAPI
	minlengthAPI.ID = minlengthDB.ID
	minlengthAPI.MinLengthPointersEncoding = minlengthDB.MinLengthPointersEncoding
	minlengthDB.CopyBasicFieldsToMinLength_WOP(&minlengthAPI.MinLength_WOP)

	c.JSON(http.StatusOK, minlengthAPI)
}

// UpdateMinLength
//
// swagger:route PATCH /minlengths/{ID} minlengths updateMinLength
//
// # Update a minlength
//
// Responses:
// default: genericError
//
//	200: minlengthDBResponse
func (controller *Controller) UpdateMinLength(c *gin.Context) {

	mutexMinLength.Lock()
	defer mutexMinLength.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMinLength", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gongxsd/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMinLength.GetDB()

	// Validate input
	var input orm.MinLengthAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var minlengthDB orm.MinLengthDB

	// fetch the minlength
	_, err := db.First(&minlengthDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	minlengthDB.CopyBasicFieldsFromMinLength_WOP(&input.MinLength_WOP)
	minlengthDB.MinLengthPointersEncoding = input.MinLengthPointersEncoding

	db, _ = db.Model(&minlengthDB)
	_, err = db.Updates(&minlengthDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	minlengthNew := new(models.MinLength)
	minlengthDB.CopyBasicFieldsToMinLength(minlengthNew)

	// redeem pointers
	minlengthDB.DecodePointers(backRepo, minlengthNew)

	// get stage instance from DB instance, and call callback function
	minlengthOld := backRepo.BackRepoMinLength.Map_MinLengthDBID_MinLengthPtr[minlengthDB.ID]
	if minlengthOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), minlengthOld, minlengthNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the minlengthDB
	c.JSON(http.StatusOK, minlengthDB)
}

// DeleteMinLength
//
// swagger:route DELETE /minlengths/{ID} minlengths deleteMinLength
//
// # Delete a minlength
//
// default: genericError
//
//	200: minlengthDBResponse
func (controller *Controller) DeleteMinLength(c *gin.Context) {

	mutexMinLength.Lock()
	defer mutexMinLength.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMinLength", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gongxsd/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMinLength.GetDB()

	// Get model if exist
	var minlengthDB orm.MinLengthDB
	if _, err := db.First(&minlengthDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&minlengthDB)

	// get an instance (not staged) from DB instance, and call callback function
	minlengthDeleted := new(models.MinLength)
	minlengthDB.CopyBasicFieldsToMinLength(minlengthDeleted)

	// get stage instance from DB instance, and call callback function
	minlengthStaged := backRepo.BackRepoMinLength.Map_MinLengthDBID_MinLengthPtr[minlengthDB.ID]
	if minlengthStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), minlengthStaged, minlengthDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
