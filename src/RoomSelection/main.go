package main

import "sys"

func main(){
	//Kill process 
	sys.Exec("dnx","web")
	explore("http://localhost:500")
}
