package calc

import (
	"github.com/cmsd2/aws_price_calc/types"
	"math"
)

func SqsPriceRps(config *types.Sqs, requests_per_second float64, is_fifo bool, message_size_kb float64) float64 {
	requests := requests_per_second * seconds_per_month()

	return SqsPrice(config, requests, is_fifo, message_size_kb)
}

func SqsPrice(config *types.Sqs, requests float64, is_fifo bool, message_size_kb float64) float64 {
	var price float64

	if is_fifo {
		price = config.Price.Requests.Fifo
	} else {
		price = config.Price.Requests.Standard
	}

	requests_price := requests * price / config.Price.Requests.Per

	data_size_gb := requests * message_size_kb / 1024.0 / 1024.0

	data_transfer_price := LadderPrice(&config.Price.Data.Out, data_size_gb)

	return requests_price + data_transfer_price
}

func LadderPrice(ladder *types.DataLadderPrice, quantity float64) float64 {
	price := 0.0

	for i := range ladder.Bands {
		band := ladder.Bands[i]

		if quantity >= band.From && band.Poa {
			return math.NaN()
		}

		band_quantity := math.Max(0.0, math.Min(quantity, band.To) - band.From)

		price += band_quantity * band.Price
	}

	return price
}

func seconds_per_month() float64 {
	return 60 * 60 * 24 * 31
}