package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("欢迎来到猜数字游戏！数字范围是 0-99，你有 7 次机会。")
	r := rand.New(rand.NewSource(time.Now().Unix()))
	target := r.Intn(100) // 0-99
	chance := 7
	var guessed []int
	for {
		fmt.Print("请输入你的猜测：")
		var guess int
		fmt.Scan(&guess)
		if guess == target {
			fmt.Println("恭喜你猜对了！")
			break
		} else if guess < target {
			fmt.Println("太小了，再试一次")
		} else if guess > target {
			fmt.Println("太大了，再试一次")
		}
		guessed = append(guessed, guess)
		chance -= 1
		if chance == 0 {
			fmt.Println("游戏结束，正确答案是：", target)
			fmt.Println("你猜测过的数字：", guessed)
			break
		}
	}
}
