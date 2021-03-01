package main

import "fmt"

func main(){
	var a int
	var history []string
	var x,y int
	var operasi string
	
	fmt.Println("halo selamat datang di kalk sederhana")
	fmt.Println("1. Kalkulator")
	fmt.Println("2. history")
	fmt.Println("3. exit")
	fmt.Print("masukkan menu:")
	fmt.Scanln(&a)
 
	switch a { 
		case 1:
			fmt.Println("kalkulator")  
			var i int
			// var currentNumber int
			i = 0
				kalkulator: for true{
						if i == 0 { 
							fmt.Print("Masukkan Angka: ")
							fmt.Scanln(&x)
							fmt.Print("Please enter Operator (+,-,/,*): ")
							fmt.Scanln(&operasi)
							if operasi == "#" {
								break kalkulator
							}
							fmt.Print("Please enter Second number: ")
							fmt.Scanln(&y)
		
						}else{
							fmt.Print("Please enter Operator (+,-,/,%,*): ")
							fmt.Scanln(&operasi)
							if operasi == "#" { 
								break kalkulator
							}
							fmt.Print("Please enter Second number: ")
							fmt.Scanln(&y)
						}
						switch operasi {
						case "+":
							history = append(history, tambah(x, y, &x))
						case "-":
							history = append(history, kurang(x, y, &x)) 
						case "*":
							history = append(history, kali(x, y, &x))
						case "/":
							history = append(history, bagi(x, y, &x)) 
						default:
							fmt.Println("Invalid Operation")
							continue
						}
						fmt.Printf("hasil : ", x) 
						i++
					}//for 

				case 2:
					fmt.Println("Menu history:")
					viewHistory(history)
				case 3:
					fmt.Println("Terimakasih")
					break 
				default:
					fmt.Println("keyword salah", a)

				}//switch
			}//main

			func viewHistory(history[]int) {
				if len(history) == 0 {
					fmt.Println("Tidak ada history!")
				} else {
					for a := 0; a < len(history); a++ {
						fmt.Printf("history: ", history[a])
					}
				}
			}
			func tambah(x,y int, jumlah *int) (history string) {
				*jumlah = x + y
				history = viewHistory(x, y, jumlah, "+")
				return history
			}
			func kali(x,y int, jumlah *int) (history string) {
				*jumlah = x * y
				history = viewHistory(x, y, jumlah, "*")
				return history
			}
			func kurang(x,y int, jumlah *int) (history string) {
				*jumlah = x - y
				history = viewHistory(x, y, jumlah, "-")
				return history
			}
			func bagi(x,y int, jumlah *int) (history string) {
				*jumlah = x / y
				history = viewHistory(x, y, jumlah, "/")
				return history
			}
			func viewHistory(x, y int, jumlah int, operasi string)(history string){
				history = fmt.Sprintf("%.1f", x) + operasi + fmt.Sprintf("%.1f", y)
				history = history + " = " + fmt.Sprintf("%.1f", *jumlah)
				return
			}  
				
			