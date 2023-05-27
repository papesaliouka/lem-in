package main

import (
	"fmt"
	"lem-in/helper"
	"math"
	"os"
//	"strings"
)

func main() {

	args := os.Args[1:]

	if len(args)>0{
		filename := args[0]
		relations,rooms,ants:=helper.ParseInputFile(filename)

		table:= helper.DFS(relations,helper.PeekStartRoom(rooms).Name)

		// neeed to add another check to see if at least start and end are connected 
		// if all the nodes are not connected
		if len(table)!=len(rooms){
			fmt.Println("ERROR: invalid data format,a not all rooms are connected")
			os.Exit(0)
		}

		start:=helper.PeekStartRoom(rooms).Name
		end:=helper.PeekEndRoom(rooms).Name

		allPaths := helper.FindAllPaths(relations,start,end)
		trimmed := helper.RemoveStart(allPaths)
		groups:= helper.MakeGroups(trimmed)


		paths,_ := helper.GetSmallestPathOfEachGroup(groups)
		thePromised:= helper.GiveTheOneWithMostCandidates(paths,groups)

		elected:= [][]string{append(thePromised,end)}
		min:= math.MaxInt32

		found:=0

		eligeables :=helper.GetEligeables(thePromised,groups)

		if len(eligeables)>0{

			allPossibilities:= helper.GenerateAllPossibilities(thePromised,eligeables)
			for _,v:=range allPossibilities{
				if !helper.HasCommonElements2(v){
					found++
					flatened, flat := helper.Flat2DArray(v)
					if len(flat)<min{
						min = len(flat)
						elected = flatened
					}
				}
			}

			// add End

			for i,path:= range elected{
				elected[i] = append(path, end)
			}


			if found ==0 {
				min := math.MaxInt32
				choosen:=[][]string{}
				for key,subPath:= range groups{
					if  key != thePromised[0]{
						_,flat:= helper.Flat2DArray(subPath)
						if len(flat)>0 && len(flat)<min{
							choosen = subPath
							if len(subPath)>1{
								for _,val:=range subPath{
									if !helper.HasCommonElements(thePromised,val){
										val = append(val, end)
										choosen = [][]string{val}
									}
								}
							}
						}
					}
				}
				elected = append(elected, choosen...)

			}
		}else if len(groups)==1 {
			// that make exampe two work 
			fmt.Println("example2")
			for i,path:= range elected{
				elected[i] = append(path, end)
			}
			elected = append(elected, []string{end})
		}else if len(eligeables)==0 && len(groups)==0{
			//that make exampke 2 work
			elected = append(elected, []string{end})
		}
		
		connextions := helper.ValidateStartingConnections(relations[start],elected)
		helper.BigTraversal(connextions,elected,ants)
	}
}

