package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Info))
}

type Info struct {
	Id          int
	Cid         int
	Content     string
	Valid_day   int
	Email       string
	Status      int
	Views       int
	Ip          string
	Create_time time.Time
}

type InfoVo struct {
	Info   Info
	Cat    Category
	Photos []Photo
}

func CreateInfo(info *Info) int {
	info.Ip = help.ClientIp
	info.Create_time = time.Now()
	i, err := orm.NewOrm().Insert(info)
	if err != nil {
		help.Log("error", err.Error())
	}

	return int(i)
}

func UpdateInfo(info *Info) error {
	_, err := orm.NewOrm().Update(info)
	help.Error(err)

	return err
}

func GetInfoById(id int) (*Info, error) {
	info := &Info{Id: id}

	err := orm.NewOrm().Read(info)
	if err != nil {
		help.Log("error", err.Error())
	}

	return info, err
}

func GetInfoByCid(cid int) []Info {
	var infos []Info
	_, err := orm.NewOrm().QueryTable("tbl_info").Filter("cid", cid).Filter("status", 0).OrderBy("-create_time").All(&infos)
	help.Error(err)

	return infos
}

func GetInfoByEmail(email string) []Info {
	var infos []Info
	_, err := orm.NewOrm().QueryTable("tbl_info").Filter("email", email).OrderBy("-create_time").All(&infos)
	if err != nil {
		help.Log("error", err.Error())
	}

	return infos
}

func GetInfoPage(cid, offset, size int) (infos []Info) {
	qs := orm.NewOrm().QueryTable("tbl_info").Filter("status", 0)
	if cid > 0 {
		qs = qs.Filter("cid", cid)
	}
	_, err := qs.OrderBy("-create_time").Limit(size, offset).All(&infos)
	help.Error(err)

	return
}

func GetInfoCount(cid int) int {
	qs := orm.NewOrm().QueryTable("tbl_info").Filter("status", 0)
	if cid > 0 {
		qs = qs.Filter("cid", cid)
	}
	count, err := qs.Count()
	help.Error(err)

	return int(count)
}

func IncInfoViews(id int) bool {
	num, err := orm.NewOrm().QueryTable("tbl_info").Filter("id", id).Update(orm.Params{"views": orm.ColValue(orm.ColAdd, 1)})
	help.Error(err)

	return num > 0
}

func SearchInfo(s string) (infos []Info) {
	_, err := orm.NewOrm().QueryTable("tbl_info").Filter("status", 0).Filter("content__icontains", s).OrderBy("-create_time").Limit(50).All(&infos)
	help.Error(err)

	return
}

func ConvertInfoToVo(info Info) InfoVo {
	vo := InfoVo{}
	vo.Info = info
	vo.Cat = GetCategoryById(info.Cid)
	vo.Photos = GetPhotoByInfoid(info.Id)

	return vo
}

func ConvertInfosToVo(infos []Info) []InfoVo {
	vos := []InfoVo{}
	for _, info := range infos {
		vo := ConvertInfoToVo(info)
		vos = append(vos, vo)
	}

	return vos
}

func DelInfoById(id int) bool {
	i, err := orm.NewOrm().QueryTable("tbl_info").Filter("id", id).Update(orm.Params{"status": -1})
	help.Error(err)

	return i > 0
}

func DelExpireInfo() bool {
	_, err := orm.NewOrm().Raw("UPDATE `tbl_info` SET status=-1  WHERE `status`=0 and `valid_day`>0 and  `create_time` < date_sub(now(),interval `valid_day` day)").Exec()

	help.Error(err)
	if err != nil {
		return false
	}

	return true
}