package service

import (
	"math/rand"
	"time"

	"go.uber.org/zap"
)

type Height struct {
	Man   int
	Woman int
}

type HeightService struct {
	heights []Height
	logger  *zap.Logger
}

func NewHeightService(logger *zap.Logger) *HeightService {
	s := &HeightService{
		heights: []Height{
			{Man: 71, Woman: 69},
			{Man: 68, Woman: 64},
			{Man: 66, Woman: 65},
			{Man: 67, Woman: 63},
			{Man: 70, Woman: 65},
			{Man: 71, Woman: 62},
			{Man: 70, Woman: 65},
			{Man: 73, Woman: 64},
			{Man: 72, Woman: 66},
			{Man: 65, Woman: 59},
			{Man: 66, Woman: 62},
		},
	}
	s.logger = logger

	return s
}

// GetRandomHeight returns a single Height pair randomly.
func (h *HeightService) GetRandomHeight() Height {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	return h.heights[random.Intn(len(h.heights))]
}

func (h *HeightService) GetDataCount() int {
	return len(h.heights)
}
