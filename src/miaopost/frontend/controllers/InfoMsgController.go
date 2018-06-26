package controllers

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"miaopost/frontend/models"
	"time"
	"yhl/help"
	"yhl/wechat"
)

type InfoMsgController struct {
	BaseController
}

// 留言/回复
func (this *InfoMsgController) CreateMsg() {
	u := this.GetSession("user")
	if u == nil {
		this.SendRes(-1, "请先登录", nil)
	}
	user := u.(*models.User)

	content := this.GetString("content")
	info_id, _ := this.GetInt("info_id")
	pid, _ := this.GetInt("pid")
	im := models.InfoMessage{
		Uid:     user.Id,
		Info_id: int(info_id),
		Pid:     int(pid),
		Content: content,
	}
	i := models.CreateInfoMessage(&im)
	if i > 0 {
		vo := models.ConvertInfoMsgToVo(&im)

		// 留言红包处理
		c := help.MongoDb.C("pre_msg_reward")
		var ir *models.InfoReward
		err := c.Find(bson.M{"info_id": int(info_id), "uid": user.Id}).One(&ir)
		if err == nil {
			ir := models.GainReward(ir.Id, user.Id)
			vo.Ireward = ir

			c.Remove(bson.M{"id": ir.Id})
		}

		// 微信提醒回复人
		if int(pid) > 0 {
			go func(pid int) {
				p, err := models.GetInfoMessageById(pid)
				if err != nil {
					return
				}

				user, _ := models.GetUserById(p.Uid)
				help.Log("wxmsg", user)
				viewUrl := this.Ctx.Input.Site() + "/info/view?id=" + help.ToStr(info_id)
				msg := user.Nickname + "回复了你的留言， 查看: " + viewUrl
				wechat.SendTextMsg(user.Openid, msg)
			}(int(pid))
		}

		this.SendRes(0, "success", vo)
	}

	this.SendRes(-1, "failed", nil)
}

// 建删
func (this *InfoMsgController) SuggestDel() {
	id, _ := this.GetInt("id")

	condition := bson.M{
		"msg_id": int(id),
	}
	job := &mgo.MapReduce{
		Map:    "function(){ emit(this.ip, 1) }",
		Reduce: "function(key, values) { return Array.sum(values) }",
	}
	type record struct {
		Ip    string "_id"
		Count int    "value"
	}
	var result []record
	_, err := help.MongoDb.C("info_msg_del_sug").Find(condition).MapReduce(job, &result)
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}
	if len(result) > 3 {
		models.DelInfoMsgById(int(id))
		this.SendRes(0, "success", nil)
	}

	m := map[string]interface{}{
		"msg_id": int(id),
		"time":   time.Now(),
		"ip":     this.Ctx.Input.IP(),
	}
	user := this.GetSession("user")
	if user != nil {
		m["uid"] = user.(*models.User).Id
	}
	help.MongoDb.C("info_msg_del_sug").Insert(m)

	this.SendRes(0, "success", nil)
}

// 赞赏
func (this *InfoMsgController) Admire() {
	mid, _ := this.GetInt("mid")
	// 生成支付订单
	_ = mid
	this.SendRes(0, "success", nil)
	// 支付后个人账号变更
}

// 点赞
func (this *InfoMsgController) Support() {
	id, _ := this.GetInt("id")
	models.Support(int(id))

	this.SendRes(0, "success", nil)
}
