<h1>Mess Menu MANAGEMENT</h1>
<p>This Project deals with various function that allow you to extract data from the excel file</p>
<h2>How To setup your Project</h2>
<p>Step1:- Open Terminal and type the following command</p>
<br>

```
go mod init project
```
<br>
then to install excelize use 
<br>


```
go get "github.com/xuri/excelize/v2"
```
<br>

<h2>Excelize Library</h2>
<p>This a very good library for extracting the data from an excel file</p>
To make a file use 
<br>

```
f, _ := excelize.OpenFile("Sample-Menu.xlsx")
```
<br>

To get the data by column wise use 

<br>

```
cols, _ := f.GetCols("Sheet1")
```
<br>

<h3>Function 1</h3>
<P>So to take day and meal as input and returns the corresponding items available for that meal   i created a function named items</p>
<br>

```
func items(cols [][]string, day string, meal string)
```

<br>
<p>The main idea behind this function is take a column from the sheet1 and then i iterate it from top to bottom then if the meal is breakfast then i look for the name "BREAKFAST" in that Column and specify the breakfast_start to the index of "BREAKFAST"</p>
<br>

```
for j := 0; j < 100; j++ {
			if (cols[daynum][j]) == "BREAKFAST" {
				breakfast_start = j
				break
			}
		}
```

<br>
<p>and for the breakfast_end i specify the index of "LUNCH" and for the items I go to that column and takes out the data starting from index breakfast_start+1 to breakfast_end-1</p>
<br>

```
fmt.Println(cols[daynum][breakfast_start+1 : breakfast_end-1])
```

<br>

<h3>Function 2</h3>
This Function return the number of iteams in that meal 
<br>

```
func num_items(cols [][]string, day string, meal string)
```

<br>
So what i have done is fisrt i get the start index and end index of that meal and then use this to get meal
<br>

```
fmt.Println(((breakfast_end) - (breakfast_start)) - 2)
```

<br>

<h3>Function 3</h3>
This Function checks if the given iteam  is in that particular meal or Not
for this I make a function named check_iteams
<br>

```
func check_iteams(cols [][]string, day string, meal string, iteam_name string)
```

<br>
and get the iteam in that meal same as function 2 and then i iterate over it and checks if that iteam is in that meal or not
<br>

```
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
```

<br>

<h2>**I don't use Error Hanlding while making a json file**</h2>

<h3>Function 3 </h3>
This Function Converts the entire data and save it as json file in the same directory
For this we need to import some commands as below
<br>

```
    "encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
```

<br>

```
conver_json, _ := json.MarshalIndent(cols, "", "\t")
```

<br>
This code takes the cols and then store all of its data into conver_json as json data
<br>
Now To save a file that contain all this data we use fuction 
<br>

```
func save_json_to_file(jsonData string) {
	cwd, _ := os.Getwd()

	filename := filepath.Join(cwd, "menu_data.json")

	os.WriteFile(filename, []byte(jsonData), fs.FileMode(0644))

	fmt.Printf("Saved JSON data to file: %s\n", filename)
}
```

<br>
First we take the our working directory location using os.Getwd()<br>
Second code initializes the filename as "menu_data.json" in the current directory.<br>
Third code used for writong the daa in that file<br>
Fourth Code if the file succefully made then print out the that the file has been saved 
<br>









