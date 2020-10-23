package datas

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var Carbines = []gormadapter.CasbinRule{
	{PType: "p", V0: "888", V1: "/base/login", V2: "POST"},
	{PType: "p", V0: "888", V1: "/base/register", V2: "POST"},
	{PType: "p", V0: "888", V1: "/api/createApi", V2: "POST"},
	{PType: "p", V0: "888", V1: "/api/getApiList", V2: "POST"},
	{PType: "p", V0: "888", V1: "/api/getApiById", V2: "POST"},
	{PType: "p", V0: "888", V1: "/api/deleteApi", V2: "POST"},
	{PType: "p", V0: "888", V1: "/api/updateApi", V2: "POST"},
	{PType: "p", V0: "888", V1: "/api/getAllApis", V2: "POST"},
	{PType: "p", V0: "888", V1: "/authority/createAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/authority/deleteAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/authority/getAuthorityList", V2: "POST"},
	{PType: "p", V0: "888", V1: "/authority/setDataAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/authority/updateAuthority", V2: "PUT"},
	{PType: "p", V0: "888", V1: "/authority/copyAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/getMenu", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/getMenuList", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/addBaseMenu", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/getBaseMenuTree", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/addMenuAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/getMenuAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/deleteBaseMenu", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/updateBaseMenu", V2: "POST"},
	{PType: "p", V0: "888", V1: "/menu/getBaseMenuById", V2: "POST"},
	{PType: "p", V0: "888", V1: "/user/changePassword", V2: "POST"},
	{PType: "p", V0: "888", V1: "/user/getUserList", V2: "POST"},
	{PType: "p", V0: "888", V1: "/user/setUserAuthority", V2: "POST"},
	{PType: "p", V0: "888", V1: "/user/deleteUser", V2: "DELETE"},
	{PType: "p", V0: "888", V1: "/fileUploadAndDownload/upload", V2: "POST"},
	{PType: "p", V0: "888", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
	{PType: "p", V0: "888", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
	{PType: "p", V0: "888", V1: "/casbin/updateCasbin", V2: "POST"},
	{PType: "p", V0: "888", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
	{PType: "p", V0: "888", V1: "/casbin/casbinTest/:pathParam", V2: "GET"},
	{PType: "p", V0: "888", V1: "/jwt/jsonInBlacklist", V2: "POST"},
	{PType: "p", V0: "888", V1: "/system/getSystemConfig", V2: "POST"},
	{PType: "p", V0: "888", V1: "/system/setSystemConfig", V2: "POST"},
	{PType: "p", V0: "888", V1: "/system/getServerInfo", V2: "POST"},
	{PType: "p", V0: "888", V1: "/customer/customer", V2: "POST"},
	{PType: "p", V0: "888", V1: "/customer/customer", V2: "PUT"},
	{PType: "p", V0: "888", V1: "/customer/customer", V2: "DELETE"},
	{PType: "p", V0: "888", V1: "/customer/customer", V2: "GET"},
	{PType: "p", V0: "888", V1: "/customer/customerList", V2: "GET"},
	{PType: "p", V0: "888", V1: "/autoCode/createTemp", V2: "POST"},
	{PType: "p", V0: "888", V1: "/autoCode/getTables", V2: "GET"},
	{PType: "p", V0: "888", V1: "/autoCode/getDB", V2: "GET"},
	{PType: "p", V0: "888", V1: "/autoCode/getColume", V2: "GET"},
	{PType: "p", V0: "888", V1: "/sysDictionaryDetail/createSysDictionaryDetail", V2: "POST"},
	{PType: "p", V0: "888", V1: "/sysDictionaryDetail/deleteSysDictionaryDetail", V2: "DELETE"},
	{PType: "p", V0: "888", V1: "/sysDictionaryDetail/updateSysDictionaryDetail", V2: "PUT"},
	{PType: "p", V0: "888", V1: "/sysDictionaryDetail/findSysDictionaryDetail", V2: "GET"},
	{PType: "p", V0: "888", V1: "/sysDictionaryDetail/getSysDictionaryDetailList", V2: "GET"},
	{PType: "p", V0: "888", V1: "/sysDictionary/createSysDictionary", V2: "POST"},
	{PType: "p", V0: "888", V1: "/sysDictionary/deleteSysDictionary", V2: "DELETE"},
	{PType: "p", V0: "888", V1: "/sysDictionary/updateSysDictionary", V2: "PUT"},
	{PType: "p", V0: "888", V1: "/sysDictionary/findSysDictionary", V2: "GET"},
	{PType: "p", V0: "888", V1: "/sysDictionary/getSysDictionaryList", V2: "GET"},
	{PType: "p", V0: "888", V1: "/sysOperationRecord/createSysOperationRecord", V2: "POST"},
	{PType: "p", V0: "888", V1: "/sysOperationRecord/deleteSysOperationRecord", V2: "DELETE"},
	{PType: "p", V0: "888", V1: "/sysOperationRecord/updateSysOperationRecord", V2: "PUT"},
	{PType: "p", V0: "888", V1: "/sysOperationRecord/findSysOperationRecord", V2: "GET"},
	{PType: "p", V0: "888", V1: "/sysOperationRecord/getSysOperationRecordList", V2: "GET"},
	{PType: "p", V0: "888", V1: "/sysOperationRecord/deleteSysOperationRecordByIds", V2: "DELETE"},
	{PType: "p", V0: "888", V1: "/user/setUserInfo", V2: "PUT"},
	{PType: "p", V0: "888", V1: "/email/emailTest", V2: "POST"},
	{PType: "p", V0: "888", V1: "/simpleUploader/upload", V2: "POST"},
	{PType: "p", V0: "888", V1: "/simpleUploader/checkFileMd5", V2: "GET"},
	{PType: "p", V0: "888", V1: "/simpleUploader/mergeFileMd5", V2: "GET"},
	{PType: "p", V0: "8881", V1: "/base/login", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/base/register", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/api/createApi", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/api/getApiList", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/api/getApiById", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/api/deleteApi", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/api/updateApi", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/api/getAllApis", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/authority/createAuthority", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/authority/deleteAuthority", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/authority/getAuthorityList", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/authority/setDataAuthority", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/getMenu", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/getMenuList", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/addBaseMenu", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/getBaseMenuTree", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/addMenuAuthority", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/getMenuAuthority", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/deleteBaseMenu", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/updateBaseMenu", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/menu/getBaseMenuById", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/user/changePassword", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/user/getUserList", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/user/setUserAuthority", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/fileUploadAndDownload/upload", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/casbin/updateCasbin", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/jwt/jsonInBlacklist", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/system/getSystemConfig", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/system/setSystemConfig", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/customer/customer", V2: "POST"},
	{PType: "p", V0: "8881", V1: "/customer/customer", V2: "PUT"},
	{PType: "p", V0: "8881", V1: "/customer/customer", V2: "DELETE"},
	{PType: "p", V0: "8881", V1: "/customer/customer", V2: "GET"},
	{PType: "p", V0: "8881", V1: "/customer/customerList", V2: "GET"},
	{PType: "p", V0: "9528", V1: "/base/login", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/base/register", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/api/createApi", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/api/getApiList", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/api/getApiById", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/api/deleteApi", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/api/updateApi", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/api/getAllApis", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/authority/createAuthority", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/authority/deleteAuthority", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/authority/getAuthorityList", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/authority/setDataAuthority", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/getMenu", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/getMenuList", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/addBaseMenu", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/getBaseMenuTree", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/addMenuAuthority", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/getMenuAuthority", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/deleteBaseMenu", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/updateBaseMenu", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/menu/getBaseMenuById", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/user/changePassword", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/user/getUserList", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/user/setUserAuthority", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/fileUploadAndDownload/upload", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/casbin/updateCasbin", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/jwt/jsonInBlacklist", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/system/getSystemConfig", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/system/setSystemConfig", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/customer/customer", V2: "POST"},
	{PType: "p", V0: "9528", V1: "/customer/customer", V2: "PUT"},
	{PType: "p", V0: "9528", V1: "/customer/customer", V2: "DELETE"},
	{PType: "p", V0: "9528", V1: "/customer/customer", V2: "GET"},
	{PType: "p", V0: "9528", V1: "/customer/customerList", V2: "GET"},
	{PType: "p", V0: "9528", V1: "/autoCode/createTemp", V2: "POST"},
}

func InitCasbinModel(db *gorm.DB) (err error) {
	return db.Transaction(func(tx *gorm.DB) error {
		if !tx.Migrator().HasTable("casbin_rule") {
			if err := tx.Migrator().CreateTable(&gormadapter.CasbinRule{}); err != nil {
				return err
			}
		}
		if tx.Create(&Carbines).Error != nil { // 遇到错误时回滚事务
			return err
		}
		return nil
	})
}
