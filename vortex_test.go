package vortex_test

import (
	"testing"

	v "github.com/s-ir/vortex"
	s "github.com/s-ir/vortex/styles"
)

func TestMain(t *testing.T) {
	a := v.App{
		Body: []v.Component{
			&v.Div{
				Styles: v.Styles{
					W:  200,
					H:  200,
					Bg: s.Cyan100,
				},
			},
			&v.Div{
				Styles: v.Styles{
					W:  200,
					H:  200,
					Bg: s.Red100,
				},
			},
			&v.P{
				Styles: v.Styles{
					Color: s.Cyan400,
					Text:  s.Text2XL,
				},
				Body: "Hello bob",
			},
		},
	}
	if err := a.Start(); err != nil {
		t.Fatal(err)
	}
	t.Fatalf("App stopped")
	defer a.Stop()
}
