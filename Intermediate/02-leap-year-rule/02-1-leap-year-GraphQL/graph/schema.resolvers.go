package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"math"

	"github.com/mkdtemplar/SimpleProblems/Intermediate/02-1-leap-year-GraphQL/graph/model"
)

// GetLeapYears is the resolver for the getLeapYears field.
func (r *queryResolver) GetLeapYears(ctx context.Context, input model.InputYear) (*model.LeapYears, error) {

	var difference float64

	for difference < 1 && input.StartYear <= input.EndYear {
		fraction := input.YearDuration - math.Trunc(input.YearDuration)
		difference += fraction

		if difference >= 0.5 {
			difference -= 1
			r.LEAPYEAR.LeapYear = append(r.LEAPYEAR.LeapYear, input.StartYear)
		}
		input.StartYear++
	}

	return &r.LEAPYEAR, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
