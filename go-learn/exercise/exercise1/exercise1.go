// Role Playing Game

package main

import (
	"fmt"
)

type Item struct {
	Name, Type string
}

type Player struct {
	Name      string
	Inventory []Item
}

func (p *Player) PickUpItem(item Item) {
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s picked up %s\n", p.Name, item.Name)
}

func (p *Player) DropItem(itemName string) {
	if i := p.findItem(itemName); i >= 0 {
		p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
		fmt.Printf("%s dropped %s\n", p.Name, itemName)
	} else {
		fmt.Printf("%s doesn't have %s", p.Name, itemName)
	}
}

func (p *Player) UseItem(itemName string) {
	if i := p.findItem(itemName); i >= 0 {
		item := p.Inventory[i]
		fmt.Printf("%s used %s\n", p.Name, itemName)
		if item.Type == "potion" {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
		}
	} else {
		fmt.Printf("%s doesn't have %s\n", p.Name, itemName)
	}
}

func (p *Player) findItem(name string) int {
	for i, item := range p.Inventory {
		if item.Name == name {
			return i
		}
	}
	return -1
}

func main() {
	player := Player{Name: "Hero"}
	player.PickUpItem(Item{Name: "Health Potion", Type: "potion"})
	player.PickUpItem(Item{Name: "Sword", Type: "weapon"})
	player.UseItem("Health Potion")
	player.DropItem("Sword")
}
