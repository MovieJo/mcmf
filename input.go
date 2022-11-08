package mcmf

import (
	"fmt"
)

type Input struct {
	Ids      []int     `json:"ids"`
	Costs    [][]int64 `json:"costs"`
	Problems []Problem `json:"problems"`
}

type Problem struct {
	Sources []Tip  `json:"sources"`
	Links   []Link `json:"links"`
	Sinks   []Tip  `json:"sinks"`
}

type Tip struct {
	Id       int   `json:"id"`
	Capacity int64 `json:"capacity"`
}

type Link struct {
	FromId   int   `json:"from"`
	ToId     int   `json:"to"`
	Capacity int64 `json:"capacity"`
}

func (i Input) Validate() error {
	exists := map[int]bool{}

	// id validation
	for _, id := range i.Ids {
		if _, e := exists[id]; !e {
			return fmt.Errorf("duplicated Id %d", id)
		}
		exists[id] = true
	}

	// matrix size validation
	matrixSize := len(i.Ids)
	if len(i.Costs) != matrixSize {
		return fmt.Errorf(
			"invalid cost matrix size: ids len=%d, rows len=%d",
			matrixSize, len(i.Costs),
		)
	}
	for idx, row := range i.Costs {
		if len(row) != matrixSize {
			return fmt.Errorf(
				"invalid cost matrix size: ids len=%d, column %d len=%d",
				matrixSize, idx, len(row),
			)
		}
	}

	// problem validation
	for _, p := range i.Problems {
		if err := p.Validate(exists); err != nil {
			return err
		}
	}

	// all check passed
	return nil
}

func (p Problem) Validate(existIds map[int]bool) error {
	// sources id validation
	for idx, tip := range p.Sources {
		if _, e := existIds[tip.Id]; !e {
			return fmt.Errorf(
				"id for sources %d not exists: %d",
				idx, tip.Id,
			)
		}
	}

	// links id validation
	for idx, link := range p.Links {
		if _, e := existIds[link.FromId]; !e {
			return fmt.Errorf(
				"from id for link %d not exists: %d",
				idx, link.FromId,
			)
		}
		if _, e := existIds[link.ToId]; !e {
			return fmt.Errorf(
				"to id for link %d not exists: %d",
				idx, link.ToId,
			)
		}
	}

	// sinks id validation
	for idx, tip := range p.Sinks {
		if _, e := existIds[tip.Id]; !e {
			return fmt.Errorf(
				"id for sinks %d not exists: %d",
				idx, tip.Id,
			)
		}
	}

	// all check passed
	return nil
}
