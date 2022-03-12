package main

import (
	"image/color"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// SysInfo saves the basic system information
type SysInfo struct {
	Hostname string `bson:hostname`
	Platform string `bson:platform`
	CPU      string `bson:cpu`
	RAM      uint64 `bson:ram`
	Disk     uint64 `bson:disk`
}

func main() {

	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("\\") // If you're in Unix change this "\\" for "/"

	info := new(SysInfo)

	info.Hostname = hostStat.Hostname
	info.Platform = hostStat.Platform
	info.CPU = cpuStat[0].ModelName
	info.RAM = vmStat.Total / 1024 / 1024
	info.Disk = diskStat.Total / 1024 / 1024

	// fmt.Printf("%+v\n", info)

	// gets GPU info
	Info := exec.Command("cmd", "/C", "wmic path win32_VideoController get name")
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := Info.Output()
	replace := strings.Replace(string(History), "Name", "", -1)
	replace2 := strings.Replace(replace, "LuminonCore IDDCX Adapter", "", -1)

	// gets BOARD info
	Infos := exec.Command("cmd", "/C", "wmic path win32_BaseBoard get Product")
	Infos.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Historys, _ := Infos.Output()
	replaces := strings.Replace(string(Historys), "Product", "", -1)

	// fmt.Println(string(Historys))

	myApp := app.New()
	r, _ := fyne.LoadResourceFromPath("icon.png")
	// change icon
	myApp.SetIcon(r)
	myWindow := myApp.NewWindow("PC-INFO")
	myWindow.CenterOnScreen()

	text1 := canvas.NewText("  Hostname: "+hostStat.Hostname, color.White)
	text2 := canvas.NewText("  Platform: "+hostStat.Platform, color.White)
	text3 := canvas.NewText("  Mainboard: "+replaces, color.White)
	text4 := canvas.NewText("  CPU: "+cpuStat[0].ModelName, color.White)
	text5 := canvas.NewText("  GPU: "+replace2, color.White)
	text6 := canvas.NewText("  RAM: "+strconv.FormatUint(vmStat.Total/1024/1024, 10)[0:2]+"gb", color.White)

	content := container.New(layout.NewHBoxLayout(), text1)
	content2 := container.New(layout.NewHBoxLayout(), text2)
	content3 := container.New(layout.NewHBoxLayout(), text3)
	content4 := container.New(layout.NewHBoxLayout(), text4)
	content5 := container.New(layout.NewHBoxLayout(), text5)
	content6 := container.New(layout.NewHBoxLayout(), text6)

	tabs := container.NewAppTabs(
		container.NewTabItem("SPECS ", container.New(layout.NewVBoxLayout(), content, content2, content3, content4, content5, content6)),
		container.NewTabItem("About ", container.New(layout.NewVBoxLayout(), widget.NewLabel(" Version: 0.0.0 11/03/2022 \n Author: Alejandro"))),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.Resize(fyne.NewSize(300, 320))
	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
