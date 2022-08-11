package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

func main() {
	var name string
	var file_name string
	fmt.Println("请输根目录要解析的文件夹名(默认/images)")
	fmt.Scanln(&file_name)
	fmt.Println("请输入图片名前缀（选填）：") //执行到这里会堵塞
	fmt.Scanln(&name)
	if file_name == "" {
		file_name = "./images/"
	} else {
		file_name = "./" + file_name + "/"
	}
	GetImgList(file_name, name)
}

//读取文件夹里的图片文件 并修改name
func GetImgList(imgPath string, file_s string) {

	if fileInfo, err := ioutil.ReadDir(imgPath); err != nil {
		// 错误处理
		fmt.Println()
		fmt.Println("-----------出错-------------")
		fmt.Println("获取images列表失败" + "---" + imgPath + "路径不存在")
		fmt.Println("----------------------------")
		fmt.Println()
		fmt.Println("操作失败 回车退出")
		fmt.Scanln()
		return
	} else {
		if len(fileInfo) == 0 {
			fmt.Println("-----------出错-------------")
			fmt.Println(imgPath + "为空文件夹")
			fmt.Println("----------------------------")
			fmt.Println()
			fmt.Println("操作失败 回车退出")
			fmt.Scanln()
			return
		}
		for index, fileInfo := range fileInfo {
			// fileInfo.Name() 文件名
			// fileInfo.ModTime().Unix()) 文件添加时间时间戳
			// str := imgPath + "/" + fileInfo.Name()
			// strArr = append(strArr, fileInfo)
			//处理图片
			file_name_end := path.Ext(fileInfo.Name()) //获取文件后缀
			//重命名文件
			err1 := os.Rename(imgPath+fileInfo.Name(), imgPath+file_s+strconv.FormatInt(int64(index), 10)+file_name_end)
			if err1 != nil {
				fmt.Println("图片命名出错")
				return
			} else {
				println(`文件重命名成功`)
			}
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("操作完成 回车退出")
	fmt.Scanln()
}
