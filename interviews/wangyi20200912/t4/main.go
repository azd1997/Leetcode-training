/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 9/12/20 4:16 PM
* @Description: The file is for
***********************************************************************/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(line))

}

func sol() {

}
