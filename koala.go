package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	items := []list.Item{
		Item{title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
		Item{title: "Nutella", desc: "It's good on toast"},
		Item{title: "Bitter melon", desc: "It cools you down"},
		Item{title: "Nice socks", desc: "And by that I mean socks without holes"},
		Item{title: "Eight hours of sleep", desc: "I had this once"},
		Item{title: "Cats", desc: "Usually"},
		Item{title: "Plantasia, the album", desc: "My plants love it too"},
		Item{title: "Pour over coffee", desc: "It takes forever to make though"},
		Item{title: "VR", desc: "Virtual reality...what is there to say?"},
		Item{title: "Noguchi Lamps", desc: "Such pleasing organic forms"},
		Item{title: "Linux", desc: "Pretty much the best OS"},
		Item{title: "Business school", desc: "Just kidding"},
		Item{title: "Pottery", desc: "Wet clay is a great feeling"},
		Item{title: "Shampoo", desc: "Nothing like clean hair"},
		Item{title: "Table tennis", desc: "It’s surprisingly exhausting"},
		Item{title: "Milk crates", desc: "Great for packing in your extra stuff"},
		Item{title: "Afternoon tea", desc: "Especially the tea sandwich part"},
		Item{title: "Stickers", desc: "The thicker the vinyl the better"},
		Item{title: "20° Weather", desc: "Celsius, not Fahrenheit"},
		Item{title: "Warm light", desc: "Like around 2700 Kelvin"},
		Item{title: "The vernal equinox", desc: "The autumnal equinox is pretty good too"},
		Item{title: "Gaffer’s tape", desc: "Basically sticky fabric"},
		Item{title: "Terrycloth", desc: "In other words, towel fabric"},
	}

	title := "My Fave Things"
	m := NewModel(title, items)
	p := tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
