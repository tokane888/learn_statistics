package service

import (
	"fmt"

	"go.uber.org/zap"
	"gonum.org/v1/gonum/stat"
)

type BootStrapService struct {
	heightService *HeightService
	logger        *zap.Logger
}

func NewBootStrapService(heightService *HeightService, logger *zap.Logger) *BootStrapService {
	s := &BootStrapService{
		heightService: heightService,
		logger:        logger,
	}

	return s
}

// GetCorrelationCoefficients generate data and caliculate correlation coefficient
// indicated times. Return results at float64 array.
func (s *BootStrapService) GetCorrelationCoefficients(num int) []float64 {
	var results []float64
	for i := 0; i < num; i++ {
		results = append(results, s.GetCorrelationCoefficient())
	}

	return results
}

// GetCorrelationCoefficient get same number of data in HeightService.heights randomly.
// After that, calculate correlation coefficient of the data and return it.
func (s *BootStrapService) GetCorrelationCoefficient() float64 {
	var manHeights []float64
	var womanHeights []float64
	for i := 0; i < s.heightService.GetDataCount(); i++ {
		height := s.heightService.GetRandomHeight()
		manHeights = append(manHeights, (float64)(height.Man))
		womanHeights = append(womanHeights, (float64)(height.Woman))
	}

	s.logger.Debug("manHeights")
	s.logger.Debug(fmt.Sprintf("%v\n", manHeights))
	s.logger.Debug("womanHeights")
	s.logger.Debug(fmt.Sprintf("%v\n", womanHeights))

	return stat.Correlation(manHeights, womanHeights, nil)
}
