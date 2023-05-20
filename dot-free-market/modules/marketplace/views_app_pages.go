package marketplace

import (
	"github.com/bobanderson-malibu/dotfree/dot-free-market/modules/util"
	"github.com/gocraft/web"
	"net/http"
)

func (c *Context) DownloadApk(w web.ResponseWriter, r *web.Request) {
	user := c.ViewUser.User
	if user.HasDownloadedApp == false {
		user.HasDownloadedApp = true
		user.Save()
	}
	http.Redirect(w, r.Request, "/tochka.apk", 302)
}

func (c *Context) TargetWebUserPage(w web.ResponseWriter, r *web.Request) {
	user := c.ViewUser.User
	if user.HasVisitedDownloadAppPage == false {
		user.HasVisitedDownloadAppPage = true
		user.Save()
	}
	util.RenderTemplate(w, "app_pages/target_web_user_page", c)
}
