package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

func items(cols [][]string, day string, meal string) {
	var daynum int
	if day == "MONDAY" {
		daynum = 0
	} else if day == "TUESDAY" {
		daynum = 1
	} else if day == "WEDNESDAY" {
		daynum = 2
	} else if day == "THURSDAY" {
		daynum = 3
	} else if day == "FRIDAY" {
		daynum = 4
	} else if day == "SATURDAY" {
		daynum = 5
	} else if day == "SUNDAY" {
		daynum = 6
	}

	fmt.Println("Day is:-", day)

	if meal == "BREAKFAST" {
		var breakfast_start int
		for j := 0; j < 100; j++ {
			if (cols[daynum][j]) == "BREAKFAST" {
				breakfast_start = j
				break
			}
		}
		var breakfast_end int
		for k := 0; k < 100; k++ {
			if (cols[0][k]) == "LUNCH" {
				breakfast_end = k
				break
			}
		}
		fmt.Println(cols[daynum][breakfast_start+1 : breakfast_end-1])
	} else if meal == "LUNCH" {
		var lunch_start int
		for j := 0; j < 100; j++ {
			if (cols[daynum][j]) == "LUNCH" {
				lunch_start = j
				break
			}
		}
		var lunch_end int
		for k := 0; k < 100; k++ {
			if (cols[0][k]) == "DINNER" {
				lunch_end = k
				break
			}
		}
		fmt.Println(cols[daynum][lunch_start+1 : lunch_end-2])
	} else if meal == "DINNER" {
		var dinner_start int
		for j := 0; j < 100; j++ {
			if (cols[daynum][j]) == "DINNER" {
				dinner_start = j
				break
			}
		}
		var dinner_end int
		for k := 0; k < 100; k++ {
			if (cols[0][k]) == "" {
				dinner_end = k
				break
			}
		}
		fmt.Println(cols[daynum][dinner_start+1 : dinner_end-1])
	}

}

func num_items(cols [][]string, day string, meal string) {
	var daynum int
	if day == "MONDAY" {
		daynum = 0
	} else if day == "TUESDAY" {
		daynum = 1
	} else if day == "WEDNESDAY" {
		daynum = 2
	} else if day == "THURSDAY" {
		daynum = 3
	} else if day == "FRIDAY" {
		daynum = 4
	} else if day == "SATURDAY" {
		daynum = 5
	} else if day == "SUNDAY" {
		daynum = 6
	}
	fmt.Println("Day is:-", day)

	if meal == "BREAKFAST" {
		var breakfast_start int
		for j := 0; j < 100; j++ {
			if (cols[daynum][j]) == "BREAKFAST" {
				breakfast_start = j
				break
			}
		}
		var breakfast_end int
		for k := 0; k < 100; k++ {
			if (cols[daynum][k]) == "LUNCH" {
				breakfast_end = k
				break
			}
		}
		fmt.Println(((breakfast_end) - (breakfast_start)) - 2)
	} else if meal == "LUNCH" {
		var lunch_start int
		for j := 0; j < 100; j++ {
			if (cols[daynum][j]) == "LUNCH" {
				lunch_start = j
				break
			}
		}
		var lunch_end int
		for k := 0; k < 100; k++ {
			if (cols[daynum][k]) == "DINNER" {
				lunch_end = k
				break
			}
		}
		fmt.Println(((lunch_end) - (lunch_start)) - 2)
	}
	if meal == "DINNER" {
		var dinner_start int
		for j := 0; j < 100; j++ {
			if (cols[daynum][j]) == "DINNER" {
				dinner_start = j
				break
			}
		}
		var dinner_end int
		for k := 0; k < 100; k++ {
			if (cols[daynum][k]) == "" {
				dinner_end = k
				break
			}
		}
		fmt.Println(((dinner_end) - (dinner_start)) - 1)
	}

}

func check_iteams(cols [][]string, day string, meal string, iteam_name string) {
	var daynum int
	if day == "MONDAY" {
		daynum = 0
	} else if day == "TUESDAY" {
		daynum = 1
	} else if day == "WEDNESDAY" {
		daynum = 2
	} else if day == "THURSDAY" {
		daynum = 3
	} else if day == "FRIDAY" {
		daynum = 4
	} else if day == "SATURDAY" {
		daynum = 5
	} else if day == "SUNDAY" {
		daynum = 6
	}
	fmt.Println("Day is:-", day)

	if meal == "BREAKFAST" {
		var breakfast_start int
		for j := 0; j < 100; j++ {
			if (cols[daynum][j]) == "BREAKFAST" {
				breakfast_start = j
				break
			}
		}
		var breakfast_end int
		for k := 0; k < 100; k++ {
			if (cols[daynum][k]) == "LUNCH" {
				breakfast_end = k
				break
			}
		}
		iteams_breakfast := (cols[daynum][breakfast_start+1 : breakfast_end-1])
		fmt.Println(iteams_breakfast)
		var checknum int = 0
		for j := 0; j < len(iteams_breakfast); j++ {
			if iteams_breakfast[j] == iteam_name {
				checknum = 1
				break
			}
		}
		if checknum == 1 {
			fmt.Println("Iteam is Present")
		} else {
			fmt.Println("Iteam is absent")
		}
	}

	if meal == "LUNCH" {
		var lunch_start int
		for j := 0; j < 100; j++ {
			if (cols[daynum][j]) == "LUNCH" {
				lunch_start = j
				break
			}

		}
		var lunch_end int
		for k := 0; k < 100; k++ {
			if (cols[daynum][k]) == "DINNER" {
				lunch_end = k
				break
			}
		}
		iteams_lunch := (cols[daynum][lunch_start+1 : lunch_end-1])
		fmt.Println(iteams_lunch)
		var checknum int = 0
		for j := 0; j < len(iteams_lunch); j++ {
			if iteams_lunch[j] == iteam_name {
				checknum = 1
				break
			}
		}
		if checknum == 1 {
			fmt.Println("Iteam is Present")
		} else {
			fmt.Println("Iteam is absent")
		}
	}

	if meal == "DINNER" {
		var dinner_start int
		for j := 0; j < 100; j++ {
			if (cols[daynum][j]) == "DINNER" {
				dinner_start = j
				break
			}

		}
		var dinner_end int
		for k := 0; k < 100; k++ {
			if (cols[daynum][k]) == "" {
				dinner_end = k
				break
			}
		}
		iteams_dinner := (cols[daynum][dinner_start+1 : dinner_end-1])
		fmt.Println(iteams_dinner)
		var checknum int = 0
		for j := 0; j < len(iteams_dinner); j++ {
			if iteams_dinner[j] == iteam_name {
				checknum = 1
				break
			}
		}
		if checknum == 1 {
			fmt.Println("Iteam is Present")
		} else {
			fmt.Println("Iteam is absent")
		}
	}
}

func convert_to_json(cols [][]string) []byte {
	conver_json, _ := json.MarshalIndent(cols, "", "\t")
	return conver_json

}

func save_json_to_file(jsonData string) {
	cwd, _ := os.Getwd()

	filename := filepath.Join(cwd, "menu_data.json")

	os.WriteFile(filename, []byte(jsonData), fs.FileMode(0644))

	fmt.Printf("Saved JSON data to file: %s\n", filename)
}

type details struct {
	day           string
	date          string
	meal          string
	iteam_in_meal map[int]string
}

func main() {
	f, _ := excelize.OpenFile("Sample-Menu.xlsx")
	cols, _ := f.GetCols("Sheet1")
	// Above syntax stores the data column wise in cols

	// Example:-1 returns the corresponding items available for that meal.
	items(cols, "MONDAY", "BREAKFAST")

	// Example:-2 returns the number of items in that meal
	num_items(cols, "MONDAY", "BREAKFAST")

	//Example:-3 Checks if the given item is in that particular meal or not
	check_iteams(cols, "MONDAY", "BREAKFAST", "CORNFLAKES")

	// Example:-4  TO save data TO Jason FIle
	jasonData := convert_to_json((cols))
	save_json_to_file(string(jasonData))

	//Example:-5  method that prints the details of each meal instance
	structure_details := details{"Monday", "05-02-2024", "BREAKFAST", map[int]string{1: "CHOICES OF EGG", 2: "CORNFLAKES", 3: "BREAD +JAM", 4: "POHA", 5: "GRAPES", 6: "TEA+ COFFEE", 7: "MILK"}}
	// fmt.Println(structure_details.iteam_in_meal)
	for _, values := range structure_details.iteam_in_meal {
		fmt.Printf("%v\n", values)
	}

}
