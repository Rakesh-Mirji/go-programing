package main

import ("fmt"
		"time")

type userId string
type transId string
type amount float64

type User struct{
	id userId
	name, password, address string
	phone int 
	role string
	*account
}

type account struct{
	accountNo, accountType string
	balance amount
	createdOn time.Time
	transactions []*transaction
}

type transaction struct{
	transactionId transId
	amount amount
	from transId
	to transId
}

type BankInterface interface{
	SendOnline(BankInterface, amount) string
	CashDeposit(amount) string
	CheckBalance() string
	AddAccount(string,string,amount)
}

func(u *User) SendOnline(person BankInterface, amount amount)string {
	// fmt.Println(person)
	user,ok := person.(*User)

	if ok == true {
		if u.balance > amount{
			u.balance -= amount
			user.balance += amount
			return  fmt.Sprintf("INR RS %v sent from %s to %s ",amount,u.name, user.name)
		}
	}
	return  fmt.Sprintf("INR RS %v exceeds the balance limit ",amount)
}

func(u *User) CashDeposit(amount amount)string {
		u.balance += amount
		return  fmt.Sprintf("INR RS %v depositted to %s",amount,u.name)
}

func (u *User) CheckBalance()string {
	return fmt.Sprintf("\nName : %v \nAddress : %v\nContact : %v\nAccount No : %v\nBalance : %v ",u.name,u.address,u.phone,u.accountNo,u.balance)
}

func (u *User) AddAccount(accNo,accType string, bal amount ){
	u.account = &account{ accountNo:"0001", accountType:"savings", balance:50000, createdOn:time.Now() }
}


func main(){

	var anil BankInterface = &User{   id:"1111",
							name:"Anil",
							password:"1234",
							address:"savadatti",
							phone:9945575310,
							role:"VIP User" }

	anil.AddAccount("0001","saving",50000)
	
	var pavan BankInterface = &User{   id:"1112",
							name:"Pavan",
							password:"1235",
							address:"Hubli",
							phone:994557540,
							role:"VVIP User" }

	pavan.AddAccount("0002","savings",50000)

	// fmt.Println(anil,pavan)
	fmt.Println(anil.SendOnline( pavan , 50000 ))

	fmt.Println(anil.CashDeposit(40000))
	
	fmt.Println(anil.CheckBalance())
	fmt.Println(pavan.CheckBalance())
}