package marketplace

import (
	"github.com/bobanderson-malibu/dotfree/dot-free-market/modules/util"
	"github.com/gocraft/web"
)

func (c *Context) ShowFeed(w web.ResponseWriter, r *web.Request) {
	if len(r.URL.Query()["section"]) > 0 {
		section := r.URL.Query()["section"][0]
		c.SelectedSection = section
	} else {
		c.SelectedSection = ""
	}
	feedItems := CacheGetPublicFeedItems()
	c.ViewFeedItems = feedItems.ViewFeedItems(c.ViewUser.Language, c.SelectedSection)
	util.RenderTemplate(w, "feed", c)
}
