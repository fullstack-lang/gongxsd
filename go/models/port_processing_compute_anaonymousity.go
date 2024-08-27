package models

func PostProcessingComputeAnonymousity(stage *StageStruct) {

	for ct := range *GetGongstructInstancesSet[ComplexType](stage) {
		ct.IsAnonymous = true
	}
	for _, ct := range SchemaSingloton.ComplexTypes {
		ct.IsAnonymous = false
	}
}
