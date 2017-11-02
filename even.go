package main

import "fmt"

func main() {
	tab := []int{1, 2, 38, 45, 847, 45}
	for i := 0; i < len(tab); i++ {
		x, err := even(tab[i])
		if err != nil {
			fmt.Print("error in even method")
			panic("error")
		}
		if x {
			fmt.Println(tab[i])
		}
	}
	for _, v := range tab {

		x, err := even(v)
		if err != nil {
			fmt.Print("error in even method")
			panic("error")
		}
		if x {
			fmt.Println(v)
		}

	}

}
func even(a int) (bool, error) {
	if a%2 == 0 {
		return true, nil
	} else {
		return false, nil
	}
}
