/*
Copyright 2022 QuanxiangCloud Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package models

import "gorm.io/gorm"

const (
	// Deleted Deleted
	Deleted int64 = 1
	// NotDeleted NotDeleted
	NotDeleted int64 = 0
)

//AppCenter AppCenter
type AppCenter struct {
	ID         string `gorm:"column:id;type:varchar(64);primary_key " json:"id"`
	AppName    string `gorm:"column:app_name;type:varchar(80);" json:"appName"`
	AccessURL  string `gorm:"column:access_url;type:varchar(200);"  json:"accessURL"` //visit url
	AppIcon    string `gorm:"column:app_icon;type:text;"  json:"appIcon"`
	CreateBy   string `gorm:"column:create_by;type:varchar(32);" json:"createBy"`
	UpdateBy   string `gorm:"column:update_by;type:varchar(32);" json:"updateBy"`
	CreateTime int64  `gorm:"column:create_time;type:bigint; " json:"createTime"`
	UpdateTime int64  `gorm:"column:update_time;type:bigint; " json:"updateTime"`
	UseStatus  int    `gorm:"column:use_status;"  json:"useStatus"` //published1，unpublished-1
	DelFlag    int64  `gorm:"column:del_flag;"  json:"delFlag"`     //delete marker 0 not deleted 1 deleted
	// The default time is five days after you click delete.
	// If you click delete in the recycle bin, the delete time changes to the current time
	DeleteTime int64  `gorm:"column:delete_time;type:bigint; " json:"deleteTime"` //default remove
	AppSign    string `gorm:"column:app_sign" json:"appSign"`
}

//TableName get the table name
func (AppCenter) TableName() string {
	return "t_app_center"
}

//AppRepo AppRepo
type AppRepo interface {
	SelectByPage(userID, name string, status, page, limit int, isAdmin bool, db *gorm.DB) ([]AppCenter, int64)
	SelectByID(ID string, db *gorm.DB) *AppCenter
	SelectByName(Name string, db *gorm.DB) *AppCenter
	Insert(app *AppCenter, tx *gorm.DB) error
	Update(app *AppCenter, tx *gorm.DB) error
	Delete(id string, tx *gorm.DB) error
	GetByIDs(tx *gorm.DB, ids ...string) ([]*AppCenter, error)
	UpdateDelFlag(db *gorm.DB, id string, deleteTime int64) error
	GetDeleteList(db *gorm.DB, deleteTime int64) ([]*AppCenter, error)
	SelectByAppSign(db *gorm.DB, appSign string) *AppCenter
}