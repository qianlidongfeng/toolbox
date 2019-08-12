package toolbox

import "time"

var(
	cstZone *time.Location
)
func init(){
	cstZone=time.FixedZone("CST", 8*3600)
}

var(
	UI=ui{}
)