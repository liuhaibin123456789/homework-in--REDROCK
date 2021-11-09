package main

import "fmt"

func main() {
	var skill,skillMod string
	fmt.Println("少年，你要使用何种技能打败怪兽？")
	_, err := fmt.Scan(&skill)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("少年，你要使用哪种姿态释放技能？")
	_, err = fmt.Scan(&skillMod)
	if err != nil {
		fmt.Println(err)
		return
	}
	ReleaseSkill(skill,skillMod, func(skillName,skillModel string) {
		fmt.Println(skillModel,"\t", skillName)
	})
}

func ReleaseSkill(skillNames,skillMod string, releaseSkillFunc func(string,string)) {
	releaseSkillFunc(skillNames,skillMod)
}
