package thefarm

import (
	"errors"
	"fmt"
)

// TODO: Define the SillyNephewError type here.

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}

	amount, err := weightFodder.FodderAmount()

	if amount < 0 {
		if err == nil {
			return 0.0, errors.New("negative fodder")
		} else {
			if err == ErrScaleMalfunction {
				return 0.0, errors.New("negative fodder")
			}
		}
	}

	if cows < 0 {
		return 0.0, fmt.Errorf("silly nephew, there cannot be %v cows", cows)
	}

	if err != nil {
		if err == ErrScaleMalfunction {
			amount = amount * 2
		} else {
			return 0, err
		}
	}
	return amount / float64(cows), nil
}
