package engine

import (
	"log"

	"github.com/gomodule/redigo/redis"

	"github.com/Proofsuite/amp-matching-engine/types"
)

// GetOrderBook fetches the complete orderbook from redis for the required pair
func (e *Engine) GetOrderBook(pair *types.Pair) (sellBook, buyBook []*map[string]float64) {
	sKey, bKey := pair.GetOrderBookKeys()
	res, err := redis.Int64s(e.redisConn.Do("SORT", sKey, "GET", sKey+"::book::*", "GET", "#")) // Add price point to order book
	if err != nil {
		log.Print(err)
	}

	for i := 0; i < len(res); i = i + 2 {
		temp := &map[string]float64{
			"volume": float64(res[i]),
			"price":  float64(res[i+1]),
		}
		sellBook = append(sellBook, temp)
	}

	res, err = redis.Int64s(e.redisConn.Do("SORT", bKey, "GET", bKey+"::book::*", "GET", "#", "DESC"))
	if err != nil {
		log.Print(err)
	}

	for i := 0; i < len(res); i = i + 2 {
		temp := &map[string]float64{
			"volume": float64(res[i]),
			"price":  float64(res[i+1]),
		}
		buyBook = append(buyBook, temp)
	}

	return
}
