package router

import (
	challenge_router "src/api/challenge"
)

func (r *Router) RouterChallenge() {
	c := challenge_router.NewChallenge()
	g := r.s.Group("/challenge")
	g.GET("/challenge", c.Challenge)
}
