package main

import (
	"GMacWallpaper/conf"
	"GMacWallpaper/flag"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/reujab/wallpaper"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

/*
目前存在的问题：
	- 桌面必须要显示(不能被其他应用全屏占用)才可以切换，否则无效，也不会报错。而互联网上也没有找到解决的方法。
 */

//查询是否连接了指定的显示器
func isConnectSpecialMonitor()(bool,error){
	cmd := exec.Command("/usr/sbin/system_profiler", "SPDisplaysDataType")
	output, err :=cmd.Output()
	if err!=nil{
		return false,err
	}
	for _,specialMonitor := range conf.GlocalConfig.Special.Monitors{
		if strings.Contains(string(output),specialMonitor){
			return true,nil
		}
	}
	return false,nil
}

//设置壁纸
type Data struct {
	Index int `json:"index"`
	LastTime int `json:"lastTime"`
}
func setWallpaper(path string)error{
	path, err := filepath.Abs(path)
	if err!=nil{
		return errors.New("路径不正确，"+err.Error())
	}
	fileInfo, err := os.Stat(path)
	if err!=nil{
		return errors.New("查询路径信息失败，"+err.Error())
	}
	if fileInfo.IsDir(){
		//读取数据
		var data Data
		dataFilePath := "./data.json"
		dataFile,err := ioutil.ReadFile(dataFilePath)
		if err==nil{
			err = json.Unmarshal(dataFile,&data)
			if err!=nil{
				return errors.New("数据解析失败，"+err.Error())
			}
		}
		//遍历目录
		files, err := ioutil.ReadDir(path)
		if err!=nil{
			return errors.New(fmt.Sprintf("获取目录%s下的文件失败，%s",path,err.Error()))
		}
		if data.LastTime != time.Now().Hour(){
			data.Index = (data.Index + 1)%len(files)
			data.LastTime = time.Now().Hour()
		}
		//换壁纸
		err = wallpaper.SetFromFile(path+"/"+files[data.Index].Name())
		if err!=nil{
			return errors.New("设置壁纸失败，"+err.Error())
		}
		//数据写入文件
		dataByte,err := json.Marshal(data)
		if err!=nil{
			return errors.New("数据Marshal失败，"+err.Error())
		}
		err = ioutil.WriteFile(dataFilePath,dataByte,0644)
		if err!=nil{
			err = os.Remove(dataFilePath)
			if err!=nil{
				return errors.New("数据写入失败后进行删除，但是删除失败，"+err.Error())
			}
		}
	}else{
		//fmt.Println("开始替换壁纸")			//Debug
		err := wallpaper.SetFromFile(path)
		if err!=nil{
			return errors.New("设置壁纸失败，"+err.Error())
		}
	}
	return nil
}

func main(){
	//解析命令行
	flag.ParseFlag()

	//解析配置文件
	conf.ParseConfig(*flag.PathConfFile)

	//判断是否连接了特殊的显示器
	isConnect,err:=isConnectSpecialMonitor()
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	if isConnect{
		err = setWallpaper(conf.GlocalConfig.Special.Path)
	}else{
		err = setWallpaper(conf.GlocalConfig.Default.Path)
	}
	if err!=nil{
		fmt.Println("切换壁纸失败，"+err.Error())
		os.Exit(1)
	}
}
