package dao

import (
	"fmt"
	"log"
	"strings"

	"github.com/saltbo/zpan/model"
)

func DirExist(uid int64, dir string) bool {
	if dir == "" {
		return true
	}

	items := strings.Split(dir, "/")
	name := items[len(items)-2]
	parent := strings.TrimSuffix(dir, name+"/")
	exist, err := DB.Where("uid=? and name=? and parent=?", uid, name, parent).Exist(&model.Matter{})
	if err != nil {
		log.Panicln(err)
	}

	return exist
}

func FileGet(uid int64, fileId interface{}) (*model.Matter, error) {
	m := new(model.Matter)
	if exist, err := DB.Id(fileId).Get(m); err != nil {
		return nil, err
	} else if !exist {
		return nil, fmt.Errorf("file not exist.")
	} else if m.Uid != uid {
		return nil, fmt.Errorf("file not belong to you.")
	}

	return m, nil
}

func FileCopy(srcFile *model.Matter, dest string) error {
	m := &model.Matter{
		Uid:    srcFile.Uid,
		Name:   srcFile.Name,
		Type:   srcFile.Type,
		Size:   srcFile.Size,
		Parent: dest,
		Object: srcFile.Object,
	}
	_, err := DB.Insert(m)
	return err
}

func FileMove(id int64, dest string) error {
	_, err := DB.ID(id).Cols("parent").Update(&model.Matter{Parent: dest})
	return err
}

func FileRename(id int64, name string) error {
	_, err := DB.ID(id).Cols("name").Update(&model.Matter{Name: name})
	return err
}

func DirRename(id int64, name string) error {
	oldMatter := new(model.Matter)
	if exist, _ := DB.Id(id).Get(oldMatter); !exist {
		return fmt.Errorf("matter not exist.")
	}

	oldParent := fmt.Sprintf("%s%s/", oldMatter.Parent, oldMatter.Name)
	newParent := fmt.Sprintf("%s%s/", oldMatter.Parent, name)
	list := make([]model.Matter, 0)
	_ = DB.Where("parent like '" + oldParent + "%'").Find(&list)

	session := DB.NewSession()
	defer session.Close()
	for _, v := range list {
		v.Parent = strings.Replace(v.Parent, oldParent, newParent, 1)
		_, err := session.ID(v.Id).Cols("parent").Update(v)
		if err != nil {
			_ = session.Rollback()
			return err
		}
	}

	m := &model.Matter{Name: name}
	if _, err := session.ID(id).Cols("name").Update(m); err != nil {
		_ = session.Rollback()
		return err
	}

	return session.Commit()
}
