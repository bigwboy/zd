<div class="alert alert-info" role="alert">
	<p class="text-warning">亲, 你已经免费解锁了两所学校，超过两所需要支付</p>
	<p class="center-block">支付金额 : <span class="pay_amount text-danger">￥2.00</span></p>
	<br>
	<button class="btn btn-success" type="button" onclick="pay()">立即支付</button>
</div>
<input type="hidden" id="gid" name="gid" value="{{.gid}}">


<script type="text/javascript">
function onBridgeReady(){
   WeixinJSBridge.invoke(
       'getBrandWCPayRequest', {
           "appId":"{{.sdk.appId}}",     //公众号名称，由商户传入     
           "timeStamp":"{{.sdk.timeStamp}}",         //时间戳，自1970年以来的秒数     
           "nonceStr":"{{.sdk.nonceStr}}", //随机串     
           "package":"{{.sdk.package}}",     
           "signType":"{{.sdk.signType}}",         //微信签名方式：     
           "paySign":"{{.sdk.paySign}}" //微信签名 
       },
       function(res){     
           if(res.err_msg == "get_brand_wcpay_request:ok" ) {
           	// 使用以上方式判断前端返回,微信团队郑重提示：res.err_msg将在用户支付成功后返回    ok，但并不保证它绝对可靠。 
           	//alert("支付成功");
            prompt("支付解锁成功！");
            setTimeout(function(){
              window.location = "/pg/list?gid="+$("#gid").val();
            }, 2000);            
           }  else if (res.err_msg == "get_brand_wcpay_request:cancel")  {
                     alert("支付过程中用户取消");
            }else{
                    //支付失败
                    alert(res.err_msg)
           }
       }
   ); 
}

function pay() {
	if (typeof WeixinJSBridge == "undefined"){
	   if( document.addEventListener ){
	       document.addEventListener('WeixinJSBridgeReady', onBridgeReady, false);
	   }else if (document.attachEvent){
	       document.attachEvent('WeixinJSBridgeReady', onBridgeReady); 
	       document.attachEvent('onWeixinJSBridgeReady', onBridgeReady);
	   }
	}else{
	   onBridgeReady();
	} 
}
</script>