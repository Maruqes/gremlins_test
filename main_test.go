package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncomendaPequena(t *testing.T) {
	assert.Equal(t, 3500, totalEncomenda(3000, false))
}

func TestEncomendaVIP(t *testing.T) {
	assert.Equal(t, 10800, totalEncomenda(12000, true))
}

func TestEncomendaSemVIP(t *testing.T) {
	assert.Equal(t, 12000, totalEncomenda(12000, false))
}

// func TestLimiteDescontoVIP(t *testing.T) {
// 	assert.Equal(t, 9000, totalEncomenda(10000, true))
// }

// func TestLimitePortesGratis(t *testing.T) {
// 	assert.Equal(t, 5000, totalEncomenda(5000, false))
// }
