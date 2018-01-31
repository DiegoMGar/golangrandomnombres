package main

import (
	"fmt"
	"math/rand"
	"time"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
) 

func main(){
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(time.Now().UnixNano())
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		oldName := strings.ToLower(file.Name())
		isEXE := strings.HasSuffix(oldName,".exe")
		isINI := strings.HasSuffix(oldName,".ini")
		if(isEXE || isINI){
			continue;
		}
		index := strings.Index(oldName,"_")
		realName := oldName[index+1:len(oldName)]
		newName := strconv.Itoa(r.Int())+"_"+realName
		err = os.Rename(file.Name(),newName)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(oldName+" to "+newName)
	}
}