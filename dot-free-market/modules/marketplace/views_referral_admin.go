package marketplace

import (
	"github.com/gocraft/web"

	"github.com/bobanderson-malibu/dotfree/dot-free-market/modules/util"
)

func (c *Context) ViewAdminListReferralPayments(w web.ResponseWriter, r *web.Request) {
	c.ReferralPayments = FindReferralPayments()
	util.RenderTemplate(w, "referral/admin/payments", c)
}
