package marketplace

import (
	"net/http"

	"github.com/bobanderson-malibu/dotfree/dot-free-market/modules/util"
	"github.com/gocraft/web"
)

func (c *Context) AdminMessagesShow(w web.ResponseWriter, r *web.Request) {
	if uuid := r.PathParams["uuid"]; uuid != "" {
		thread, err := FindThreadByUuid(uuid)
		if err != nil {
			http.NotFound(w, r.Request)
			return
		}
		c.Thread = *thread
		viewThread := thread.ViewThread(c.ViewUser.Language, c.ViewUser.User)
		c.ViewThread = &viewThread
	}
	util.RenderTemplate(w, "board/admin/message", c)
}
