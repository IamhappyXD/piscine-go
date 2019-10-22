package piscine 

func PrintNbrInOrder(n int) {
	if n <= 0{
		fmt.Print(0)
	}else if 1 <= n && n<=9 {
		fmt.Print(n)
	} else{
		i :=0
		ch :=n
		var a[] int
		for 1<ch{
			a=append(a, ch%10)
			ch -= (ch%10) 
			ch /= 10
			i++
		}
		for j:=0; j< i-1; j++{
			if a[j] > a[j+1]{
				temp := a[j+1]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}

		for j:=0; j< i;j++{
			fmt.Print(a[j])
		}

	
		
	}
}




