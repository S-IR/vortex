package vortex_test

import (
	"testing"

	v "github.com/s-ir/vortex"
	s "github.com/s-ir/vortex/styles"
)

func TestMain(t *testing.T) {
	// divStyles := v.Styles{
	// 	W:  200,
	// 	H:  200,
	// 	Bg: s.Cyan100,
	// }

	a := v.App{
		Body: []v.Component{
			// &v.Div{
			// 	Styles: v.Styles{
			// 		W:  400,
			// 		H:  200,
			// 		Bg: s.Cyan100,
			// 		Transitions: map[v.StyleField]v.TransitionProperty{
			// 			v.All: {
			// 				TimingFunction: v.Linear,
			// 				Duration:       300,
			// 			},
			// 		},
			// 	},
			// 	StylesHover: func(start v.Styles) v.Styles {
			// 		start.Bg = s.Red600
			// 		return start
			// 	},
			// 	OnClick: func(div *v.Div, e *v.MouseEvent) {
			// 	},
			// 	OnHover: func(div *v.Div, e *v.MouseEvent) {
			// 		// div.Styles.Bg = s.Black
			// 		// fmt.Println("YOU HOVERED")
			// 	},
			// },
			&v.Div{
				Styles: v.Styles{
					W:  200,
					H:  200,
					Bg: s.Red500,
					Transitions: map[v.StyleField]v.TransitionProperty{
						v.All: {
							TimingFunction: v.Linear,
							Duration:       100,
						},
					},
				},
				StylesHover: func(start v.Styles) v.Styles {
					// start.Bg = s.Black
					start.W = 600
					return start
				},

				OnClick: func(t *v.Div, e *v.MouseEvent) {
				},
				OnHover: func(div *v.Div, e *v.MouseEvent) {
					// fmt.Println("YOU HOVERED 2")

				},
			},
			// &v.P{
			// 	Styles: v.Styles{
			// 		Color: s.Cyan400,
			// 		Text:  48,
			// 		Font:  s.TimesNewRoman,
			// 	},

			// 	Body: "Hello bob",
			// },
		},
	}

	if err := a.Start(); err != nil {
		t.Fatal(err)
	}
	t.Fatalf("App stopped")
	defer a.Stop()
}
