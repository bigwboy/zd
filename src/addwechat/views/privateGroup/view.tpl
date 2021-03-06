<div class="view_pg">
	<div class="viewinfo">
		<div class="row">
			<div class="col-md-9">
				<h4>{{.pg.Name}}</h4>
			</div>
			<div class="col-md-3">
				<a class="glink" href="/pg/list?gid={{.group.Id}}"><span class="glyphicon glyphicon-align-left" aria-hidden="true"></span> {{if .group.Name}} {{.group.Name}}{{else}}{{.group.En_name}}{{end}}</a> 
				<button class="btn btn-warning btn-sm report_pg">举报</button>
			</div>
		</div>
		<div>
			<span>发布时间：{{showtime .pg.Createtime}}</span>
		</div>
		<p class="desc">
			{{.pg.Introduction}}
		</p>
		<div class="row qrcode">
			{{if .pg.Qrcode}}
			 <div class="col-md-2 ">				 				 
					 <img class="img-rounded" src="{{.pg.Qrcode}}" alt="群二维码" />			  
					 <p class="qr1">群二维码</p> 					 
			 </div>
			 {{end}}
			 {{if .pg.Ower_qrcode}}
			 <div class="col-md-2">				 				 
					 <img class="img-rounded" src="{{.pg.Ower_qrcode}}" alt="群主二维码" />	 
					<p class="qr1">群主二维码</p>									  
			 </div>
			 {{end}}
		</div>	
		{{if .pg.Wechat_id}} 
		<p class="qr2"><label>微信号: </label> {{.pg.Wechat_id}}</p> 
		{{end}}
	</div>
	<div class="comment ">
		<h4 class="text-muted">留言</h4>
		<div class="row">	
			<div class="col-md-7">	
				<input type="hidden" id="pg_id" name="pg_id" value="{{.pg.Id}}">
				<textarea id="pg_msg" class="form-control"></textarea>
			</div>
		</div>		 
		<div class="comment_btn col-sm-offset-6">
			<button class="btn btn-success pgmsg-btn">提交</button>
		</div>
		<ul id="commentlist">
			{{range .pgMsgs}}
			<li>				
				<h5>{{.User.Nickname}}</h5>
				<p>{{.Pgm.Content}}</p>
				<p>{{showtime .Pgm.Createtime}}</p>				
			</li>
			{{end}}
		</ul>
	</div>

	<div id="pgReportModal" class="modal fade" tabindex="-1" role="dialog" aria-labelledby="gridSystemModalLabel">
	  <div class="modal-dialog" role="document">
	    <div class="modal-content">
	      <div class="modal-body">
	      	<div class="">
	               <p class="text-muted">如您认为此群不能为同学们带来便利，或有与其他影响正常交流秩序的行为，请举报。</p>
	            </div>
	            <div>
	               <textarea id="pg_report_content" class="form-control" placeholder="举报理由(必填)"></textarea>
	            </div>
	      </div>
	      <div class="modal-footer">
	        <button type="button" class="btn btn-info" data-dismiss="modal">取消</button>
	        <button type="button" class="btn btn-primary" id="report_pg_btn">提交举报</button>
	      </div>
	    </div><!-- /.modal-content -->
	  </div><!-- /.modal-dialog -->
	</div><!-- /.modal -->

</div>