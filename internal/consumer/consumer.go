package consumer

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/mmfshirokan/nbrb/internal/model"
	log "github.com/sirupsen/logrus"
)

type Adder interface {
	Add(crs []model.Currency) error
}

type Consumer struct {
	sv Adder
}

func New(sv Adder) *Consumer {
	return &Consumer{
		sv: sv,
	}
}

func (c *Consumer) Consume(ctx context.Context, target string) {
	req, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		log.Fatal("Wrong http request: ", err)
	}

	reqAndSave := func() {
		// Sending http request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal("Can't obtain http response, uttempt failed: ", err)
		}

		// Parsing JSON
		crs := make([]model.Currency, 0, 32)
		if err = json.NewDecoder(resp.Body).Decode(&crs); err != nil {
			log.Fatal("Can't decode JSON: ", err)
		}
		resp.Body.Close()

		// Adding currencies
		if err = c.sv.Add(crs); err != nil {
			log.Error("Can't Add in DB: ", err, " another attempt in aprox 24h")
		} else {
			log.Info("Data successfully saved on ", time.Now().UTC())
		}
	}

	reqAndSave()

	tiker := time.NewTicker(tillNextDayUTC())

	for {
		select {
		case <-ctx.Done():
			{
				tiker.Stop()
				return
			}
		case <-tiker.C:
			{
				reqAndSave()
				tiker.Reset(tillNextDayUTC())
			}
		}
	}

}

func tillNextDayUTC() time.Duration {
	return time.Date(
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day()+1,
		0, 0, 1, 0,
		time.UTC,
	).Sub(time.Now().UTC())
}
