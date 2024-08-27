package models

func PostProcessing(stage *StageStruct) {
	PostProcessingComputeAnonymousity(stage)
	PostProcessingNames(stage)
	PostProcessingAnalyzeXSDStructure(stage)
}
