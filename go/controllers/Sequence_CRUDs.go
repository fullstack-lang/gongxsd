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
var __Sequence__dummysDeclaration__ models.Sequence
var __Sequence_time__dummyDeclaration time.Duration

var mutexSequence sync.Mutex

// An SequenceID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSequence updateSequence deleteSequence
type SequenceID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SequenceInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSequence updateSequence
type SequenceInput struct {
	// The Sequence to submit or modify
	// in: body
	Sequence *orm.SequenceAPI
}

// GetSequences
//
// swagger:route GET /sequences sequences getSequences
//
// # Get all sequences
//
// Responses:
// default: genericError
//
//	200: sequenceDBResponse
func (controller *Controller) GetSequences(c *gin.Context) {

	// source slice
	var sequenceDBs []orm.SequenceDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSequences", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSequence.GetDB()

	_, err := db.Find(&sequenceDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	sequenceAPIs := make([]orm.SequenceAPI, 0)

	// for each sequence, update fields from the database nullable fields
	for idx := range sequenceDBs {
		sequenceDB := &sequenceDBs[idx]
		_ = sequenceDB
		var sequenceAPI orm.SequenceAPI

		// insertion point for updating fields
		sequenceAPI.ID = sequenceDB.ID
		sequenceDB.CopyBasicFieldsToSequence_WOP(&sequenceAPI.Sequence_WOP)
		sequenceAPI.SequencePointersEncoding = sequenceDB.SequencePointersEncoding
		sequenceAPIs = append(sequenceAPIs, sequenceAPI)
	}

	c.JSON(http.StatusOK, sequenceAPIs)
}

// PostSequence
//
// swagger:route POST /sequences sequences postSequence
//
// Creates a sequence
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSequence(c *gin.Context) {

	mutexSequence.Lock()
	defer mutexSequence.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSequences", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSequence.GetDB()

	// Validate input
	var input orm.SequenceAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create sequence
	sequenceDB := orm.SequenceDB{}
	sequenceDB.SequencePointersEncoding = input.SequencePointersEncoding
	sequenceDB.CopyBasicFieldsFromSequence_WOP(&input.Sequence_WOP)

	_, err = db.Create(&sequenceDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSequence.CheckoutPhaseOneInstance(&sequenceDB)
	sequence := backRepo.BackRepoSequence.Map_SequenceDBID_SequencePtr[sequenceDB.ID]

	if sequence != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), sequence)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, sequenceDB)
}

// GetSequence
//
// swagger:route GET /sequences/{ID} sequences getSequence
//
// Gets the details for a sequence.
//
// Responses:
// default: genericError
//
//	200: sequenceDBResponse
func (controller *Controller) GetSequence(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSequence", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSequence.GetDB()

	// Get sequenceDB in DB
	var sequenceDB orm.SequenceDB
	if _, err := db.First(&sequenceDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var sequenceAPI orm.SequenceAPI
	sequenceAPI.ID = sequenceDB.ID
	sequenceAPI.SequencePointersEncoding = sequenceDB.SequencePointersEncoding
	sequenceDB.CopyBasicFieldsToSequence_WOP(&sequenceAPI.Sequence_WOP)

	c.JSON(http.StatusOK, sequenceAPI)
}

// UpdateSequence
//
// swagger:route PATCH /sequences/{ID} sequences updateSequence
//
// # Update a sequence
//
// Responses:
// default: genericError
//
//	200: sequenceDBResponse
func (controller *Controller) UpdateSequence(c *gin.Context) {

	mutexSequence.Lock()
	defer mutexSequence.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSequence", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSequence.GetDB()

	// Validate input
	var input orm.SequenceAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var sequenceDB orm.SequenceDB

	// fetch the sequence
	_, err := db.First(&sequenceDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	sequenceDB.CopyBasicFieldsFromSequence_WOP(&input.Sequence_WOP)
	sequenceDB.SequencePointersEncoding = input.SequencePointersEncoding

	db, _ = db.Model(&sequenceDB)
	_, err = db.Updates(sequenceDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	sequenceNew := new(models.Sequence)
	sequenceDB.CopyBasicFieldsToSequence(sequenceNew)

	// redeem pointers
	sequenceDB.DecodePointers(backRepo, sequenceNew)

	// get stage instance from DB instance, and call callback function
	sequenceOld := backRepo.BackRepoSequence.Map_SequenceDBID_SequencePtr[sequenceDB.ID]
	if sequenceOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), sequenceOld, sequenceNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the sequenceDB
	c.JSON(http.StatusOK, sequenceDB)
}

// DeleteSequence
//
// swagger:route DELETE /sequences/{ID} sequences deleteSequence
//
// # Delete a sequence
//
// default: genericError
//
//	200: sequenceDBResponse
func (controller *Controller) DeleteSequence(c *gin.Context) {

	mutexSequence.Lock()
	defer mutexSequence.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSequence", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSequence.GetDB()

	// Get model if exist
	var sequenceDB orm.SequenceDB
	if _, err := db.First(&sequenceDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&sequenceDB)

	// get an instance (not staged) from DB instance, and call callback function
	sequenceDeleted := new(models.Sequence)
	sequenceDB.CopyBasicFieldsToSequence(sequenceDeleted)

	// get stage instance from DB instance, and call callback function
	sequenceStaged := backRepo.BackRepoSequence.Map_SequenceDBID_SequencePtr[sequenceDB.ID]
	if sequenceStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), sequenceStaged, sequenceDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
