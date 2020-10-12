package main

import (
	"fmt"
	"os"

	"github.com/tokane888/learn_statistics/statistics_intro/p66_bootstrap/src/service"
)

func main() {
	logger := NewLogger()
	heightService := service.NewHeightService(logger)
	bootstrapService := service.NewBootStrapService(heightService, logger)

	correlations := bootstrapService.GetCorrelationCoefficients(200)
	outputResult(correlations)
}

func outputResult(correlations []float64) {
	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for _, correlation := range correlations {
		b := []byte(fmt.Sprintf("%f\n", correlation))
		_, err := file.Write(b)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}
