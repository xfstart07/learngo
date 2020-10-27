// Author: xufei
// Date: 2020/7/2

package department

import (
	"context"
	"errors"
	"learngo/microservices/gokit/one-part/config"
	"learngo/microservices/gokit/one-part/protobuf-spec/common"
	"log"
	"time"

	"google.golang.org/grpc"
)

type DepartmentManager struct {
	Departments         []*common.Department
	UManagerServiceConn *grpc.ClientConn
	UManagerServiceCli  common.UserManagerServiceClient
}

// List ...
func (m *DepartmentManager) List() (departmentList []*common.Department, err error) {
	return m.Departments, nil
}

// Create ...
func (m *DepartmentManager) Create(department *common.Department) (err error) {
	for _, d := range m.Departments {
		if d.Name == department.Name {
			log.Fatal(department)
			return errors.New("目标部门已存在")
		}
	}
	m.Departments = append(m.Departments, department)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	userList := &common.UserList{
		List: department.Users,
	}
	m.UManagerServiceCli.AddUser(ctx, userList)
	return
}

// PersonnelChange 人事调动
func (m *DepartmentManager) PersonnelChange(user *common.User, company string) (err error) {
	for _, department := range m.Departments {
		// 从旧部门中移出
		if department.Name == user.Company {
			for index, _user := range department.Users {
				if _user.Name == user.Name {
					// 没有好的移除切片元素的方法
					department.Users = append(department.Users[:index], department.Users[index+1:]...)
				}
			}
		}
		// 加入新部门
		if department.Name == company {
			user.Company = department.Name
			department.Users = append(department.Users, user)
		}
	}
	return
}

func NewService() *DepartmentManager {
	manager := &DepartmentManager{
		Departments: []*common.Department{
			&common.Department{
				Name: "百度",
				Users: []*common.User{
					&common.User{
						Name:    "李彦宏",
						Company: "百度",
					},
				},
			},
			&common.Department{
				Name: "阿里",
				Users: []*common.User{
					&common.User{
						Name:    "马云",
						Company: "阿里",
					},
				},
			},
			&common.Department{
				Name: "腾讯",
				Users: []*common.User{
					&common.User{
						Name:    "马化腾",
						Company: "腾讯",
					},
				},
			},
		},
	}

	return manager
}

func (manger *DepartmentManager) connectUManagerService() {
	conn, err := grpc.Dial(config.UserManagerGrpcTransportAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	manger.UManagerServiceConn = conn
	manger.UManagerServiceCli = common.NewUserManagerServiceClient(conn)
	log.Println("user manager service connected")
}
