function removeAda(){
	var ids = ["xl_dp_DIV","fdiv","cs_right_bottom","ad_rightBottom"]
	for(var i=0;i<ids.length;i++){
		var id = ids[i]
		if(document.getElementById(id)){
			document.getElementById(id).remove()
		}
	}
}
removeAda()
setTimeout("removeAda()",3000)
