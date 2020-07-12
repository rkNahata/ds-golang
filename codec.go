package main

import (
	"ds-golang/bTree"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

func (c *Codec) Serialize(root *bTree.Node) string {
	var result string
	encode(root, result)
	return result
}

func encode(root *bTree.Node, result string) {
	if root == nil {
		return
	}
	result = result + strconv.FormatInt(root.Val, 10) + ","
	encode(root.Left, result)
	encode(root.Right, result)
}

func (c *Codec) DeSerialize(data string) *bTree.Node {
	dataList := strings.Split(data, ",")
	intList := convertStringToInt(dataList)
	low, high := 0, len(intList)-1
	node := decode(intList, low, high)
	return node
}

func decode(list []int64, low, high int) *bTree.Node {
	if low > high {
		return nil
	}
	node := bTree.New(list[0])
	leftMostIndex := getIndex(list, node.Val, low+1, high)
	node.Left = decode(list, low+1, leftMostIndex-1)
	node.Right = decode(list, leftMostIndex, high)
	return node
}

func getIndex(dataList []int64, val int64, low, high int) int {
	var i int
	for i = low; i <= high; i++ {
		if val > dataList[i] {
			break
		}
	}
	return i
}

func convertStringToInt(data []string) []int64 {
	var result []int64
	for i := range data {
		val, _ := strconv.Atoi(data[i])
		result = append(result, int64(val))
	}
	return result
}
