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
var __Length__dummysDeclaration__ models.Length
var __Length_time__dummyDeclaration time.Duration

var mutexLength sync.Mutex

// An LengthID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLength updateLength deleteLength
type LengthID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LengthInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLength updateLength
type LengthInput struct {
	// The Length to submit or modify
	// in: body
	Length *orm.LengthAPI
}

// GetLengths
//
// swagger:route GET /lengths lengths getLengths
//
// # Get all lengths
//
// Responses:
// default: genericError
//
//	200: lengthDBResponse
func (controller *Controller) GetLengths(c *gin.Context) {

	// source slice
	var lengthDBs []orm.LengthDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLengths", "Name", stackPath)
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
	db := backRepo.BackRepoLength.GetDB()

	_, err := db.Find(&lengthDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	lengthAPIs := make([]orm.LengthAPI, 0)

	// for each length, update fields from the database nullable fields
	for idx := range lengthDBs {
		lengthDB := &lengthDBs[idx]
		_ = lengthDB
		var lengthAPI orm.LengthAPI

		// insertion point for updating fields
		lengthAPI.ID = lengthDB.ID
		lengthDB.CopyBasicFieldsToLength_WOP(&lengthAPI.Length_WOP)
		lengthAPI.LengthPointersEncoding = lengthDB.LengthPointersEncoding
		lengthAPIs = append(lengthAPIs, lengthAPI)
	}

	c.JSON(http.StatusOK, lengthAPIs)
}

// PostLength
//
// swagger:route POST /lengths lengths postLength
//
// Creates a length
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLength(c *gin.Context) {

	mutexLength.Lock()
	defer mutexLength.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLengths", "Name", stackPath)
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
	db := backRepo.BackRepoLength.GetDB()

	// Validate input
	var input orm.LengthAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create length
	lengthDB := orm.LengthDB{}
	lengthDB.LengthPointersEncoding = input.LengthPointersEncoding
	lengthDB.CopyBasicFieldsFromLength_WOP(&input.Length_WOP)

	_, err = db.Create(&lengthDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLength.CheckoutPhaseOneInstance(&lengthDB)
	length := backRepo.BackRepoLength.Map_LengthDBID_LengthPtr[lengthDB.ID]

	if length != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), length)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, lengthDB)
}

// GetLength
//
// swagger:route GET /lengths/{ID} lengths getLength
//
// Gets the details for a length.
//
// Responses:
// default: genericError
//
//	200: lengthDBResponse
func (controller *Controller) GetLength(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLength", "Name", stackPath)
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
	db := backRepo.BackRepoLength.GetDB()

	// Get lengthDB in DB
	var lengthDB orm.LengthDB
	if _, err := db.First(&lengthDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var lengthAPI orm.LengthAPI
	lengthAPI.ID = lengthDB.ID
	lengthAPI.LengthPointersEncoding = lengthDB.LengthPointersEncoding
	lengthDB.CopyBasicFieldsToLength_WOP(&lengthAPI.Length_WOP)

	c.JSON(http.StatusOK, lengthAPI)
}

// UpdateLength
//
// swagger:route PATCH /lengths/{ID} lengths updateLength
//
// # Update a length
//
// Responses:
// default: genericError
//
//	200: lengthDBResponse
func (controller *Controller) UpdateLength(c *gin.Context) {

	mutexLength.Lock()
	defer mutexLength.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLength", "Name", stackPath)
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
	db := backRepo.BackRepoLength.GetDB()

	// Validate input
	var input orm.LengthAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var lengthDB orm.LengthDB

	// fetch the length
	_, err := db.First(&lengthDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	lengthDB.CopyBasicFieldsFromLength_WOP(&input.Length_WOP)
	lengthDB.LengthPointersEncoding = input.LengthPointersEncoding

	db, _ = db.Model(&lengthDB)
	_, err = db.Updates(&lengthDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	lengthNew := new(models.Length)
	lengthDB.CopyBasicFieldsToLength(lengthNew)

	// redeem pointers
	lengthDB.DecodePointers(backRepo, lengthNew)

	// get stage instance from DB instance, and call callback function
	lengthOld := backRepo.BackRepoLength.Map_LengthDBID_LengthPtr[lengthDB.ID]
	if lengthOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), lengthOld, lengthNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the lengthDB
	c.JSON(http.StatusOK, lengthDB)
}

// DeleteLength
//
// swagger:route DELETE /lengths/{ID} lengths deleteLength
//
// # Delete a length
//
// default: genericError
//
//	200: lengthDBResponse
func (controller *Controller) DeleteLength(c *gin.Context) {

	mutexLength.Lock()
	defer mutexLength.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLength", "Name", stackPath)
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
	db := backRepo.BackRepoLength.GetDB()

	// Get model if exist
	var lengthDB orm.LengthDB
	if _, err := db.First(&lengthDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&lengthDB)

	// get an instance (not staged) from DB instance, and call callback function
	lengthDeleted := new(models.Length)
	lengthDB.CopyBasicFieldsToLength(lengthDeleted)

	// get stage instance from DB instance, and call callback function
	lengthStaged := backRepo.BackRepoLength.Map_LengthDBID_LengthPtr[lengthDB.ID]
	if lengthStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), lengthStaged, lengthDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
