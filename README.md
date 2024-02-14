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

<h3>Function 1</h3>
<P>So to take day and meal as input and returns the corresponding items available for that meal   i created a function named items</p>
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
<p>and for the breakfast_end i specify the index of "LUNCH" and for the items I go to that column and takes out the data starting from index breakfast_start to breakfast_end</p>
<br>

```
fmt.Println(cols[daynum][breakfast_start+1 : breakfast_end-1])
```

<br>



