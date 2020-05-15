package webp2other

import (
	"errors"
	"fmt"
	"golang.org/x/image/webp"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
)

var (
	outSuffix string
	inName    string
	outName   string
	inFile    io.Reader
	outFile   io.Writer
)

func init() {
	// 处理用户输入的文件名
	fmt.Println("请把 webp 文件拖动此到窗口，后面空格隔开\n然后输入要转换的格式：png 或 jpg(默认为jpg)\n然后按回车：")
	_, _ = fmt.Scanln(&inName, &outSuffix)

	// 校验
	if inName == "" || inName[len(inName)-4:] != "webp" {
		err := errors.New("请拖入后缀为 .wepb 的文件")
		outputErr(err)
	}
	if strings.Index("jpgpng", outSuffix) == -1 || outSuffix == "" {
		outSuffix = "jpg"
	}

	// 定义输出文件名
	outName = inName[:len(inName)-4] + outSuffix
}

func Exec() {
	// 打开输入文件
	file1, err := os.Open(inName)
	if err != nil {
		outputErr(err)
	}
	defer file1.Close()
	inFile = file1

	// 打开输出文件
	flag := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file2, err := os.OpenFile(outName, flag, 0666)
	if err != nil {
		outputErr(err)
	}
	defer file2.Close()
	outFile = file2

	// 转换格式
	m, err := webp.Decode(inFile)
	if err != nil {
		outputErr(err)
	}

	switch outSuffix {
	case "png":
		err = png.Encode(outFile, m)
	case "jpg":
		err = jpeg.Encode(outFile, m, &jpeg.Options{Quality: 100})
	}
	if err != nil {
		outputErr(err)
	}

	fmt.Printf("\n文件已保存到：%s\n", outName)
	exit()
}

func outputErr(err error) {
	fmt.Printf("\n程序中止：%s\n", err)
	exit()
}

func exit() {
	fmt.Printf("\n按回车键退出程序...")
	b := make([]byte, 1)
	_, _ = os.Stdin.Read(b)
	os.Exit(0)
}
