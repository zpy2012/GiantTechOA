//chart.html
var arrays = [];
//从页面上隐藏
$(".zjy_lists .am-btn-danger").click(function(){
	var list = $(this).parents(".zjy_lists").index();
//  		$(".zjy_lists").eq(list).remove();
	$(".zjy_lists").eq(list).css("display","none");
	if($(".zjy_lists input[type='checkbox']").eq(list).is(':checked')){ 
		arrays.splice($.inArray(list,arrays),1);
		console.log(arrays)
	}
});
//复选框选中的是列表中的第几个
$(".zjy_lists input[type='checkbox']").click(function(){
	var list = $(this).parents(".zjy_lists").index();
	if($(this).is(':checked')){ 
		arrays.push(list);
		if(arrays.length>3){
			var shiftfirst = arrays.shift();
			$(".zjy_lists input[type='checkbox']").eq(shiftfirst).uCheck('uncheck');
		} 
	}else{
		arrays.splice($.inArray(list,arrays),1)
	}
	console.log(arrays)
});

//单选框选中的是列表中的第几个
$(".zjy_list input[type='radio']").click(function(){
	var list = $(this).parents(".zjy_list").index();
});

//department编辑
$("tr .department_edit_button").click(function(){
	var list = $(this).parents(".zjy_style_list tr").index();
	//这里取了id
 	var id= $(".DepartmentId").eq(list).text();
  var URL = "/department/edit/?id="+id;
  window.location.href=URL;
});

//user编辑
$("tr .user_edit_button").click(function(){
  var list = $(this).parents(".zjy_style_list tr").index();
  //这里取了id
  var id= $(".UserId").eq(list).text();
  var URL = "/user/edit/?id="+id;
  window.location.href=URL;
});

//修改手机号
$('#updatePhone').click(function(){
  $('#updateph').modal({
    relatedTarget: this,
    closeOnConfirm: false,
    closeViaDimmer: false,
    onConfirm: function(e) {
        var url = "/compareCode"
        var phoneNumber = $('#phoneLabel').text();
        var code = $('#phoneCode').val();
        console.log(code);
        if (code == "") {
          // $('#phoneCode').popover('setContent',"请输入验证码！");
          // setTimeout(function(){
          //   $('#phoneCode').popover('close');
          // },1000)//2秒后自动关闭
        }else{
          $.post(url,
          {
            phoneNumber:phoneNumber,
            code:code,
            type:"checkPhone"
          },
          function(data,status){
            if (data == "验证成功") {
              $('#updateph').modal('close');
              $('#newPhone').modal({
                relatedTarget: this,
                closeOnConfirm: false,
                closeViaDimmer: false,
                onConfirm: function(e) {
                    var url = "/compareCode"
                    var newphone = $('#newphoneLabel').val();
                    var newcode = $('#newPhoneCode').val();
                    if (newphone == "") {
                      // $('#phoneCode').popover('setContent',"请输入验证码！");
                      // setTimeout(function(){
                      //   $('#phoneCode').popover('close');
                      // },1000)//2秒后自动关闭
                    }if (newcode == "") {

                    }else{
                      $.post(url,
                      {
                        phoneNumber:newphone,
                        code:newcode,
                        type:"updatePhone"
                      },
                      function(data,status){
                        if (data == "修改成功") {
                          $('#title').text(data);
                          $('#alert').modal('open');
                          setTimeout(function(){
                            $('#alert').modal('close');
                          },1000)//2秒后自动关闭
                          $('#newPhone').modal('close');
                          window.location.reload;
                        }else {
                          $('#title').text(data);
                          $('#alert').modal('open');
                          setTimeout(function(){
                            $('#alert').modal('close');
                          },1000)//2秒后自动关闭
                        }
                      });
                    }
                }
              });
            }else {
              $('#title').text(data);
              $('#alert').modal('open');
              setTimeout(function(){
                $('#alert').modal('close');
              },1000)//2秒后自动关闭
            }
          });
        }
    }
  });
});

//修改密码
$('#updatePasswd').click(function(){
  $('#updateph').modal({
    relatedTarget: this,
    closeOnConfirm: false,
    closeViaDimmer: false,
    onConfirm: function(e) {
        var url = "/compareCode"
        var phoneNumber = $('#phoneLabel').text();
        var code = $('#phoneCode').val();
        if (code == "") {
          // $('#phoneCode').popover('setContent',"请输入验证码！");
          // setTimeout(function(){
          //   $('#phoneCode').popover('close');
          // },1000)//2秒后自动关闭
        }else{
          $.post(url,
          {
            phoneNumber:phoneNumber,
            code:code,
            type:"checkPhone"
          },
          function(data,status){
            if (data == "验证成功") {
              $('#updateph').modal('close');
              $('#newpwd').modal({
                relatedTarget: this,
                closeOnConfirm: false,
                closeViaDimmer: false,
                onConfirm: function(e) {
                    var url = "/user/updatePwd"
                    var passwd1 = $('#pwdLabel1').val();
                    var passwd2 = $('#pwdLabel2').val();
                    if (passwd1 == "") {
                      $('#title').text("密码不能为空");
                      $('#alert').modal('open');
                      setTimeout(function(){
                        $('#alert').modal('close');
                      },1000);//2秒后自动关闭
                    }if (passwd2 == "") {
                      $('#title').text("密码不能为空");
                      $('#alert').modal('open');
                      setTimeout(function(){
                        $('#alert').modal('close');
                      },1000);//2秒后自动关闭
                    }else if (passwd1 != passwd2) {
                      $('#title').text("两次密码应该相同");
                      $('#alert').modal('open');
                      setTimeout(function(){
                        $('#alert').modal('close');
                      },1000);//2秒后自动关闭
                    }else {
                      $.post(url,
                      {
                        passWord:passwd1
                      },
                      function(data,status){
                        if (data == "修改成功") {
                          $('#title').text(data);
                          $('#alert').modal('open');
                          setTimeout(function(){
                            $('#alert').modal('close');
                          },1000)//2秒后自动关闭
                          $('#newpwd').modal('close');
                          window.location.reload;
                        }else {
                          $('#title').text(data);
                          $('#alert').modal('open');
                          setTimeout(function(){
                            $('#alert').modal('close');
                          },1000)//2秒后自动关闭
                        }
                      });
                    }
                }
              });
            }else {
              $('#title').text(data);
              $('#alert').modal('open');
              setTimeout(function(){
                $('#alert').modal('close');
              },1000)//2秒后自动关闭
            }
          });
        }
    }
  });
});

//获取当前手机的验证码
$('#getoldcode').click(function(){
      var phoneNumber = $('#phoneLabel').text();
      var url = "/getCode/?phoneNumber="+phoneNumber;
      $('#title').text("正在获取验证码。。。");
      $('#alert').modal('open');
      $.get(url,
      function(data,status){
        if (data == "获取成功") {
          $('#alert').modal('close');
          $('#getoldcode').addClass("am-disabled");
          Countdown();
        }else {
              $('#title').text("获取失败，请重试!");
              $('#alert').modal('open');
            setTimeout(function(){
              $('#alert').modal('close');
            },1000)//2秒后自动关闭
            }
      });
      var timer = 60;
      function Countdown() {
      if (timer >= 1) {
        console.log("2");
        $('#getoldcode').html("("+timer+")秒后获取");
          timer -= 1;
          setTimeout(function() {
              Countdown();
          }, 1000);
      }
      else{
        $('#getoldcode').removeClass("am-disabled");
        $('#getoldcode').html("获取验证码");
      }
    };
});

//获取新手机验证码
$('#getnewcode').click(function(){
      var phoneNumber = $('#newphoneLabel').val();
      var url = "/getCode/?phoneNumber="+phoneNumber;
      $('#title').text("正在获取验证码。。。");
      $('#alert').modal('open');
      $.get(url,
      function(data,status){
        if (data == "获取成功") {
	        $('#alert').modal('close');
	        $('#getnewcode').addClass("am-disabled");
	        Countdown();
        }else {
            $('#title').text("获取失败，请重试!");
            $('#alert').modal('open');
            setTimeout(function(){
              	$('#alert').modal('close');
            },1000)//2秒后自动关闭
        }
      });
      var timer = 60;
      function Countdown() {
      if (timer >= 1) {
        $('#getnewcode').html("("+timer+")秒后获取");
          timer -= 1;
          setTimeout(function() {
              Countdown();
          }, 1000);
      }
      else{
        $('#getnewcode').removeClass("am-disabled");
        $('#getnewcode').html("获取验证码");
      }
    };
});

//上传按钮
$('#file_upload').click(function(){
  var projectId = $('#ProjectId').text();
  var URL = "/project/upload/?projectId="+projectId;
  window.location.href=URL;
});

//project附件
$("tr .project_file_button").click(function(){
  var list = $(this).parents(".zjy_style_list tr").index();
  //这里取了id
  var id= $(".ProjectId").eq(list).text();
  var URL = "/project/file/?page=1&id="+id;
  window.location.href=URL;
});

//project编辑
$("tr .project_info_button").click(function(){
  var list = $(this).parents(".zjy_style_list tr").index();
  //这里取了id
  var id= $(".ProjectId").eq(list).text();
  var URL = "/project/info/?id="+id;
  window.location.href=URL;
});

//删除部门
$(".department_delete_button").click(function(){
	var list = $(this).parents(".zjy_style_list tr").index();
	//这里取了id
 	var id= $(".DepartmentId").eq(list).text();
  var URL = "/department/delete/?id="+id;
  window.location.href = URL;
});

//删除用户
$(".user_delete_button").click(function(){
  var list = $(this).parents(".zjy_style_list tr").index();
  //这里取了id
  var id= $(".UserId").eq(list).text();
  var URL = "/user/delete/?id="+id;
  window.location.href = URL;
});

//删除项目
$(".project_delete_button").click(function(){
  var list = $(this).parents(".zjy_style_list tr").index();
  //这里取了id
  var id= $(".ProjectId").eq(list).text();
  var URL = "/project/delete/?id="+id;
  window.location.href = URL;
});


$(".zjy_style_list input[type='checkbox']").click(function(){
	var list = $(this).parents(".zjy_style_list tr").index();
	if($(this).is(':checked')){ 
		stylearrays.push(list);
		if(stylearrays.length>3){
			var shiftfirst = stylearrays.shift();
			$(".zjy_style_list input[type='checkbox']").eq(shiftfirst).uCheck('uncheck');
		}
	}else{
		stylearrays.splice($.inArray(list,stylearrays),1)
	}
	console.log(stylearrays)
});
//新增部门
$(".department_new").click(function(){
	window.location.href="/department/add";
});

//新增用户
$(".user_new").click(function(){
  window.location.href="/user/add";
});

//新增项目
$(".project_new").click(function(){
  window.location.href="/project/add";
});

//修改项目名称
$("#updateProjectName").click(function(){
	var projectId = $("#PId").text();
	$('#update_project_name').modal({
	    relatedTarget: this,
	    closeOnConfirm: true,
	    closeViaDimmer: false,
	    onConfirm: function(e) {
	        var url = "/project/edit";
	        var name = $('#project_newname').val();
	        if (name == "") {
	          	$('#reminder_title').text("请输入项目名称");
	            $('#reminder').modal('open');
	            setTimeout(function(){
	              $('#reminder').modal('close');
	            },1000);//2秒后自动关闭
	       	}else {
	        	$.post(url,
	        	{
	        		ProjectId:projectId,
					Type:"newName",
					Name:name
	        	},
	          	function(data,status){
	              	$('#reminder_title').text(data);
		            $('#reminder').modal('open');
		            setTimeout(function(){
		              $('#reminder').modal('close');
		              window.location.reload();
		            },1000);//2秒后自动关闭
	           	});
	       	}
	    }
    });
});

//修改项目介绍
$("#updateProjectDescription").click(function(){
	var projectId = $("#PId").text();
	$('#update_project_description').modal({
	    relatedTarget: this,
	    closeOnConfirm: true,
	    closeViaDimmer: false,
	    onConfirm: function(e) {
	        var url = "/project/edit";
	        var description = $('#project_newdescription').val();
	        if (description == "") {
	          	$('#reminder_title').text("请输入项目介绍");
	            $('#reminder').modal('open');
	            setTimeout(function(){
	              $('#reminder').modal('close');
	            },1000);//2秒后自动关闭
	       	}else {
	        	$.post(url,
	        	{
	        		ProjectId:projectId,
					Type:"newDescription",
					Description:description
	        	},
	          	function(data,status){
	              	$('#reminder_title').text(data);
		            $('#reminder').modal('open');
		            setTimeout(function(){
		              $('#reminder').modal('close');
		              window.location.reload();
		            },1000);//2秒后自动关闭
	           	});
	       	}
	    }
    });
});

//修改项目来源
$("#updateProjectSource").click(function(){
	var projectId = $("#PId").text();
	$('#update_project_source').modal({
	    relatedTarget: this,
	    closeOnConfirm: true,
	    closeViaDimmer: false,
	    onConfirm: function(e) {
	        var url = "/project/edit";
	        var source = $('#project_newsource').val();
        	$.post(url,
        	{
        		ProjectId:projectId,
				Type:"newSource",
				Source:source
        	},
          	function(data,status){
              	$('#reminder_title').text(data);
	            $('#reminder').modal('open');
	            setTimeout(function(){
	              $('#reminder').modal('close');
	              window.location.reload();
	            },1000);//2秒后自动关闭
           	});
       	}
    });
});

//修改项目类型
$("#updateProjectType").click(function(){
	var projectId = $("#PId").text();
	$('#update_project_type').modal({
	    relatedTarget: this,
	    closeOnConfirm: true,
	    closeViaDimmer: false,
	    onConfirm: function(e) {
	        var url = "/project/edit";
	        var projectType = $('#project_newtype').val();
        	$.post(url,
        	{
        		ProjectId:projectId,
				Type:"newType",
				ProjectType:projectType
        	},
          	function(data,status){
              	$('#reminder_title').text(data);
	            $('#reminder').modal('open');
	            setTimeout(function(){
	              $('#reminder').modal('close');
	              window.location.reload();
	            },1000);//2秒后自动关闭
           	});
       	}
    });
});

//项目放弃或者删除
$("#project_del_btn").click(function(){
  var projectId = $("#PId").text();
  var URL = "/project/delete";
  $('#wornning').modal({
	    relatedTarget: this,
	    closeOnConfirm: true,
	    closeViaDimmer: false,
	    onConfirm: function(e) {
    	  $.post(URL,
		  {
		    ProjectId:projectId
		  },
		  function(data,status){
		  	if(data == "刷新"){
		  		window.location.reload();
		  	}else if(data == "成功"){
		  		window.location.href = "/project/?page=1"
		  	}else{
		  		$('#reminder_title').text(data);
		        $('#reminder').modal('open');
		        setTimeout(function(){
		          $('#reminder').modal('close');
		          window.location.reload();
		        },1000);//2秒后自动关闭
		  	}
		  });
	    }
	});
});

//项目操作  
$("#project_operation_btn").click(function(){
  var projectId = $("#PId").text();
  var URL = "/project/operation/?ProjectId="+projectId;
  $.get(URL,
  function(data,status){
  	var type = data;
    switch(data){
    	case "提醒报价":{
    		$.post(URL,
		  	{
		    	ProjectId:projectId,
		    	type:type
		  	},
		  	function(data,status){
		    	$('#reminder_title').text(data);
	            $('#reminder').modal('open');
	            setTimeout(function(){
	              $('#reminder').modal('close');
	              window.location.reload();
	            },1000);//2秒后自动关闭
		  	});
    		break;
    	}
    	case "签约":{
    		$('#project_deal').modal({
                relatedTarget: this,
                closeOnConfirm: true,
                closeViaDimmer: false,
                onConfirm: function(e) {
                    var url = "/project/operation"
                    var dealprice = $('#project_dealprice').val();
                    var dealtime = $('#project_dealtime').val();
                    if (dealprice == "") {
                      	$('#reminder_title').text("请输入成交价");
			            $('#reminder').modal('open');
			            setTimeout(function(){
			              $('#reminder').modal('close');
			            },1000);//2秒后自动关闭
                   	}if (dealtime == "") {
                   		$('#reminder_title').text("请输入成交工期");
			            $('#reminder').modal('open');
			            setTimeout(function(){
			              $('#reminder').modal('close');
			            },1000);//2秒后自动关闭
                   	}else {
                    	$.post(url,
                    	{
                    		ProjectId:projectId,
		    				type:type,
		    				price:dealprice,
		    				time:dealtime
                    	},
                      	function(data,status){
	                      	$('#reminder_title').text(data);
				            $('#reminder').modal('open');
				            setTimeout(function(){
				              $('#reminder').modal('close');
				              window.location.reload();
				            },1000);//2秒后自动关闭
                       	});
                   	}
                }
            });
    		break;
    	}
    	case "提醒结项":{
    		$.post(URL,
		  	{
		    	ProjectId:projectId,
		    	type:type
		  	},
		  	function(data,status){
		    	$('#reminder_title').text(data);
	            $('#reminder').modal('open');
	            setTimeout(function(){
	              $('#reminder').modal('close');
	            },1000);//2秒后自动关闭
		  	});
    		break;
    	}
    	case "报价":{
    		$('#project_baojia').modal({
                relatedTarget: this,
                closeOnConfirm: true,
                closeViaDimmer: false,
                onConfirm: function(e) {
                    var url = "/project/operation"
                    var price = $('#project_price').val();
                    var time = $('#project_time').val();
                    if (price == "") {
                      	$('#reminder_title').text("请输入项目报价");
			            $('#reminder').modal('open');
			            setTimeout(function(){
			              $('#reminder').modal('close');
			            },1000);//2秒后自动关闭
                   	}if (time == "") {
                   		$('#reminder_title').text("请输入预估工期");
			            $('#reminder').modal('open');
			            setTimeout(function(){
			              $('#reminder').modal('close');
			            },1000);//2秒后自动关闭
                   	}else {
                    	$.post(url,
                    	{
	                      	ProjectId:projectId,
			    			type:type,
			    			price:price,
			    			time:time
                    	},
                    	function(data,status){
	                      	$('#reminder_title').text(data);
				            $('#reminder').modal('open');
				            setTimeout(function(){
				              $('#reminder').modal('close');
				              window.location.reload();
				            },1000);//2秒后自动关闭
                    	});
                   	}
                }
            });
    		break;
    	}
    	case "开始项目":{
    		$('#project_start').modal({
                relatedTarget: this,
                closeOnConfirm: true,
                closeViaDimmer: false,
                onConfirm: function(e) {
                	var url = "/project/operation"
	            	$.post(url,
	            	{
		              	ProjectId:projectId,
		    			type:type
	            	},
	            	function(data,status){
		              	$('#reminder_title').text(data);
			            $('#reminder').modal('open');
			            setTimeout(function(){
			              $('#reminder').modal('close');
			              window.location.reload();
			            },1000);//2秒后自动关闭
	             	});
               	}
            });
    		break;
    	}
    	case "结项":{
    		$('#project_end').modal({
                relatedTarget: this,
                closeOnConfirm: true,
                closeViaDimmer: false,
                onConfirm: function(e) {
                	var url = "/project/operation"
                	$.post(url,
                	{
	                  	ProjectId:projectId,
		    			type:type
                	},
                	function(data,status){
	                  	$('#reminder_title').text(data);
			            $('#reminder').modal('open');
			            setTimeout(function(){
			              $('#reminder').modal('close');
			              window.location.reload();
			            },1000);//2秒后自动关闭
                	});
               	}
            });
    		break;
    	}
    	case "变更状态":{
    		$('#project_status').modal({
                relatedTarget: this,
                closeOnConfirm: true,
                closeViaDimmer: false,
                onConfirm: function(e) {
                    var url = "/project/operation"
                    var status = $('#status_select').val();
					$.post(url,
					{
						ProjectId:projectId,
						type:type,
						status:status
					},
					function(data,status){
						$('#reminder_title').text(data);
						$('#reminder').modal('open');
						setTimeout(function(){
							$('#reminder').modal('close');
							window.location.reload();
						},1000);//2秒后自动关闭
					});
                }
            });
    		break;
    	}
    }
  });
});

//imgform-lines.html
//增加图片
$("#doc-form-file").change(function() {
var $file = $(this);
var fileObj = $file[0];
var windowURL = window.URL || window.webkitURL;
var dataURL;
var $img = $("#preview");
if(fileObj && fileObj.files && fileObj.files[0]){
dataURL = windowURL.createObjectURL(fileObj.files[0]);
$img.attr('src',dataURL);
}else{
dataURL = $file.val();
var imgObj = document.getElementById("preview");
imgObj.style.filter = "progid:DXImageTransform.Microsoft.AlphaImageLoader(sizingMethod=scale)";
imgObj.filters.item("DXImageTransform.Microsoft.AlphaImageLoader").src = dataURL;
 
}
});

$(".submit_button").click(function() {
    $(".am-modal-bd").text("正在上传，请稍后....");
    $('#your-modal').modal('open');
    var option = {
        headers : {
            "ClientCallMode" : "ajax"
        }, //添加请求头部
        dataType: "json",
        success : function(data) {
            $('#your-modal').modal('close');
            if (data.Code == "200") {
                $(".am-modal-bd").text(data.Msg);
  		      	  $('#your-modal').modal('open');
      		  		setTimeout(function(){
      		  			$('#your-modal').modal('close')
      		  		},1000);//2秒后自动关闭
                setTimeout(function () {
                    window.location.href = data.JumpUrl;
                }, 1500);
                return false
            }else {
                $(".am-modal-bd").text(data.Msg);
    		      	$('#your-modal').modal('open');
    			  		setTimeout(function(){
    			  			$('#your-modal').modal('close')
    			  		},1000);//2秒后自动关闭
                    return false;
                }
        },
        error : function(data) {
        	$(".am-modal-bd").text(data.Msg);
	      	$('#your-modal').modal('open');
		  		setTimeout(function(){
		  			$('#your-modal').modal('close')
		  		},10000);//2秒后自动关闭
            return false;
        }
    };
    $(".ajax-form").ajaxSubmit(option);
    return false; //最好返回false，因为如果按钮类型是submit,则表单自己又会提交一次;返回false阻止表单再次提交      
});

$("#userImg").click(function() {
	$("#userFile").click();
});

$("#userFile").change(function() {
		$("#userImage").click();
});

$("#userImage").click(function(){
	$("#title").text("正在上传，请稍后....");
    $('#alert').modal('open');
    var option = {
        headers : {
            "ClientCallMode" : "ajax"
        }, //添加请求头部
//      dataType: "json",
        success : function(data) {
            $('#alert').modal('close');
            if (data == "成功") {
                window.location.reload();
                return false
            }else {
                $("#title").text(data);
    		      	$('#alert').modal('open');
    			  		setTimeout(function(){
    			  			$('#alert').modal('close')
    			  		},1000);//2秒后自动关闭
                    return false;
                }
        },
        error : function(data) {
        	$("#title").text(data);
	      	$('#alert').modal('open');
		  		setTimeout(function(){
		  			$('#alert').modal('close')
		  		},10000);//2秒后自动关闭
            return false;
        }
    };
    $(".ajax-form").ajaxSubmit(option);
    return false; //最好返回false，因为如果按钮类型是submit,则表单自己又会提交一次;返回false阻止表单再次提交
});
