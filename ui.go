package toolbox

type ui struct{

}

func (this *ui) GetPagesBarInfo(current int,total int,showNum int) (map[string]interface{},error){
	info:=map[string]interface{}{
		"pages":[]int{},
	}
	if showNum>total{
		showNum=total
	}
	front:=current-showNum/2
	if front<1{
		front=1
	}
	back:=showNum+front-1
	if back>=total{
		back=total
		front=back-showNum+1
	}
	for i:=front;i<=back;i++{
		info["pages"]=append((info["pages"]).([]int),i)
	}
	info["current"]=current
	if front>1{
		info["prev"]=current-1
		info["first"]=1
	}else{
		info["prev"]=nil
		info["first"]=nil
	}
	if back<total{
		info["next"]=current+1
		info["last"]=total
		info["total"]=total
	}else{
		info["next"]=nil
		info["last"]=nil
		info["total"]=nil
	}
	return info,nil
}