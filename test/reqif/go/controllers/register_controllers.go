// generated code - do not edit
package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fullstack-lang/gongxsd/test/reqif/go/orm"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// registerControllers register controllers
func registerControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gongxsd/test/reqif/go")
	{ // insertion point for registrations
		v1.GET("/v1/alternative_ids", GetController().GetALTERNATIVE_IDs)
		v1.GET("/v1/alternative_ids/:id", GetController().GetALTERNATIVE_ID)
		v1.POST("/v1/alternative_ids", GetController().PostALTERNATIVE_ID)
		v1.PATCH("/v1/alternative_ids/:id", GetController().UpdateALTERNATIVE_ID)
		v1.PUT("/v1/alternative_ids/:id", GetController().UpdateALTERNATIVE_ID)
		v1.DELETE("/v1/alternative_ids/:id", GetController().DeleteALTERNATIVE_ID)

		v1.GET("/v1/attribute_definition_booleans", GetController().GetATTRIBUTE_DEFINITION_BOOLEANs)
		v1.GET("/v1/attribute_definition_booleans/:id", GetController().GetATTRIBUTE_DEFINITION_BOOLEAN)
		v1.POST("/v1/attribute_definition_booleans", GetController().PostATTRIBUTE_DEFINITION_BOOLEAN)
		v1.PATCH("/v1/attribute_definition_booleans/:id", GetController().UpdateATTRIBUTE_DEFINITION_BOOLEAN)
		v1.PUT("/v1/attribute_definition_booleans/:id", GetController().UpdateATTRIBUTE_DEFINITION_BOOLEAN)
		v1.DELETE("/v1/attribute_definition_booleans/:id", GetController().DeleteATTRIBUTE_DEFINITION_BOOLEAN)

		v1.GET("/v1/attribute_definition_dates", GetController().GetATTRIBUTE_DEFINITION_DATEs)
		v1.GET("/v1/attribute_definition_dates/:id", GetController().GetATTRIBUTE_DEFINITION_DATE)
		v1.POST("/v1/attribute_definition_dates", GetController().PostATTRIBUTE_DEFINITION_DATE)
		v1.PATCH("/v1/attribute_definition_dates/:id", GetController().UpdateATTRIBUTE_DEFINITION_DATE)
		v1.PUT("/v1/attribute_definition_dates/:id", GetController().UpdateATTRIBUTE_DEFINITION_DATE)
		v1.DELETE("/v1/attribute_definition_dates/:id", GetController().DeleteATTRIBUTE_DEFINITION_DATE)

		v1.GET("/v1/attribute_definition_enumerations", GetController().GetATTRIBUTE_DEFINITION_ENUMERATIONs)
		v1.GET("/v1/attribute_definition_enumerations/:id", GetController().GetATTRIBUTE_DEFINITION_ENUMERATION)
		v1.POST("/v1/attribute_definition_enumerations", GetController().PostATTRIBUTE_DEFINITION_ENUMERATION)
		v1.PATCH("/v1/attribute_definition_enumerations/:id", GetController().UpdateATTRIBUTE_DEFINITION_ENUMERATION)
		v1.PUT("/v1/attribute_definition_enumerations/:id", GetController().UpdateATTRIBUTE_DEFINITION_ENUMERATION)
		v1.DELETE("/v1/attribute_definition_enumerations/:id", GetController().DeleteATTRIBUTE_DEFINITION_ENUMERATION)

		v1.GET("/v1/attribute_definition_integers", GetController().GetATTRIBUTE_DEFINITION_INTEGERs)
		v1.GET("/v1/attribute_definition_integers/:id", GetController().GetATTRIBUTE_DEFINITION_INTEGER)
		v1.POST("/v1/attribute_definition_integers", GetController().PostATTRIBUTE_DEFINITION_INTEGER)
		v1.PATCH("/v1/attribute_definition_integers/:id", GetController().UpdateATTRIBUTE_DEFINITION_INTEGER)
		v1.PUT("/v1/attribute_definition_integers/:id", GetController().UpdateATTRIBUTE_DEFINITION_INTEGER)
		v1.DELETE("/v1/attribute_definition_integers/:id", GetController().DeleteATTRIBUTE_DEFINITION_INTEGER)

		v1.GET("/v1/attribute_definition_reals", GetController().GetATTRIBUTE_DEFINITION_REALs)
		v1.GET("/v1/attribute_definition_reals/:id", GetController().GetATTRIBUTE_DEFINITION_REAL)
		v1.POST("/v1/attribute_definition_reals", GetController().PostATTRIBUTE_DEFINITION_REAL)
		v1.PATCH("/v1/attribute_definition_reals/:id", GetController().UpdateATTRIBUTE_DEFINITION_REAL)
		v1.PUT("/v1/attribute_definition_reals/:id", GetController().UpdateATTRIBUTE_DEFINITION_REAL)
		v1.DELETE("/v1/attribute_definition_reals/:id", GetController().DeleteATTRIBUTE_DEFINITION_REAL)

		v1.GET("/v1/attribute_definition_strings", GetController().GetATTRIBUTE_DEFINITION_STRINGs)
		v1.GET("/v1/attribute_definition_strings/:id", GetController().GetATTRIBUTE_DEFINITION_STRING)
		v1.POST("/v1/attribute_definition_strings", GetController().PostATTRIBUTE_DEFINITION_STRING)
		v1.PATCH("/v1/attribute_definition_strings/:id", GetController().UpdateATTRIBUTE_DEFINITION_STRING)
		v1.PUT("/v1/attribute_definition_strings/:id", GetController().UpdateATTRIBUTE_DEFINITION_STRING)
		v1.DELETE("/v1/attribute_definition_strings/:id", GetController().DeleteATTRIBUTE_DEFINITION_STRING)

		v1.GET("/v1/attribute_definition_xhtmls", GetController().GetATTRIBUTE_DEFINITION_XHTMLs)
		v1.GET("/v1/attribute_definition_xhtmls/:id", GetController().GetATTRIBUTE_DEFINITION_XHTML)
		v1.POST("/v1/attribute_definition_xhtmls", GetController().PostATTRIBUTE_DEFINITION_XHTML)
		v1.PATCH("/v1/attribute_definition_xhtmls/:id", GetController().UpdateATTRIBUTE_DEFINITION_XHTML)
		v1.PUT("/v1/attribute_definition_xhtmls/:id", GetController().UpdateATTRIBUTE_DEFINITION_XHTML)
		v1.DELETE("/v1/attribute_definition_xhtmls/:id", GetController().DeleteATTRIBUTE_DEFINITION_XHTML)

		v1.GET("/v1/attribute_value_booleans", GetController().GetATTRIBUTE_VALUE_BOOLEANs)
		v1.GET("/v1/attribute_value_booleans/:id", GetController().GetATTRIBUTE_VALUE_BOOLEAN)
		v1.POST("/v1/attribute_value_booleans", GetController().PostATTRIBUTE_VALUE_BOOLEAN)
		v1.PATCH("/v1/attribute_value_booleans/:id", GetController().UpdateATTRIBUTE_VALUE_BOOLEAN)
		v1.PUT("/v1/attribute_value_booleans/:id", GetController().UpdateATTRIBUTE_VALUE_BOOLEAN)
		v1.DELETE("/v1/attribute_value_booleans/:id", GetController().DeleteATTRIBUTE_VALUE_BOOLEAN)

		v1.GET("/v1/attribute_value_dates", GetController().GetATTRIBUTE_VALUE_DATEs)
		v1.GET("/v1/attribute_value_dates/:id", GetController().GetATTRIBUTE_VALUE_DATE)
		v1.POST("/v1/attribute_value_dates", GetController().PostATTRIBUTE_VALUE_DATE)
		v1.PATCH("/v1/attribute_value_dates/:id", GetController().UpdateATTRIBUTE_VALUE_DATE)
		v1.PUT("/v1/attribute_value_dates/:id", GetController().UpdateATTRIBUTE_VALUE_DATE)
		v1.DELETE("/v1/attribute_value_dates/:id", GetController().DeleteATTRIBUTE_VALUE_DATE)

		v1.GET("/v1/attribute_value_enumerations", GetController().GetATTRIBUTE_VALUE_ENUMERATIONs)
		v1.GET("/v1/attribute_value_enumerations/:id", GetController().GetATTRIBUTE_VALUE_ENUMERATION)
		v1.POST("/v1/attribute_value_enumerations", GetController().PostATTRIBUTE_VALUE_ENUMERATION)
		v1.PATCH("/v1/attribute_value_enumerations/:id", GetController().UpdateATTRIBUTE_VALUE_ENUMERATION)
		v1.PUT("/v1/attribute_value_enumerations/:id", GetController().UpdateATTRIBUTE_VALUE_ENUMERATION)
		v1.DELETE("/v1/attribute_value_enumerations/:id", GetController().DeleteATTRIBUTE_VALUE_ENUMERATION)

		v1.GET("/v1/attribute_value_integers", GetController().GetATTRIBUTE_VALUE_INTEGERs)
		v1.GET("/v1/attribute_value_integers/:id", GetController().GetATTRIBUTE_VALUE_INTEGER)
		v1.POST("/v1/attribute_value_integers", GetController().PostATTRIBUTE_VALUE_INTEGER)
		v1.PATCH("/v1/attribute_value_integers/:id", GetController().UpdateATTRIBUTE_VALUE_INTEGER)
		v1.PUT("/v1/attribute_value_integers/:id", GetController().UpdateATTRIBUTE_VALUE_INTEGER)
		v1.DELETE("/v1/attribute_value_integers/:id", GetController().DeleteATTRIBUTE_VALUE_INTEGER)

		v1.GET("/v1/attribute_value_reals", GetController().GetATTRIBUTE_VALUE_REALs)
		v1.GET("/v1/attribute_value_reals/:id", GetController().GetATTRIBUTE_VALUE_REAL)
		v1.POST("/v1/attribute_value_reals", GetController().PostATTRIBUTE_VALUE_REAL)
		v1.PATCH("/v1/attribute_value_reals/:id", GetController().UpdateATTRIBUTE_VALUE_REAL)
		v1.PUT("/v1/attribute_value_reals/:id", GetController().UpdateATTRIBUTE_VALUE_REAL)
		v1.DELETE("/v1/attribute_value_reals/:id", GetController().DeleteATTRIBUTE_VALUE_REAL)

		v1.GET("/v1/attribute_value_strings", GetController().GetATTRIBUTE_VALUE_STRINGs)
		v1.GET("/v1/attribute_value_strings/:id", GetController().GetATTRIBUTE_VALUE_STRING)
		v1.POST("/v1/attribute_value_strings", GetController().PostATTRIBUTE_VALUE_STRING)
		v1.PATCH("/v1/attribute_value_strings/:id", GetController().UpdateATTRIBUTE_VALUE_STRING)
		v1.PUT("/v1/attribute_value_strings/:id", GetController().UpdateATTRIBUTE_VALUE_STRING)
		v1.DELETE("/v1/attribute_value_strings/:id", GetController().DeleteATTRIBUTE_VALUE_STRING)

		v1.GET("/v1/attribute_value_xhtmls", GetController().GetATTRIBUTE_VALUE_XHTMLs)
		v1.GET("/v1/attribute_value_xhtmls/:id", GetController().GetATTRIBUTE_VALUE_XHTML)
		v1.POST("/v1/attribute_value_xhtmls", GetController().PostATTRIBUTE_VALUE_XHTML)
		v1.PATCH("/v1/attribute_value_xhtmls/:id", GetController().UpdateATTRIBUTE_VALUE_XHTML)
		v1.PUT("/v1/attribute_value_xhtmls/:id", GetController().UpdateATTRIBUTE_VALUE_XHTML)
		v1.DELETE("/v1/attribute_value_xhtmls/:id", GetController().DeleteATTRIBUTE_VALUE_XHTML)

		v1.GET("/v1/datatype_definition_booleans", GetController().GetDATATYPE_DEFINITION_BOOLEANs)
		v1.GET("/v1/datatype_definition_booleans/:id", GetController().GetDATATYPE_DEFINITION_BOOLEAN)
		v1.POST("/v1/datatype_definition_booleans", GetController().PostDATATYPE_DEFINITION_BOOLEAN)
		v1.PATCH("/v1/datatype_definition_booleans/:id", GetController().UpdateDATATYPE_DEFINITION_BOOLEAN)
		v1.PUT("/v1/datatype_definition_booleans/:id", GetController().UpdateDATATYPE_DEFINITION_BOOLEAN)
		v1.DELETE("/v1/datatype_definition_booleans/:id", GetController().DeleteDATATYPE_DEFINITION_BOOLEAN)

		v1.GET("/v1/datatype_definition_dates", GetController().GetDATATYPE_DEFINITION_DATEs)
		v1.GET("/v1/datatype_definition_dates/:id", GetController().GetDATATYPE_DEFINITION_DATE)
		v1.POST("/v1/datatype_definition_dates", GetController().PostDATATYPE_DEFINITION_DATE)
		v1.PATCH("/v1/datatype_definition_dates/:id", GetController().UpdateDATATYPE_DEFINITION_DATE)
		v1.PUT("/v1/datatype_definition_dates/:id", GetController().UpdateDATATYPE_DEFINITION_DATE)
		v1.DELETE("/v1/datatype_definition_dates/:id", GetController().DeleteDATATYPE_DEFINITION_DATE)

		v1.GET("/v1/datatype_definition_enumerations", GetController().GetDATATYPE_DEFINITION_ENUMERATIONs)
		v1.GET("/v1/datatype_definition_enumerations/:id", GetController().GetDATATYPE_DEFINITION_ENUMERATION)
		v1.POST("/v1/datatype_definition_enumerations", GetController().PostDATATYPE_DEFINITION_ENUMERATION)
		v1.PATCH("/v1/datatype_definition_enumerations/:id", GetController().UpdateDATATYPE_DEFINITION_ENUMERATION)
		v1.PUT("/v1/datatype_definition_enumerations/:id", GetController().UpdateDATATYPE_DEFINITION_ENUMERATION)
		v1.DELETE("/v1/datatype_definition_enumerations/:id", GetController().DeleteDATATYPE_DEFINITION_ENUMERATION)

		v1.GET("/v1/datatype_definition_integers", GetController().GetDATATYPE_DEFINITION_INTEGERs)
		v1.GET("/v1/datatype_definition_integers/:id", GetController().GetDATATYPE_DEFINITION_INTEGER)
		v1.POST("/v1/datatype_definition_integers", GetController().PostDATATYPE_DEFINITION_INTEGER)
		v1.PATCH("/v1/datatype_definition_integers/:id", GetController().UpdateDATATYPE_DEFINITION_INTEGER)
		v1.PUT("/v1/datatype_definition_integers/:id", GetController().UpdateDATATYPE_DEFINITION_INTEGER)
		v1.DELETE("/v1/datatype_definition_integers/:id", GetController().DeleteDATATYPE_DEFINITION_INTEGER)

		v1.GET("/v1/datatype_definition_reals", GetController().GetDATATYPE_DEFINITION_REALs)
		v1.GET("/v1/datatype_definition_reals/:id", GetController().GetDATATYPE_DEFINITION_REAL)
		v1.POST("/v1/datatype_definition_reals", GetController().PostDATATYPE_DEFINITION_REAL)
		v1.PATCH("/v1/datatype_definition_reals/:id", GetController().UpdateDATATYPE_DEFINITION_REAL)
		v1.PUT("/v1/datatype_definition_reals/:id", GetController().UpdateDATATYPE_DEFINITION_REAL)
		v1.DELETE("/v1/datatype_definition_reals/:id", GetController().DeleteDATATYPE_DEFINITION_REAL)

		v1.GET("/v1/datatype_definition_strings", GetController().GetDATATYPE_DEFINITION_STRINGs)
		v1.GET("/v1/datatype_definition_strings/:id", GetController().GetDATATYPE_DEFINITION_STRING)
		v1.POST("/v1/datatype_definition_strings", GetController().PostDATATYPE_DEFINITION_STRING)
		v1.PATCH("/v1/datatype_definition_strings/:id", GetController().UpdateDATATYPE_DEFINITION_STRING)
		v1.PUT("/v1/datatype_definition_strings/:id", GetController().UpdateDATATYPE_DEFINITION_STRING)
		v1.DELETE("/v1/datatype_definition_strings/:id", GetController().DeleteDATATYPE_DEFINITION_STRING)

		v1.GET("/v1/datatype_definition_xhtmls", GetController().GetDATATYPE_DEFINITION_XHTMLs)
		v1.GET("/v1/datatype_definition_xhtmls/:id", GetController().GetDATATYPE_DEFINITION_XHTML)
		v1.POST("/v1/datatype_definition_xhtmls", GetController().PostDATATYPE_DEFINITION_XHTML)
		v1.PATCH("/v1/datatype_definition_xhtmls/:id", GetController().UpdateDATATYPE_DEFINITION_XHTML)
		v1.PUT("/v1/datatype_definition_xhtmls/:id", GetController().UpdateDATATYPE_DEFINITION_XHTML)
		v1.DELETE("/v1/datatype_definition_xhtmls/:id", GetController().DeleteDATATYPE_DEFINITION_XHTML)

		v1.GET("/v1/embedded_values", GetController().GetEMBEDDED_VALUEs)
		v1.GET("/v1/embedded_values/:id", GetController().GetEMBEDDED_VALUE)
		v1.POST("/v1/embedded_values", GetController().PostEMBEDDED_VALUE)
		v1.PATCH("/v1/embedded_values/:id", GetController().UpdateEMBEDDED_VALUE)
		v1.PUT("/v1/embedded_values/:id", GetController().UpdateEMBEDDED_VALUE)
		v1.DELETE("/v1/embedded_values/:id", GetController().DeleteEMBEDDED_VALUE)

		v1.GET("/v1/enum_values", GetController().GetENUM_VALUEs)
		v1.GET("/v1/enum_values/:id", GetController().GetENUM_VALUE)
		v1.POST("/v1/enum_values", GetController().PostENUM_VALUE)
		v1.PATCH("/v1/enum_values/:id", GetController().UpdateENUM_VALUE)
		v1.PUT("/v1/enum_values/:id", GetController().UpdateENUM_VALUE)
		v1.DELETE("/v1/enum_values/:id", GetController().DeleteENUM_VALUE)

		v1.GET("/v1/relation_groups", GetController().GetRELATION_GROUPs)
		v1.GET("/v1/relation_groups/:id", GetController().GetRELATION_GROUP)
		v1.POST("/v1/relation_groups", GetController().PostRELATION_GROUP)
		v1.PATCH("/v1/relation_groups/:id", GetController().UpdateRELATION_GROUP)
		v1.PUT("/v1/relation_groups/:id", GetController().UpdateRELATION_GROUP)
		v1.DELETE("/v1/relation_groups/:id", GetController().DeleteRELATION_GROUP)

		v1.GET("/v1/relation_group_types", GetController().GetRELATION_GROUP_TYPEs)
		v1.GET("/v1/relation_group_types/:id", GetController().GetRELATION_GROUP_TYPE)
		v1.POST("/v1/relation_group_types", GetController().PostRELATION_GROUP_TYPE)
		v1.PATCH("/v1/relation_group_types/:id", GetController().UpdateRELATION_GROUP_TYPE)
		v1.PUT("/v1/relation_group_types/:id", GetController().UpdateRELATION_GROUP_TYPE)
		v1.DELETE("/v1/relation_group_types/:id", GetController().DeleteRELATION_GROUP_TYPE)

		v1.GET("/v1/req_ifs", GetController().GetREQ_IFs)
		v1.GET("/v1/req_ifs/:id", GetController().GetREQ_IF)
		v1.POST("/v1/req_ifs", GetController().PostREQ_IF)
		v1.PATCH("/v1/req_ifs/:id", GetController().UpdateREQ_IF)
		v1.PUT("/v1/req_ifs/:id", GetController().UpdateREQ_IF)
		v1.DELETE("/v1/req_ifs/:id", GetController().DeleteREQ_IF)

		v1.GET("/v1/req_if_contents", GetController().GetREQ_IF_CONTENTs)
		v1.GET("/v1/req_if_contents/:id", GetController().GetREQ_IF_CONTENT)
		v1.POST("/v1/req_if_contents", GetController().PostREQ_IF_CONTENT)
		v1.PATCH("/v1/req_if_contents/:id", GetController().UpdateREQ_IF_CONTENT)
		v1.PUT("/v1/req_if_contents/:id", GetController().UpdateREQ_IF_CONTENT)
		v1.DELETE("/v1/req_if_contents/:id", GetController().DeleteREQ_IF_CONTENT)

		v1.GET("/v1/req_if_headers", GetController().GetREQ_IF_HEADERs)
		v1.GET("/v1/req_if_headers/:id", GetController().GetREQ_IF_HEADER)
		v1.POST("/v1/req_if_headers", GetController().PostREQ_IF_HEADER)
		v1.PATCH("/v1/req_if_headers/:id", GetController().UpdateREQ_IF_HEADER)
		v1.PUT("/v1/req_if_headers/:id", GetController().UpdateREQ_IF_HEADER)
		v1.DELETE("/v1/req_if_headers/:id", GetController().DeleteREQ_IF_HEADER)

		v1.GET("/v1/req_if_tool_extensions", GetController().GetREQ_IF_TOOL_EXTENSIONs)
		v1.GET("/v1/req_if_tool_extensions/:id", GetController().GetREQ_IF_TOOL_EXTENSION)
		v1.POST("/v1/req_if_tool_extensions", GetController().PostREQ_IF_TOOL_EXTENSION)
		v1.PATCH("/v1/req_if_tool_extensions/:id", GetController().UpdateREQ_IF_TOOL_EXTENSION)
		v1.PUT("/v1/req_if_tool_extensions/:id", GetController().UpdateREQ_IF_TOOL_EXTENSION)
		v1.DELETE("/v1/req_if_tool_extensions/:id", GetController().DeleteREQ_IF_TOOL_EXTENSION)

		v1.GET("/v1/specifications", GetController().GetSPECIFICATIONs)
		v1.GET("/v1/specifications/:id", GetController().GetSPECIFICATION)
		v1.POST("/v1/specifications", GetController().PostSPECIFICATION)
		v1.PATCH("/v1/specifications/:id", GetController().UpdateSPECIFICATION)
		v1.PUT("/v1/specifications/:id", GetController().UpdateSPECIFICATION)
		v1.DELETE("/v1/specifications/:id", GetController().DeleteSPECIFICATION)

		v1.GET("/v1/specification_types", GetController().GetSPECIFICATION_TYPEs)
		v1.GET("/v1/specification_types/:id", GetController().GetSPECIFICATION_TYPE)
		v1.POST("/v1/specification_types", GetController().PostSPECIFICATION_TYPE)
		v1.PATCH("/v1/specification_types/:id", GetController().UpdateSPECIFICATION_TYPE)
		v1.PUT("/v1/specification_types/:id", GetController().UpdateSPECIFICATION_TYPE)
		v1.DELETE("/v1/specification_types/:id", GetController().DeleteSPECIFICATION_TYPE)

		v1.GET("/v1/spec_hierarchys", GetController().GetSPEC_HIERARCHYs)
		v1.GET("/v1/spec_hierarchys/:id", GetController().GetSPEC_HIERARCHY)
		v1.POST("/v1/spec_hierarchys", GetController().PostSPEC_HIERARCHY)
		v1.PATCH("/v1/spec_hierarchys/:id", GetController().UpdateSPEC_HIERARCHY)
		v1.PUT("/v1/spec_hierarchys/:id", GetController().UpdateSPEC_HIERARCHY)
		v1.DELETE("/v1/spec_hierarchys/:id", GetController().DeleteSPEC_HIERARCHY)

		v1.GET("/v1/spec_objects", GetController().GetSPEC_OBJECTs)
		v1.GET("/v1/spec_objects/:id", GetController().GetSPEC_OBJECT)
		v1.POST("/v1/spec_objects", GetController().PostSPEC_OBJECT)
		v1.PATCH("/v1/spec_objects/:id", GetController().UpdateSPEC_OBJECT)
		v1.PUT("/v1/spec_objects/:id", GetController().UpdateSPEC_OBJECT)
		v1.DELETE("/v1/spec_objects/:id", GetController().DeleteSPEC_OBJECT)

		v1.GET("/v1/spec_object_types", GetController().GetSPEC_OBJECT_TYPEs)
		v1.GET("/v1/spec_object_types/:id", GetController().GetSPEC_OBJECT_TYPE)
		v1.POST("/v1/spec_object_types", GetController().PostSPEC_OBJECT_TYPE)
		v1.PATCH("/v1/spec_object_types/:id", GetController().UpdateSPEC_OBJECT_TYPE)
		v1.PUT("/v1/spec_object_types/:id", GetController().UpdateSPEC_OBJECT_TYPE)
		v1.DELETE("/v1/spec_object_types/:id", GetController().DeleteSPEC_OBJECT_TYPE)

		v1.GET("/v1/spec_relations", GetController().GetSPEC_RELATIONs)
		v1.GET("/v1/spec_relations/:id", GetController().GetSPEC_RELATION)
		v1.POST("/v1/spec_relations", GetController().PostSPEC_RELATION)
		v1.PATCH("/v1/spec_relations/:id", GetController().UpdateSPEC_RELATION)
		v1.PUT("/v1/spec_relations/:id", GetController().UpdateSPEC_RELATION)
		v1.DELETE("/v1/spec_relations/:id", GetController().DeleteSPEC_RELATION)

		v1.GET("/v1/spec_relation_types", GetController().GetSPEC_RELATION_TYPEs)
		v1.GET("/v1/spec_relation_types/:id", GetController().GetSPEC_RELATION_TYPE)
		v1.POST("/v1/spec_relation_types", GetController().PostSPEC_RELATION_TYPE)
		v1.PATCH("/v1/spec_relation_types/:id", GetController().UpdateSPEC_RELATION_TYPE)
		v1.PUT("/v1/spec_relation_types/:id", GetController().UpdateSPEC_RELATION_TYPE)
		v1.DELETE("/v1/spec_relation_types/:id", GetController().DeleteSPEC_RELATION_TYPE)

		v1.GET("/v1/xhtml_contents", GetController().GetXHTML_CONTENTs)
		v1.GET("/v1/xhtml_contents/:id", GetController().GetXHTML_CONTENT)
		v1.POST("/v1/xhtml_contents", GetController().PostXHTML_CONTENT)
		v1.PATCH("/v1/xhtml_contents/:id", GetController().UpdateXHTML_CONTENT)
		v1.PUT("/v1/xhtml_contents/:id", GetController().UpdateXHTML_CONTENT)
		v1.DELETE("/v1/xhtml_contents/:id", GetController().DeleteXHTML_CONTENT)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)

		v1.GET("/v1/ws/commitfrombacknb", GetController().onWebSocketRequestForCommitFromBackNb)
		v1.GET("/v1/ws/stage", GetController().onWebSocketRequestForBackRepoContent)

		v1.GET("/v1/stacks", GetController().stacks)
	}
}

func (controller *Controller) stacks(c *gin.Context) {

	var res []string

	for k, _ := range controller.Map_BackRepos {
		res = append(res, k)
	}

	c.JSON(http.StatusOK, res)
}

// onWebSocketRequestForCommitFromBackNb is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequestForCommitFromBackNb(c *gin.Context) {

	// log.Println("Stack github.com/fullstack-lang/gongxsd/test/reqif/go, onWebSocketRequestForCommitFromBackNb")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:8080" || origin == "http://localhost:4200"
		},
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go, Unkown stack", stackPath)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb()

	for nbCommitBackRepo := range updateCommitBackRepoNbChannel {

		// Send elapsed time as a string over the WebSocket connection
		err = wsConnection.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%d", nbCommitBackRepo)))
		if err != nil {
			log.Println("client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
			fmt.Println(err)
			return
		}
	}
}

// onWebSocketRequestForBackRepoContent is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequestForBackRepoContent(c *gin.Context) {

	// log.Println("Stack github.com/fullstack-lang/gongxsd/test/reqif/go, onWebSocketRequestForBackRepoContent")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:8080" || origin == "http://localhost:4200"
		},
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go, Unkown stack", stackPath)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb()

	backRepoData := new(orm.BackRepoData)
	orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)

	err = wsConnection.WriteJSON(backRepoData)
	// log.Println("Stack github.com/fullstack-lang/gongxsd/test/reqif/go, onWebSocketRequestForBackRepoContent, first sent back repo of", stackPath)
	if err != nil {
		log.Println("client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	}

	for nbCommitBackRepo := range updateCommitBackRepoNbChannel {
		_ = nbCommitBackRepo

		backRepoData := new(orm.BackRepoData)
		orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)

		// Send backRepo data
		err = wsConnection.WriteJSON(backRepoData)

		// log.Println("Stack github.com/fullstack-lang/gongxsd/test/reqif/go, onWebSocketRequestForBackRepoContent, sent back repo of", stackPath)

		if err != nil {
			log.Println("client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
			fmt.Println(err)
			return
		}
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastPushFromFrontNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongxsd/test/reqif/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
