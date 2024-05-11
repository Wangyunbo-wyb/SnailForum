package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

func main() {
	//Create a new Node with a Node number of1 node,err:=snowflake.NewNode(1) if err!=ni1
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	//Generate a snowflake ID.
	id := node.Generate()
	//Print out the ID in a few different ways.
	fmt.Printf("Int64 ID:%d\n", id)
	fmt.Printf("StringID:%s ln", id)
	fmt.Printf("Base2ID:%s ln", id.Base2())
	fmt.Printf("Base64ID:%sln", id.Base64())
	//Print out the ID's timestamp
	fmt.Printf("ID Time:%d\n", id.Time())
	//Print out the ID's node number
	fmt.Printf("ID Node :%dln", id.Node())
	//Print out the ID's sequence number
	fmt.Printf("ID Step :%dln[", id.Step())
	// Generate and print,all in one.
	fmt.Printf("ID :%d \n[", node.Generate().Int64())
}
