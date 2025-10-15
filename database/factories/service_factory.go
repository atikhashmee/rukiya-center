package factories

import "github.com/brianvoe/gofakeit/v7"

type ServiceFactory struct {
}

// Definition Define the model's default state.
func (f *ServiceFactory) Definition() map[string]any {

	return map[string]any{
		"name":        gofakeit.Name(),
		"description": gofakeit.Sentence(5),
		"price":       gofakeit.RandomInt([]int{10, 500, 40}),
		"image":       gofakeit.Image(200, 200),
		"status":      gofakeit.RandomString([]string{"active", "inactive"}),
	}
}
