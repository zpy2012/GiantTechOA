var flower = "0",//选花
	manwoman1 = 0,//  选自己男女
	manwoman2 = 1,//  选对方男女
	relations = 0,//选关系
	styles = "0";//选风格
animation();
function animation(){
	$(".zjy_swiper-slide").css("display","none");
	setTimeout(function(){
		 $(".zjy_swiper-slide").fadeIn(1000);
	},100);
}
//选花
$(".hualist li .zjy_chosebox").click(function(){
	var huali =$(this).parent().index();
	if(huali==0){
		flower = $("#eflowerid1").text();
		$(".hualist li .hua").removeClass("mask");
		$(".hualist li .choose").removeClass("selecteds");
		$(".hualist li .choose").eq(huali).addClass("selected");
		$(".hualist li .hua").eq(huali).addClass("mask");
	}else if(huali==1){
		flower = $("#eflowerid2").text();
		$(".hualist li .hua").removeClass("mask");
		$(".hualist li .choose").removeClass("selecteds").removeClass("selected");
		$(".hualist li .choose").eq(huali).addClass("selecteds");
		$(".hualist li .hua").eq(huali).addClass("mask");
	}else if(huali==2){
		flower = $("#eflowerid3").text();
		$(".hualist li .hua").removeClass("mask");
		$(".hualist li .choose").removeClass("selecteds").removeClass("selected");
		$(".hualist li .choose").eq(huali).addClass("selecteds");
		$(".hualist li .hua").eq(huali).addClass("mask");
	}
});
//  选男女
$(".zjy_manwoman1 li img").click(function(){
	var wm1 = $(this).parent().index();
	if(wm1==0){
		$(this).attr("src","static/img/man.png");
		$(".zjy_manwoman1 li span").removeClass("wmactive");
		$(".zjy_manwoman1 li span").eq(wm1).addClass("wmactive");
		$(".zjy_manwoman1 li img").eq(1).attr("src","static/img/woman01.png");
	}else if(wm1==1){
		$(this).attr("src","static/img/woman.png");
		$(".zjy_manwoman1 li span").removeClass("wmactive");
		$(".zjy_manwoman1 li span").eq(wm1).addClass("wmactive");
		$(".zjy_manwoman1 li img").eq(0).attr("src","static/img/man01.png");
	}
	manwoman1 = wm1;
});
$(".zjy_manwoman2 li img").click(function(){
	var wm2 = $(this).parent().index();
	if(wm2==0){
		$(this).attr("src","static/img/man.png");
		$(".zjy_manwoman2 li span").removeClass("wmactive");
		$(".zjy_manwoman2 li span").eq(wm2).addClass("wmactive");
		$(".zjy_manwoman2 li img").eq(1).attr("src","static/img/woman01.png");
	}else if(wm2==1){
		$(this).attr("src","static/img/woman.png");
		$(".zjy_manwoman2 li span").removeClass("wmactive");
		$(".zjy_manwoman2 li span").eq(wm2).addClass("wmactive");
		$(".zjy_manwoman2 li img").eq(0).attr("src","static/img/man01.png");
	}
	manwoman2 = wm2;
});

//选关系
$(".relation li").click(function(){
	var relation = $(this).index();
	if(relation==0){
		$(".relation li button").removeClass("btncolor").removeClass("btnc").removeClass("btnd");
		$(".relation_a .btna").addClass("btncolor").addClass("btnc");
	}else if(relation==1){
		$(".relation li button").removeClass("btncolor").removeClass("btnc").removeClass("btnd");
		$(".relation_b .btnb").addClass("btncolor").addClass("btnd");
	}else if(relation==2){
		$(".relation li button").removeClass("btncolor").removeClass("btnc").removeClass("btnd");
		$(".relation_c .btnb").addClass("btncolor").addClass("btnd");
	}else if(relation==3){
		$(".relation li button").removeClass("btncolor").removeClass("btnc").removeClass("btnd");
		$(".relation_d .btna").addClass("btncolor").addClass("btnc");
	}
	relations = relation;
});
//选风格
$(".fivestyle button").click(function(){
	var styler = $(this).index();
	if (styler == 0) {
		styles = $("#styleid1").text();
	}else if (styler == 3) {
		styles = $("#styleid2").text();
	}else if (styler == 5) {
		styles = $("#styleid3").text();
	}
	$(".fivestyle button").removeClass("wmactive").removeClass("stylebtn");
	$(this).addClass("wmactive").addClass("stylebtn");
});
//查看
$(".Roselist li a").click(function(){
	var rose = $(this).parent().index();
	$(".Roselist li a").removeClass("Rosebtn").removeClass("wmactive");
	$(this).addClass("wmactive").addClass("Rosebtn");
});
$(".hhsj").click(function(){
	$(this).addClass("wmactive").addClass("hhsjbg");
});

//智能选花
$(".choosebtn").click(function(){
	var URL="/viewdetail";
	$(".choosebtn").addClass("choosebtnbg");
	var id = $(".UserId").text()
	if (flower == "0") {
		flower = $("#eflowerid1").text();
	};
	if (styles == "0") {
		styles = $("#styleid1").text();
	};
	$.post(URL, 
	{
	 	eflower:flower,
      	sex:manwoman1,
      	othersex:manwoman2,
       	relations:relations,
       	styles:styles,
       	userid:id
	},
	function(data) {  
        if (data == "成功") {
        	window.location.href=URL;
        }
    });  
});

//查看花信息
$(".Roselist .detail_button").click(function(){
	var list = $(this).parents(".Roselist").index()
	var id = $(".flowerid").eq(list).text()
	$.post("/viewdetail/flowerdetail", {
	       	flowerid : id
		},
	    function(data) {  
	        if (data == "成功") {
	        	window.location.href="/viewdetail/flowerdetail";
	        }
	    }  
	);  
});
