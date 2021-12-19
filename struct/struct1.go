package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Older(p1, p2 Person) (Person, int) {
	//var tom Person
	//tom.Name, tom.Age = "Tom", 18
	//bob := Person{Name: "Bob", Age: 25}
	//paul := Person{"Paul", 43}
	//tbOlder, tbDiff := Older(tom, bob)
	//tpOlder, tpDiff := Older(tom, paul)
	//bpOlder, bpDiff := Older(bob, paul)
	//fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.Name, bob.Name, tbOlder.Name, tbDiff)
	//fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.Name, paul.Name, tpOlder.Name, tpDiff)
	//fmt.Printf("Of %s and %s, %s is older by %d years\n", bob.Name, paul.Name, bpOlder.Name, bpDiff)
	if p1.Age > p2.Age {
		return p1, p1.Age - p2.Age
	}
	return p2, p2.Age - p1.Age
}

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human	// 匿名字段，struct
	Skills	// 匿名字段，自定义的类型string slice
	int // 内置类型作为匿名字段
	speciality string
}

func main() {
	// 初始化学生jane
	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology"}
	// 现在我们来访问相应的字段
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	// 我们来修改他的skill技能字段
	jane.Skills = []string{"anatomt"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physic", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is ", jane.int)
}
