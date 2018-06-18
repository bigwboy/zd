package controllers

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"miaopost/frontend/models"
	"time"
	"yhl/help"
)

type InfoMsgController struct {
	BaseController
}

// 留言/回复
func (this *InfoMsgController) CreateMsg() {
	user := this.GetSession("user")
	if user == nil {
		this.SendRes(-1, "请先登录", nil)
	}

	content := this.GetString("content")
	info_id, _ := this.GetInt("info_id")
	pid, _ := this.GetInt("pid")
	im := models.InfoMessage{
		Uid:     user.(*models.User).Id,
		Info_id: int(info_id),
		Pid:     int(pid),
		Content: content,
	}
	i := models.CreateInfoMessage(&im)
	if i > 0 {
		vo := models.ConvertInfoMsgToVo(&im)
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