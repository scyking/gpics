package windows

import (
	"github.com/lxn/walk"
	"log"
)

// 释放及清空容器中绑定组件
func ClearWidgets(container walk.Container) {
	widgets := container.Children()
	if widgets != nil {
		container.SetSuspended(true)
		defer container.SetSuspended(false)

		for i := widgets.Len() - 1; i >= 0; i-- {
			widgets.At(i).Dispose()
		}

		if err := widgets.Clear(); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("widgets clear ok!")
}

// 清空容器中图片组件背景
func ClearImageViewBackground(container walk.Container) {
	widgets := container.Children()
	if widgets != nil {
		for i := widgets.Len() - 1; i >= 0; i-- {
			if iv, ok := widgets.At(i).(*walk.ImageView); ok {
				iv.SetBackground(walk.NullBrush())
			}
		}
	}
}
