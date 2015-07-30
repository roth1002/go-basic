/*************************
** Mission: Where Am I? **
**************************
**
** Well, this is odd.  The normal GPS tracking system is in place,
** but there's a cryptic function that's restricting your access.
** There's a lot going on, but if you can follow the code, you
** should be able to bypass the restrictions.
**
**/

package main

import (
	"fmt"
	"strings"
)

func main() {

	request := Request{agent: "Dee Fercloze", mission: "Operation Go"}

	location := pingSattelite(request)
	fmt.Println("Location:", location)
}

type Request struct {
	agent   string
	mission string
}

func pingSattelite(request Request) string {
	isAccessible := accessible(request.agent)
	return revealLocation(isAccessible)
}

func accessible(agent string) bool {
	a := strings.Split(agent, " ")
	if len(a) != 3 {
		return false
	}
	b := a[0]
	c := a[1]
	d := a[2]
	x := strings.EqualFold(b, c)
	y := b != strings.ToLower(c)
	z := strings.Index(d, b+c) == 1 && len(d) == 5
	return x && y && z
}
