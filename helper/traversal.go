package helper

import "fmt"




func BigTraversal(paths [][]string,ants int){

	lenPath := len(paths)-1
	chooser :=0
	for i:=1;i<=ants;i++{
		if chooser <= lenPath{
			TraverseOnePath(paths[chooser],i)
			chooser++
		}else{
			chooser =0
			TraverseOnePath(paths[chooser],i)
		}
	}
}


func TraverseOnePath(path []string, antNumb int ){

	for i,v:=range path{
		if i!=0{
			fmt.Printf("L%d-%s\n",antNumb,v)
		}
	}

}