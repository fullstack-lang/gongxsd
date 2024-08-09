package models

func PostProcessing(stage *StageStruct) {
	PostProcessingNames(stage)
	PostProcessingAnalyzeXSDStructure(stage)
}
