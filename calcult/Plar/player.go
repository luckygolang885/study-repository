/* ПРАКТИКА 24.09.2026
package plar

type Player struct {
	Name string
	Lvl int
	Points float64
}
*/

package plar

type Player struct {
	Name string
	Lvl int
	Points float64
}

func LvlUP(player *Player) {
	player.Lvl++
}