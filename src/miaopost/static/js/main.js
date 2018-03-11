$(function(){

	// create.tpl
	$("#create_info_btn").click(function(){
		$this = $(this);
		var cid = $("#cid").val();
		var info_content = $("#info_content").val();
		var valid_day = $("#valid_day").val();
		var email = $("#email").val();

		var flag = true;
		$(".error_tips").html("");
		if(cid == "") {
			$(".error_tips").append("<p>请先选择分类</p>");
                        		flag = false;
		}
		if(info_content == "") {
			$(".error_tips").append("<p>信息内容为必填</p>");
                        		flag = false;
		}
		if(valid_day!="") {
			if(!/^[0-9]+$/.test(valid_day)){
			        $(".error_tips").append("<p>自动删除发布请填写数字</p>");
                        		        flag = false;
			}
		}
		if(email!="") {
			if(!/^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$/.test(email)){
			        $(".error_tips").append("<p>邮箱格式不正确</p>");
                        		        flag = false;
			}
		}

		 if (!flag) {
	                        return false;
	                }

	                $this.attr("disabled","disabled");
	                $.post("/info/create", {cid:cid, content:info_content, valid_day:valid_day, email:email}, function(e){
	                	$this.removeAttr("disabled");
	                        	if(e.code<0) {
	                                	$(".error_tips").append(e.msg);
	                                	return false;
	                        	}

	                        	window.location = "/info/view?id="+e.data.Id;
	                });



	});

	
});