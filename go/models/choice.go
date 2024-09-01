package models

type Choice struct {
	Name string
	Annotated

	OccurrenceDefinitionAbstract

	ModelGroup

	ParticleAbstract

	IsDuplicatedInXSD bool

	OuterParticle Particle
}
