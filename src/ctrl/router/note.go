package router

import (
	"src/api/note"
	"src/ctrl/mid"
)

func (r *Router) RouterNote() {
	u := note.NewNote()
	g := r.s.Group("/note")
	g.Use(mid.AuthMid())
	g.POST("/add", u.CreateNote)
	g.POST("/list", u.GetNoteList)
	g.POST("/get", u.GetNote)
}
