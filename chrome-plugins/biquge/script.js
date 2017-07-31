function removeAda(){
	var ids = ["xl_dp_DIV","fdiv","cs_right_bottom","ad_rightBottom"]
	for(var i=0;i<ids.length;i++){
		var id = ids[i]
		if(document.getElementById(id)){
			document.getElementById(id).remove()
		}
	}
	var classNames = ["adshow"]
	for(var i=0;i<classNames.length;i++){
		var className = classNames[i]
		if(document.getElementsByClassName(className).length == 1){
			document.getElementsByClassName(className)[0].remove()
		}
	}
}
removeAda()
setTimeout("removeAda()",3000)
