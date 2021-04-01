package windows

import (
	"errors"
	"gpics/base"
	"gpics/base/config"
	"gpics/base/git"
	"log"
)

import (
	"github.com/lxn/walk"
)

type MyMainWindow struct {
	*walk.MainWindow
	tv        *MyTreeView
	sv        *walk.ScrollView
	le        *walk.LineEdit
	ImageName string
	DBSource  map[string]int
}

func (mw *MyMainWindow) errMBox(err error) {
	log.Println(err)
	walk.MsgBox(mw.MainWindow, "错误提示", err.Error(), walk.MsgBoxIconError)
}

func (mw *MyMainWindow) infoMBox(msg string) {
	walk.MsgBox(mw.MainWindow, "消息提示", msg, walk.MsgBoxOK)
}

func (mw *MyMainWindow) dropFiles(fps []string) {
	rootPath := walk.Resources.RootDirPath()
	for _, fp := range fps {
		name, err := base.CopyFile(fp, rootPath)
		if err != nil {
			mw.errMBox(err)
		} else {
			mw.addImageViewWidget(name, mw.sv)
			if err := git.AutoCommit(); err != nil {
				mw.errMBox(err)
			}
		}
	}
}

func (mw *MyMainWindow) commit() {
	if err := git.AutoCommit(); err != nil {
		mw.errMBox(err)
	}
}

func (mw *MyMainWindow) addPic() {
	name, err := mw.openImage()
	if err != nil {
		mw.errMBox(err)
	}
	if name != "" {
		mw.addImageViewWidget(name, mw.sv)
		if err := git.AutoCommit(); err != nil {
			mw.errMBox(err)
		}
	}
}

func (mw *MyMainWindow) config() {
	cf := new(config.Config)
	ws, ok := config.Workspace()

	if !ok {
		mw.errMBox(errors.New("获取工作空间配置失败"))
		return
	}

	cf.Workspace = ws

	cmd, err := RunConfigDialog(mw, cf)
	if err != nil {
		mw.errMBox(err)
		return
	}

	if cmd == walk.DlgCmdOK {

		if err := config.Save(cf); err != nil {
			mw.errMBox(err)
			return
		}

		mw.ImageName = ""

		model := mw.tv.Model().(*DirectoryTreeModel)
		root := NewDirectory(cf.Workspace, nil)
		model.roots = []*Directory{root}

		if err := mw.tv.SetModel(model); err != nil {
			mw.errMBox(err)
			return
		}

		if err := mw.tv.SetCurrentItem(root); err != nil {
			mw.errMBox(err)
			return
		}
	}
}

func (mw *MyMainWindow) clickRadio() {
	log.Println("textType:", mw.DBSource[base.DBTextType])
	if mw.ImageName != "" {
		if err := base.Copy(mw.ImageName, mw.DBSource[base.DBTextType]); err != nil {
			mw.errMBox(err)
		}
	}
}

func OpenDir(owner walk.Form, dir string) (string, error) {
	dlg := new(walk.FileDialog)

	dlg.Title = "选择文件夹"
	dlg.FilePath = dir

	ok, err := dlg.ShowBrowseFolder(owner)

	if err != nil {
		return "", err
	}

	if !ok {
		return "", err
	}

	return dlg.FilePath, nil
}
