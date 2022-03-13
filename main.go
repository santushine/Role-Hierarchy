package main

import (
	"fmt"
)

type role struct {
	name     string
	boss     string
	user     string
	children []*role
}

var roleData = role{}

func main() {

	var rootNodeAvailable bool
	var roolName string

	fmt.Println("~~~~~~~ WELCOME ~~~~~~~~")
	if !rootNodeAvailable {
		fmt.Println("Enter root role name")
		fmt.Scan(&roolName)
		fmt.Println("Role : ", roolName)
		roleData.name = roolName
		rootNodeAvailable = true
	}

	for {
		var operationCode = operations()
		switch operationCode {
		case 1:
			addSubRole()
			fmt.Printf("%+v\n", roleData)
		case 2:
			displayNodes()
		case 3:
		case 4:
			addUserToRole()
		case 5:
			displayUserNames()
		case 6:
			displaySubUsers()
		case 7:
		case 8:
		case 9:
		case 10:
		case 11:
			fmt.Println("--------- Exiting Operations ------------ ")
			return
		default:
			fmt.Println("--------- Exiting Operations ------------ ")
			return
		}

	}
}

func operations() int {
	var operationCode int
	fmt.Println("\n")
	fmt.Println("Please Select an Operatoin")
	fmt.Println("1. Add Sub Role")
	fmt.Println("2. Display Role")
	fmt.Println("3. Delete Role")
	fmt.Println("4. Add User")
	fmt.Println("5. Display User")
	fmt.Println("6. Display Users and Sub-users")
	fmt.Println("7. Delete Users")
	fmt.Println("8. No of Users from Top")
	fmt.Println("9. Height of Role Hierarchy")
	fmt.Println("10. Common Boss of User")
	fmt.Println("11. EXIT")
	fmt.Scan(&operationCode)
	fmt.Println("Entered Operartion is :", operationCode)
	return operationCode
}

func addSubRole() {
	var roleName string
	var boss string
	fmt.Println("Enter Sub Role Name")
	fmt.Scan(&roleName)
	fmt.Println("Enter Reporting to")
	fmt.Scan(&boss)

	child := role{
		name: roleName,
		boss: boss,
	}
	roleData.addChildren(child)
	//bossNode.children = append(bossNode.children, child)
}

func displayNodes() {
	var rolesOrder string
	queue := make([]*role, 0)
	queue = append(queue, &roleData)
	for len(queue) > 0 {
		nextUp := queue[0]
		rolesOrder += nextUp.name + " "
		queue = queue[1:]
		if len(nextUp.children) > 0 {
			for _, child := range nextUp.children {
				queue = append(queue, child)
			}
		}
	}
	fmt.Println("Roles Order")
	fmt.Println(rolesOrder)
	return
}

func (root *role) addChildren(childnode role) {
	fmt.Println("add Child", root.name)

	if root.name == childnode.boss {
		root.children = append(root.children, &childnode)
		// fmt.Printf("%+v\n", root)
		return
	} else {
		for _, child := range root.children {
			child.addChildren(childnode)
		}
	}

}

func addUserToRole() {
	var username string
	var roleName string
	fmt.Println("Enter user Name")
	fmt.Scan(&username)
	fmt.Println("Enter Role")
	fmt.Scan(&roleName)
	roleData.addUser(username, roleName)
}

func (root *role) addUser(user string, role string) {
	fmt.Println("add Child", root.name)

	if root.name == role {
		root.user = user
		return
	} else {
		for _, child := range root.children {
			child.addUser(user, role)
		}
	}

}

func displayUserNames() {
	queue := make([]*role, 0)
	queue = append(queue, &roleData)
	for len(queue) > 0 {
		nextUp := queue[0]
		fmt.Println(nextUp.user + " - " + nextUp.name)
		queue = queue[1:]
		if len(nextUp.children) > 0 {
			for _, child := range nextUp.children {
				queue = append(queue, child)
			}
		}
	}
	return
}

func displaySubUsers() {
	var rootUser = roleData.user
	for _, subUser := range roleData.children {
		var users string
		queue := make([]*role, 0)
		queue = append(queue, subUser)
		for len(queue) > 0 {
			nextUp := queue[0]
			users += nextUp.user + " , "
			queue = queue[1:]
			if len(nextUp.children) > 0 {
				for _, child := range nextUp.children {
					queue = append(queue, child)
				}
			}
		}
		fmt.Println(rootUser + " - " + users)
	}

	return
}
