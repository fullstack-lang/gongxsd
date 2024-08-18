package models

type Particle interface {
	GetOrder() int
}

var _ Particle = (*Element)(nil)
var _ Particle = (*Group)(nil)

type ParticleAbstract struct {
	// Order is the order at wich the particle was unmarshalled in the xsd
	// It is important to preserve the order output that is defined in the xsd
	Order int `xml:"-"`
	Depth int `xml:"-"`
}

// GetOrder implements Particle.
func (p *ParticleAbstract) GetOrder() int {
	return p.Order
}
