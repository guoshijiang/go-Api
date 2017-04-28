//==================================================================
//创建时间：2017-4-23 首次创建
//功能描述：数据相关结构体
//创建人：郭世江
//修改记录：若要修改请记录 郭世江修改加入路径处理
//==================================================================
package types
type Manager struct {
	AccountName       string  `json:"account_name"`       //管理员名字
	ManagerCode       string  `json:"manager_code"`       //管理员操作号
	HosCode           string  `json:"hos_code"`          //医院代码
	OrgCode           string  `json:"org_code"`          //组织代码
}
type HeartData struct {
	Watch_mark       *string    `json:"watch_mark"`          //腕表号
	Examination_no   *string    `json:"examination_no"`     //体检单号
	Heart_rate       *string    `json:"heart_rate"`         //心律
	Time_pd_rr       *string    `json:"time_pd_rr"`         //rr间期
	Time_pd_pr       *string    `json:"time_pd_pr"`         //P-R间期
	Qrs_front        *string    `json:"qrs_front"`          //心电轴
	Pd_qrs           *string    `json:"pd_qrs"`             //QRS波
	Pd_qt            *string    `json:"pd_qt"`              //Q—T间期
	Pd_qtc           *string    `json:"pd_qtc"`             //QTC间期
	Wave_rs          *string    `json:"wave_rs"`            //rs波
	Wave_rv5         *string    `json:"wave_rv_5"`         //rv5波
	Wave_rv6         *string    `json:"wave_rv_6"`         //rv6波
	Heart_analysis   *string    `json:"heart_analysis"`   //心拍分析
	Gread            *string    `json:"gread"`		//心电机输出结果
	Discode          *string    `json:"discode"`           //明苏达码+结果
	Miscode          *string    `json:"miscode"`           //明苏达码
	Confirm          *string    `json:"confirm"`           //医生确认
	Annotation       *string    `json:"annotation"`        //医生注解
	Result           *string    `json:"result"`            //体检结果
}

type Person struct {
	Examination_no    string              `json:"examination_no"`  //体检单号
	Name              string              `json:"name"`            //体检人名
	Sex               interface{}        `json:"sex"`		//性别
	Age               interface{}        `json:"age"`             //年龄
	CheckupDate       string              `json:"checkup_date"`  //体检时间
}

type Data struct {
	Personinfo Person `json:"personinfo"`
	Heartdata []HeartData `json:"heartdata"`
}
