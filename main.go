package main

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"os"
	"patterns/utils"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type Patterns interface {
	PatternSemi() ([]int, []int, int, bool)
	PatternInverted() ([]int, []int, int, bool)
	PatternM() ([]int, []int, int, bool)
	PatternW() ([]int, []int, int, bool)
}
type Genericfiles1 interface {
	Genericcsv(filepath string) []Date
	Genericcsvtwo(filepath string) []Date1
}
type Date struct {
	Index  int
	Date   string
	Value1 int
	Value2 int
	Value3 int
	Value4 int
}
type Date1 struct {
	Index  int
	Date   string
	Value4 int
}
type Date_struct struct {
	Date_struct []Date
}

func main() {
	var genfile Genericfiles1 = Date{}
	Date := genfile.Genericcsv("/home/tatva.j@ah.zymrinc.com/Desktop/go-graph/data.csv")
	Date1 := genfile.Genericcsvtwo("/home/tatva.j@ah.zymrinc.com/Desktop/go-graph/data.csv")
	var p Patterns = &Date_struct{Date}
	start, end, count, bol := p.PatternSemi()
	if bol {
		fmt.Println("semicircle/s found:=>", count)
		fmt.Println("Start: ", start, "End: ", end, "Count: ", count)
		// fmt.Println("Start: ", Date[start[0]], "End: ", Date[end[0]], "Count: ", count)
	} else {
		fmt.Println("semicircle /s not found")
	}
	start1, end1, count1, bol1 := p.PatternInverted()
	if bol1 {
		fmt.Println("Inverted/s found:=>", count1)
		fmt.Println("Start: ", start1, "End: ", end1, "Count: ", count1)
		// fmt.Println("Start: ", Date[start1[0]], "End: ", Date[end[0]], "Count: ", count1)
	} else {
		fmt.Println("Inverted /s not found")
	}
	start2, end2, count2, bol2 := p.PatternM()
	if bol2 {
		fmt.Println("M/s found:=>", count2)
		fmt.Println("Start: ", start2, "End: ", end2, "Count: ", count2)
		// fmt.Println("Start: ", Date[start2[0]], "End: ", Date[end2[0]], "Count: ", count2)
		// Plot_Pattern(Date1, count, start, end, "PatternM.png")
	} else {
		fmt.Println("M /s not found")
	}

	start3, end3, count3, bol3 := p.PatternW()
	if bol3 {
		fmt.Println("W/s found:=>", count3)
		fmt.Println("Start: ", start3, "End: ", end3, "Count: ", count3)
		// fmt.Println("Start: ", Date[start3[0]], "End: ", Date[end3[0]], "Count: ", count3)
		// Plot_Pattern(Date1, count1, start1, end1, "PatternW.png")
	} else {
		fmt.Println("W /s not found")
	}
	// var start4, end4, count4 []int
	// count4 = append(count4, count)
	// count4 = append(count4, count1)
	// count4 = append(count4, count2)
	// // count4 = append(count4, count3)

	// start4 = append(start4, start...)
	// start4 = append(start4, start1...)
	// start4 = append(start4, start3...)
	// // start4 = append(start4, start...)

	// // end4 = append(end4, end1...)
	// end4 = append(end4, end...)
	// end4 = append(end4, end1...)
	// end4 = append(end4, end2...)

	Plot_Pattern(Date1, count, count1, count2, count3, start, start1, start2, start3, end, end1, end2, end3, "PatternAll.png")
}
func RandomPoints(x []Date1) plotter.XYs {
	pts := make(plotter.XYs, len(x))

	for i := range x {
		// b, _ := strconv.Atoi(x[i].Date[0:2] + "." + x[i].Date[3:5] + "." + x[i].Date[7:11])

		pts[i].X = float64(x[i].Index)
		pts[i].Y = float64(x[i].Value4)
	}
	return pts
}

func Plot_Pattern(t []Date1, count1, count2, count3, count4 int, a1 []int, a2 []int, a3 []int, a4 []int, b1 []int, b2 []int, b3 []int, b4 []int, Name string) {
	//rand.Seed(int64(0))
	// fmt.Println(a, b, count)
	p := plot.New()
	// a2 := a[:count[0]]
	// b2 := b[:count[0]]

	// a3 := a[count[0] : count[0]+count[1]]
	// b3 := b[count[0] : count[0]+count[1]]

	// a4 := a[count[0]+count[1]-1:]
	// b4 := b[count[0]+count[1]:]
	// a1 := a[count[1]+1:]
	// b1 := b[count[1]+1:]

	// a4 := a[count[2]-1:]

	// b4 := b[count[2]-1:]

	p.Title.Text = Name
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints1(p,
		"Original", RandomPoints(t))

	if err != nil {
		panic(err)
	}
	p.Add(plotter.NewGrid())
	//p.BackgroundColor.RGBA()
	plotter.DefaultLineStyle.Width = vg.Points(2)

	for i := 0; i < count1; i++ {
		plotter.DefaultLineStyle.Color = color.RGBA{R: 255, A: 255}
		plotter.DefaultLineStyle.Width = vg.Points(4)
		// l.LineStyle.Color = color.RGBA{B: 255, A: 255}
		plotutil.AddLinePoints(p, Name[:len(Name)-4], 120, Linepoints(t[a1[i]:b1[i]]))

	}
	for i := 0; i < count2; i++ {
		plotter.DefaultLineStyle.Color = color.RGBA{R: 255, A: 255}
		plotter.DefaultLineStyle.Width = vg.Points(4)
		// l.LineStyle.Color = color.RGBA{B: 255, A: 255}
		plotutil.AddLinePoints(p, Name[:len(Name)-4], 4, Linepoints(t[a2[i]:b2[i]]))

	}
	for i := 0; i < count3; i++ {
		plotter.DefaultLineStyle.Color = color.RGBA{R: 255, A: 255}
		plotter.DefaultLineStyle.Width = vg.Points(4)
		// l.LineStyle.Color = color.RGBA{B: 255, A: 255}
		plotutil.AddLinePoints(p, Name[:len(Name)-4], 8, Linepoints(t[a3[i]:b3[i]]))

	}
	for i := 0; i < count4; i++ {
		plotter.DefaultLineStyle.Color = color.RGBA{R: 255, A: 255}
		plotter.DefaultLineStyle.Width = vg.Points(4)
		// l.LineStyle.Color = color.RGBA{B: 255, A: 255}
		plotutil.AddLinePoints(p, Name[:len(Name)-4], 0, Linepoints(t[a4[i]:b4[i]]))
	}

	// for i := 0; i < count[3]; i++ {
	// 	plotter.DefaultLineStyle.Color = color.RGBA{R: 255, A: 255}
	// 	plotter.DefaultLineStyle.Width = vg.Points(4)
	// 	// l.LineStyle.Color = color.RGBA{B: 255, A: 255}
	// 	plotutil.AddLinePoints(p, Name[:len(Name)-4], 3, Linepoints(t[a4[i]:b4[i]]))

	// }
	if err := p.Save(18*vg.Inch, 18*vg.Inch, Name); err != nil {
		panic(err)
	}
}
func Linepoints(x []Date1) plotter.XYer {
	//pts := make([]plotter.XYer, 3)
	pts := make(plotter.XYs, len(x))
	// var d int
	for i, x1 := range x {

		pts[i].X = float64(x1.Index)
		pts[i].Y = float64(x[i].Value4)
		// fmt.Println(pts[i].X, pts[i].Y)

	}
	return pts
}
func (d Date) Genericcsvtwo(filepath string) []Date1 {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)

	}
	// Get value from cell by given worksheet name and axis.
	var date []Date1
	if err != nil {
		fmt.Println(err)

	}
	csvReader := csv.NewReader(f)
	records, _ := csvReader.ReadAll()
	records = records[1:]
	var Index int

	for _, row := range records {
		Index = Index + 1

		b3, _ := strconv.Atoi(row[4])
		date = append(date, Date1{Index: Index, Date: row[0], Value4: b3})

	}
	return date
}
func (d Date) Genericcsv(filepath string) []Date {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)

	}
	// Get value from cell by given worksheet name and axis.
	var date []Date
	if err != nil {
		fmt.Println(err)

	}
	csvReader := csv.NewReader(f)
	records, _ := csvReader.ReadAll()
	records = records[1:]
	var Index int

	for _, row := range records {
		Index = Index + 1
		b, _ := strconv.Atoi(row[1])
		b1, _ := strconv.Atoi(row[2])
		b2, _ := strconv.Atoi(row[3])
		b3, _ := strconv.Atoi(row[4])
		date = append(date, Date{Index: Index, Date: row[0], Value1: b, Value2: b1, Value3: b2, Value4: b3})

	}
	return date
}

func (l *Date_struct) PatternSemi() ([]int, []int, int, bool) {
	peaks := peaks(l.Date_struct)
	troughs := troughs(l.Date_struct)
	// fmt.Println("Peaks:", peaks)
	// fmt.Println("Troughs:", troughs)
	var All_trends_string []string
	var mixed_unsorted []int
	for i := range peaks {
		mixed_unsorted = append(mixed_unsorted, peaks[i])
	}
	for i := range troughs {
		mixed_unsorted = append(mixed_unsorted, troughs[i])
	}
	mixed_sorted := utils.Sort(mixed_unsorted) //mixed values
	var All_trends_index_values []int
	for _, i := range mixed_sorted {
		if P_up(l.Date_struct[i : i+2]) {
			All_trends_string = append(All_trends_string, "Upward")
			All_trends_index_values = append(All_trends_index_values, i)

		} else {
			All_trends_string = append(All_trends_string, "Downward")
			All_trends_index_values = append(All_trends_index_values, i)
		}
	}
	// fmt.Println(All_trends_index_values)
	var start []int
	var end []int
	var count int
	var bol bool //just to check if the pattern is semicircle or not,this is boolean that will tell us that.

	for i := range All_trends_string {
		if All_trends_string[i] == "Upward" && i != len(All_trends_string)-1 {
			if All_trends_string[i+1] == "Downward" && i+1 != len(All_trends_string)-1 { //Task:=>check here for first and last point to be in the align manner.

				var a1 []int //start point
				var x1 int

				for i1 := 0; i1 < 80; i1++ {
					x1 = l.Date_struct[All_trends_index_values[i]+1].Value2 + i1
					a1 = append(a1, x1)
				}
				for i2 := 0; i2 < 80; i2++ {
					x1 = l.Date_struct[All_trends_index_values[i]+1].Value2 - i2
					a1 = append(a1, x1)
				}

				var a2 []int //end point
				var x int

				for i1 := 0; i1 < 80; i1++ {
					x = l.Date_struct[All_trends_index_values[i+2]].Value3 + i1
					a2 = append(a2, x)
				}
				for i2 := 0; i2 < 80; i2++ {
					x = l.Date_struct[All_trends_index_values[i+2]].Value3 - i2
					a2 = append(a2, x)
				}
				//task:try using this below code from the individual files likes semi and inverted.
				// if Is_Valid12(l.Date_struct[All_trends_index_values[i]].Value2, a1) && Is_Valid12(l.Date_struct[All_trends_index_values[i+2]+2].Value3, a2) {
				// 	start = append(start, All_trends_index_values[i])
				// 	end = append(end, All_trends_index_values[i+2]+3)
				// 	count++
				// 	bol = true
				// } else if Is_Valid12(l.Date_struct[All_trends_index_values[i]-2].Value2, a1) && Is_Valid12(l.Date_struct[All_trends_index_values[i+2]].Value3, a2) {
				// 	start = append(start, All_trends_index_values[i])
				// 	end = append(end, All_trends_index_values[i+2]-2)
				// 	count++
				// 	bol = true
				// } else if Is_Valid12(l.Date_struct[All_trends_index_values[i]].Value2, a1) && Is_Valid12(l.Date_struct[All_trends_index_values[i+2]].Value3, a2) {
				// 	start = append(start, All_trends_index_values[i]+1)
				// 	end = append(end, All_trends_index_values[i+2]-2)
				// 	count++
				// 	bol = true
				// } else if Is_Valid12(l.Date_struct[All_trends_index_values[i]].Value2, a1) && Is_Valid12(l.Date_struct[All_trends_index_values[i+3]-1].Value3, a2) {
				// 	start = append(start, All_trends_index_values[i])
				// 	end = append(end, All_trends_index_values[i+3])
				// 	count++
				// 	bol = true
				// } else if Is_Valid12(l.Date_struct[All_trends_index_values[i]+1].Value2, a1) && Is_Valid12(l.Date_struct[All_trends_index_values[i+2]].Value3, a2) {
				// 	start = append(start, All_trends_index_values[i])
				// 	end = append(end, All_trends_index_values[i+2]+1)
				// 	count++
				// 	bol = true
				// }
				if IsValidCategoryu(l.Date_struct[All_trends_index_values[i]].Value3, a2) && IsValidCategoryu(l.Date_struct[All_trends_index_values[i+2]].Value2, a1) {
					start = append(start, All_trends_index_values[i])
					end = append(end, All_trends_index_values[i+2]+1)
					count++
					bol = true
				} else if IsValidCategoryu(l.Date_struct[All_trends_index_values[i]].Value2, a1) && IsValidCategoryu(l.Date_struct[All_trends_index_values[i+3]].Value2, a2) {
					start = append(start, All_trends_index_values[i])
					end = append(end, All_trends_index_values[i+3]+1)
					count++
					bol = true
				} else if IsValidCategoryu(l.Date_struct[All_trends_index_values[i]+1].Value2, a1) && IsValidCategoryu(l.Date_struct[All_trends_index_values[i+2]-3].Value3, a1) {
					start = append(start, All_trends_index_values[i]+1)
					end = append(end, All_trends_index_values[i+2]-2)
					count++
					bol = true
				} else if IsValidCategoryu(l.Date_struct[All_trends_index_values[i]].Value2, a1) && IsValidCategoryu(l.Date_struct[All_trends_index_values[i+2]-3].Value2, a2) {
					start = append(start, All_trends_index_values[i]+1)
					end = append(end, All_trends_index_values[i+2]-2)
					count++
					bol = true
				}
				// else if IsValidCategoryu(l.Date_struct[All_trends_index_values[i+1]].Value2, a1) {
				// 	start = append(start, All_trends_index_values[i+1])
				// 	end = append(end, All_trends_index_values[i+3])
				// 	count++
				// 	bol = true //3
				// }
				// else if IsValidCategoryu(l.Date_struct[All_trends_index_values[i]].Value2, a1) && IsValidCategoryu(l.Date_struct[All_trends_index_values[i+2]].Value3, a2) {
				// 	start = append(start, All_trends_index_values[i])
				// 	end = append(end, All_trends_index_values[i+2]+1)
				// 	count++
				// 	bol = true
				// }

			}
		}

	}

	// fmt.Println("Count =", count)
	return start, end, count, bol
}

func peaks(pts []Date) (a []int) {
	var b []int
	for i := range pts {
		if pts[i].Value2 == pts[0].Value2 {
			continue
		} else if pts[i].Value2 == pts[len(pts)-1].Value2 {
			continue
		} else if pts[i-1].Value2 < pts[i].Value2 && pts[i+1].Value2 < pts[i].Value2 && pts[i-1] != pts[0] && pts[i+1] != pts[len(pts)-1] {

			var a1 []int
			var x int
			for i1 := 0; i1 < 80; i1++ {
				x = pts[i].Value2 + i1
				a1 = append(a1, x)
			}
			for i2 := 0; i2 < 80; i2++ {
				x = pts[i].Value2 - i2
				a1 = append(a1, x)
			}
			if IsValidCategoryu(pts[i-1].Value2, a1) && pts[i-1].Value2 > pts[i-2].Value2 { //pts[i-1].Value == pts[i].Value || pts[i+1].Value == pts[i].Value
				b = append(b, i)
				continue
			}
			if IsValidCategoryu(pts[i+1].Value2, a1) && pts[i+1].Value2 > pts[i+2].Value2 { //pts[i-1].Value == pts[i].Value || pts[i+1].Value == pts[i].Value
				b = append(b, i)
				continue

			}

		} else {
			continue
		}
	}
	return b
}

func troughs(pts []Date) (a []int) {
	var b []int
	for i := range pts {
		if pts[i].Value3 == pts[0].Value3 {
			continue
		} else if pts[i].Value3 == pts[len(pts)-1].Value3 {
			continue
		} else if pts[i-1].Value3 > pts[i].Value3 && pts[i+1].Value3 > pts[i].Value3 && pts[i-1] != pts[0] && pts[i+1] != pts[len(pts)-1] {

			b = append(b, i)
			continue

		}

	}
	return b

}
func Is_Valid1(category int, n []int) bool {
	switch category {
	case
		n[0],
		n[1],
		n[2],
		n[3],
		n[4],
		n[5],
		n[6],
		n[7],
		n[8],
		n[9],
		n[10],
		n[11],
		n[12],
		n[13],
		n[14],
		n[15],
		n[16],
		n[17],
		n[18],
		n[19],
		n[20],
		n[21],
		n[22],
		n[23]:
		// n[24],
		// n[25],
		// n[26],
		// n[27],
		// n[28],
		// n[29],
		// n[30],
		// n[31],
		// n[32],
		// n[33],
		// n[34]:
		// n[35],
		// n[36],
		// n[37],
		// n[38],
		// n[39]:
		// n[40],
		// n[41],
		// n[42],
		// n[43],
		// n[44],
		// n[45],
		// n[46],
		// n[47],
		// n[48]:

		return true
	}
	return false
}

func Is_Valid11(category int, n []int) bool {
	switch category {
	case
		n[0],
		n[1],
		n[2],
		n[3],
		n[4],
		n[5],
		n[6],
		n[7],
		n[8],
		n[9],
		n[10],
		n[11],
		n[12],
		n[13],
		n[14],
		n[15],
		n[16],
		n[17],
		n[18],
		n[19],
		n[20],
		n[21],
		n[22],
		n[23],
		n[24],
		n[25],
		n[26],
		n[27],
		n[28],
		n[29],
		n[30],
		n[31],
		n[32],
		n[33],
		n[34],
		n[35],
		n[36],
		n[37],
		n[38],
		n[39],
		n[40],
		n[41],
		n[42],
		n[43],
		n[44],
		n[45],
		n[46],
		n[47],
		n[48],
		n[49],
		n[50],
		n[51],
		n[52],
		n[53],
		n[54],
		n[55],
		n[56],
		n[57],

		n[58],
		n[59]:
		return true
	}
	return false
}
func Is_Valid30(category int, n []int) bool {
	switch category {
	case
		n[0],
		n[1],
		n[2],
		n[3],
		n[4],
		n[5],
		n[6],
		n[7],
		n[8],
		n[9],
		n[10],
		n[11],
		n[12],
		n[13],
		n[14],
		n[15],
		n[16],
		n[17],
		n[18],
		n[19],
		n[20],
		n[21],
		n[22],
		n[23],
		n[24],
		n[25],
		n[26],
		n[27],
		n[28],
		n[29]:
		return true
	}
	return false
}

func P_up(pts []Date) (a bool) { //pts [][]int

	var b bool = false
	for i := range pts {
		if len(pts) != 1 {
			if pts[i].Value2 < pts[i+1].Value2 && pts[i+1].Date != pts[len(pts)-1].Date {
				continue
			} else if pts[i].Value2 < pts[i+1].Value2 && pts[i+1].Date == pts[len(pts)-1].Date {
				//fmt.Println("Down trend breaks here", pts[i+1])
				b = true
				break
			} else {
				b = false
				break
			}
		} else {
			b = true
			break
		}

	}
	return b
}

func P_down(pts []Date) (a bool) {
	var b bool = false
	for i := range pts {
		if pts[i].Value3 > pts[i+1].Value3 && pts[i+1].Date != pts[len(pts)-1].Date {
			continue
		} else if pts[i].Value3 > pts[i+1].Value3 && pts[i+1].Date == pts[len(pts)-1].Date {
			//fmt.Println("Down trend breaks here", pts[i+1])
			b = true
			break
		} else {
			b = false
			break
		}
	}
	return b
}
func Is_Valid12(category int, n []int) bool {
	switch category {
	case
		n[0],
		n[1],
		n[2],
		n[3],
		n[4],
		n[5],
		n[6],
		n[7],
		n[8],
		n[9],
		n[10],
		n[11],
		n[12],
		n[13],
		n[14],
		n[15],
		n[16],
		n[17],
		n[18],
		n[19],
		n[20],
		n[21],
		n[22],
		n[23],
		n[24],
		n[25],
		n[26],
		n[27],
		n[28],
		n[29],
		n[30],
		n[31],
		n[32],
		n[33],
		n[34],
		n[35],
		n[36],
		n[37],
		n[38],
		n[39],
		n[40],
		n[41],
		n[42],
		n[43],
		n[44],
		n[45],
		n[46],
		n[47],
		n[48],
		n[49],
		n[50],
		n[51],
		n[52],
		n[53],
		n[54],
		n[55],
		n[56],
		n[57],
		n[58],
		n[59],
		n[60],
		n[61],
		n[62],
		n[63],
		n[64],
		n[65],
		n[66],
		n[67],
		n[68],
		n[69],
		n[70],
		n[71],
		n[72],
		n[73],
		n[74],
		n[75],
		n[76],
		n[77],
		n[78],
		n[79],
		n[80],
		n[81],
		n[82],
		n[83],
		n[84],
		n[85],
		n[86],
		n[87],
		n[88],
		n[89],
		n[90],
		n[91],
		n[92],
		n[93],
		n[94],
		n[95],
		n[96],
		n[97],
		n[98],
		n[99]:
		return true
	}
	return false
}

func (l *Date_struct) PatternInverted() ([]int, []int, int, bool) {
	peaks := peaks1(l.Date_struct)
	troughs := troughs1(l.Date_struct)
	// fmt.Println("Peaks:", peaks)
	// fmt.Println("Troughs:", troughs)
	var All_trends_string []string
	var mixed_unsorted []int
	for i := range peaks {
		mixed_unsorted = append(mixed_unsorted, peaks[i])
	}
	for i := range troughs {
		mixed_unsorted = append(mixed_unsorted, troughs[i])
	}
	mixed_sorted := utils.Sort(mixed_unsorted) //mixed values
	var All_trends_index_values []int
	for _, i := range mixed_sorted {
		if P_up(l.Date_struct[i : i+2]) {
			All_trends_string = append(All_trends_string, "Upward")
			All_trends_index_values = append(All_trends_index_values, i)

		} else {
			All_trends_string = append(All_trends_string, "Downward")
			All_trends_index_values = append(All_trends_index_values, i)
		}
	}

	var start []int
	var end []int
	var count int
	var bol bool //just to check if the pattern is semicircle or not,this is boolean that will tell us that.

	for i := range All_trends_string {
		if All_trends_string[i] == "Downward" && i != len(All_trends_string)-1 {
			if All_trends_string[i+1] == "Upward" && i+1 != len(All_trends_string)-1 { //Task:=>check here for first and last point to be in the align manner.
				var a1 []int //end point
				var x int

				for i1 := 0; i1 < 80; i1++ {
					x = l.Date_struct[All_trends_index_values[i+2]].Value3 + i1
					a1 = append(a1, x)
				}
				for i2 := 0; i2 < 80; i2++ {
					x = l.Date_struct[All_trends_index_values[i+2]].Value3 - i2
					a1 = append(a1, x)
				}
				var a2 []int //start point
				var x1 int

				for i1 := 0; i1 < 80; i1++ {
					x1 = l.Date_struct[All_trends_index_values[i]+1].Value2 + i1
					a2 = append(a2, x1)
				}
				for i2 := 0; i2 < 80; i2++ {
					x1 = l.Date_struct[All_trends_index_values[i]+1].Value2 - i2
					a2 = append(a2, x1)
				}

				// if IsValidCategoryu(l.Date_struct[All_trends_index_values[i]+1].Value2, a1) && IsValidCategoryu(l.Date_struct[All_trends_index_values[i+2]-1].Value3, a2) {
				// 	start = append(start, All_trends_index_values[i]+1)
				// 	end = append(end, All_trends_index_values[i+2])
				// 	count++
				// 	bol = true
				// } else if IsValidCategoryu(l.Date_struct[All_trends_index_values[i]-1].Value2, a1) && IsValidCategoryu(l.Date_struct[All_trends_index_values[i+2]+2].Value3, a2) {
				// 	start = append(start, All_trends_index_values[i]-1)
				// 	end = append(end, All_trends_index_values[i+2])
				// 	count++
				// 	bol = true
				// } else if IsValidCategoryu(l.Date_struct[All_trends_index_values[i]+2].Value2, a1) && IsValidCategoryu(l.Date_struct[All_trends_index_values[i+2]-2].Value3, a2) {
				// 	start = append(start, All_trends_index_values[i]+2)
				// 	end = append(end, All_trends_index_values[i+2]-1)
				// 	count++
				// 	bol = true
				// } else if IsValidCategoryu(l.Date_struct[All_trends_index_values[i]].Value2, a1) && IsValidCategoryu(l.Date_struct[All_trends_index_values[i+2]].Value3, a2) {
				// 	start = append(start, All_trends_index_values[i])
				// 	end = append(end, All_trends_index_values[i+2]+1)
				// 	count++
				// 	bol = true
				// } else if IsValidCategoryu(l.Date_struct[All_trends_index_values[i]+3].Value2, a1) && IsValidCategoryu(l.Date_struct[All_trends_index_values[i+2]-3].Value3, a2) {
				// 	start = append(start, All_trends_index_values[i]+3)
				// 	end = append(end, All_trends_index_values[i+2]-2)
				// 	count++
				// 	bol = true
				// } else if IsValidCategoryu(l.Date_struct[All_trends_index_values[i]-3].Value2, a1) && IsValidCategoryu(l.Date_struct[All_trends_index_values[i+2]+3].Value3, a2) {
				// 	start = append(start, All_trends_index_values[i]-3)
				// 	end = append(end, All_trends_index_values[i+2]+4)
				// 	count++
				// 	bol = true
				// } else if IsValidCategoryu(l.Date_struct[All_trends_index_values[i]].Value2, a1) && IsValidCategoryu(l.Date_struct[All_trends_index_values[i+2]].Value3, a2) {
				// 	start = append(start, All_trends_index_values[i]+1)
				// 	end = append(end, All_trends_index_values[i+2]-2)
				// 	count++
				// 	bol = true
				// }
				if Is_Valid11(l.Date_struct[All_trends_index_values[i]+1].Value3, a1) {
					start = append(start, All_trends_index_values[i]+1)
					end = append(end, All_trends_index_values[i+2]) //last U
					count++
					bol = true
				} else if Is_Valid11(l.Date_struct[All_trends_index_values[i]].Value3, a2) {
					start = append(start, All_trends_index_values[i])
					end = append(end, All_trends_index_values[i+2]) //last U])
					count++
					bol = true
				} else if IsValidCategoryu(l.Date_struct[All_trends_index_values[i]+1].Value3, a1) {
					start = append(start, All_trends_index_values[i]+1)
					end = append(end, All_trends_index_values[i+2])
					count++
					bol = true
				} else if Is_Valid11(l.Date_struct[All_trends_index_values[i]+1].Value2, a1) && IsValidCategoryu(l.Date_struct[All_trends_index_values[i+2]].Value3, a1) {
					start = append(start, All_trends_index_values[i]+1)
					end = append(end, All_trends_index_values[i+2])
					count++
					bol = true
				}
			}
		}

	}

	// fmt.Println("Count =", count)
	return start, end, count, bol
}
func IsValidCategoryu(category int, n []int) bool {
	switch category {
	case
		n[0],
		n[1],
		n[2],
		n[3],
		n[4],
		n[5],
		n[6],
		n[7],
		n[8],
		n[9],
		n[10],
		n[11],
		n[12],
		n[13],
		n[14],
		n[15],
		n[16],
		n[17],
		n[18],
		n[19],
		n[20],
		n[21],
		n[22],
		n[23],
		n[24],
		n[25],
		n[26],
		n[27],
		n[28],
		n[29],
		n[30],
		n[31],
		n[32],
		n[33],
		n[34],
		n[35],
		n[36],
		n[37],
		n[38],
		n[39],
		n[40],
		n[41],
		n[42],
		n[43],
		n[44],
		n[45],
		n[46],
		n[47],
		n[48],
		n[49],
		n[50],
		n[51],
		n[52],
		n[53],
		n[54],
		n[55],
		n[56],
		n[57],
		n[58],
		n[59],
		n[60],
		n[61],
		n[62],
		n[63],
		n[64],
		n[65],
		n[66],
		n[67],
		n[68],
		n[69],
		n[70],
		n[71],
		n[72],
		n[73],
		n[74],
		n[75],
		n[76],
		n[77],
		n[78],
		n[79],
		n[80],
		n[81],
		n[82],
		n[83],
		n[84],
		n[85],
		n[86],
		n[87],
		n[88],
		n[89],
		n[90],
		n[91],
		n[92],
		n[93],
		n[94],
		n[95],
		n[96],
		n[97],
		n[98],
		n[99],
		n[100],
		n[101],
		n[102],
		n[103],
		n[104],
		n[105],
		n[106],
		n[107],
		n[108],
		n[109],
		n[110],
		n[111],
		n[112],
		n[113],
		n[114],
		n[115],
		n[116],
		n[117],
		n[118],
		n[119],
		n[120],
		n[121],
		n[122],
		n[123],
		n[124],
		n[125],
		n[126],
		n[127],
		n[128],
		n[129],
		n[130],
		n[131],
		n[132],
		n[133],
		n[134],
		n[135],
		n[136],
		n[137],
		n[138],
		n[139],
		n[140],
		n[141],
		n[142],
		n[143],
		n[144],
		n[145],
		n[146],
		n[147],
		n[148],
		n[149],
		n[150],
		n[151],
		n[152],
		n[153],
		n[154],
		n[155],
		n[156],
		n[157],
		n[158],
		n[159]:
		return true
	}
	return false
}
func peaks1(pts []Date) (a []int) {
	var b []int
	for i := range pts {
		if pts[i].Value2 == pts[0].Value2 {
			continue
		} else if pts[i].Value2 == pts[len(pts)-1].Value2 {
			continue
		} else if pts[i-1].Value2 < pts[i].Value2 && pts[i+1].Value2 < pts[i].Value2 && pts[i-1] != pts[0] && pts[i+1] != pts[len(pts)-1] {

			b = append(b, i)
			continue

		} else {
			continue
		}
	}
	return b
}
func troughs1(pts []Date) (a []int) {
	var b []int
	for i := range pts {
		if pts[i].Value3 == pts[0].Value3 {
			continue
		} else if pts[i].Value3 == pts[len(pts)-1].Value3 {
			continue
		} else if pts[i-1].Value3 > pts[i].Value3 && pts[i+1].Value3 > pts[i].Value3 && pts[i-1] != pts[0] && pts[i+1] != pts[len(pts)-1] {
			var a1 []int
			var x int
			for i1 := 0; i1 < 80; i1++ {
				x = pts[i].Value3 + i1
				a1 = append(a1, x)
			}
			for i2 := 0; i2 < 80; i2++ {
				x = pts[i].Value3 - i2
				a1 = append(a1, x)
			}
			if IsValidCategoryu(pts[i-1].Value3, a1) && pts[i-1].Value3 < pts[i-2].Value3 { //pts[i-1].Value == pts[i].Value || pts[i+1].Value == pts[i].Value
				b = append(b, i)
				continue
			}
			if IsValidCategoryu(pts[i+1].Value3, a1) && pts[i+1].Value3 < pts[i+2].Value3 { //pts[i-1].Value == pts[i].Value || pts[i+1].Value == pts[i].Value
				b = append(b, i)
				continue

			}

		} else {
			continue
		}
	}
	return b

}
func (l *Date_struct) PatternM() ([]int, []int, int, bool) {
	Peaks := Peaks(l.Date_struct)
	// fmt.Println("Peaks =", Peaks)
	Troughs := Troughs(l.Date_struct)
	// fmt.Println("Troughs =", Troughs)
	// l.P = []string{"Upward", "Downward", "Upward", "Downward"}
	var All_trends_string []string
	var mixed_unsorted []int
	for i := range Peaks {
		mixed_unsorted = append(mixed_unsorted, Peaks[i])
	}
	for i := range Troughs {
		mixed_unsorted = append(mixed_unsorted, Troughs[i])
	}
	mixed_sorted := utils.Sort(mixed_unsorted)
	var All_trends_values []int
	for _, i := range mixed_sorted {
		if P_up(l.Date_struct[i : i+2]) {
			All_trends_string = append(All_trends_string, "Upward")
			All_trends_values = append(All_trends_values, i)

		} else {
			All_trends_string = append(All_trends_string, "Downward")
			All_trends_values = append(All_trends_values, i)
		}
	}
	// fmt.Println("All_trends_string =", All_trends_string)
	var bol bool //to check whether the code is of the patter that we expect it to be.
	// fmt.Println("Mix =", mixed_sorted)
	// fmt.Println("Pattern =", All_trends_string)
	var start []int
	var end []int
	var count int
	for i := range All_trends_string {

		if All_trends_string[i] == "Upward" && i != len(All_trends_string)-1 {
			if All_trends_string[i+1] == "Downward" && i+1 != len(All_trends_string)-1 {
				if All_trends_string[i+2] == "Upward" && i+2 != len(All_trends_string)-1 {
					if All_trends_string[i+3] == "Downward" && i+3 != len(All_trends_string)-1 {
						// fmt.Println("Count =", count)
						// fmt.Println("Starting Point =", l.Date_struct[All_trends_values[i]])
						// fmt.Println("Ending Point =", l.Date_struct[All_trends_values[i+3]+1])
						// && IsValidCategoryu(All_trends_values[i+2], a1)
						// && IsValidCategoryu(All_trends_values[i+3], a2)
						var a1 []int //end point
						var x int

						for i1 := 0; i1 < 80; i1++ {
							x = l.Date_struct[All_trends_values[i]].Value3 + i1
							a1 = append(a1, x)
						}
						for i2 := 0; i2 < 80; i2++ {
							x = l.Date_struct[All_trends_values[i]].Value3 - i2
							a1 = append(a1, x)
						}
						var a2 []int //start point
						var x1 int

						for i1 := 0; i1 < 80; i1++ {
							x1 = l.Date_struct[All_trends_values[i+1]].Value2 + i1
							a2 = append(a2, x1)
						}
						for i2 := 0; i2 < 80; i2++ {
							x1 = l.Date_struct[All_trends_values[i+1]].Value2 - i2
							a2 = append(a2, x1)
						}
						if IsValidCategoryu(l.Date_struct[All_trends_values[i]-1].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]+1].Value2, a2) {
							start = append(start, All_trends_values[i]-1)
							end = append(end, All_trends_values[i+3]+2)
							count++
							bol = true

							// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]-2].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]+2].Value2, a2) {
							// 	start = append(start, All_trends_values[i]-2)
							// 	end = append(end, All_trends_values[i+3]+3)
							// 	count++
							// 	bol = true

						} else if IsValidCategoryu(l.Date_struct[All_trends_values[i]-3].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+4]].Value2, a2) {
							start = append(start, All_trends_values[i]-3)
							end = append(end, All_trends_values[i+4]+1)
							count++
							bol = true

						} else if IsValidCategoryu(l.Date_struct[All_trends_values[i]+1].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]+1].Value2, a2) {
							start = append(start, All_trends_values[i]+1)
							end = append(end, All_trends_values[i+3]+2)
							count++
							bol = true
						} else if IsValidCategoryu(l.Date_struct[All_trends_values[i]+1].Value2, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]+1].Value3, a2) {
							start = append(start, All_trends_values[i]+1)
							end = append(end, All_trends_values[i+3]+2)
							count++
							bol = true
						}
						// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]+2].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]+2].Value2, a2) {
						// 	start = append(start, All_trends_values[i]+2)
						// 	end = append(end, All_trends_values[i+3]+3)
						// 	count++
						// 	bol = true
						// }
						// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]+3].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]+3].Value2, a2) {
						// 	start = append(start, All_trends_values[i]+3)
						// 	end = append(end, All_trends_values[i+3]+4)
						// 	count++
						// 	bol = true
						// }
						// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]-1].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]+1].Value2, a2) {
						// 	start = append(start, All_trends_values[i]-1)
						// 	end = append(end, All_trends_values[i+3]+2)
						// 	count++
						// 	bol = true
						// }
						// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]-2].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]+2].Value2, a2) {
						// 	start = append(start, All_trends_values[i]-2)
						// 	end = append(end, All_trends_values[i+3]+3)
						// 	count++
						// 	bol = true

						// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]-3].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]+3].Value2, a2) {
						// 	start = append(start, All_trends_values[i]-3)
						// 	end = append(end, All_trends_values[i+3]+4)
						// 	count++
						// 	bol = true
						// }
						// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]+1].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+4]-1].Value2, a2) {
						// 	start = append(start, All_trends_values[i]+1)
						// 	end = append(end, All_trends_values[i+4]-1)
						// 	count++
						// 	bol = true
						// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]+2].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+4]-2].Value2, a2) {
						// 	start = append(start, All_trends_values[i]+2)
						// 	end = append(end, All_trends_values[i+4]-2)
						// 	count++
						// 	bol = true
						// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]+3].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+4]-3].Value2, a2) {
						// 	start = append(start, All_trends_values[i]+3)
						// 	end = append(end, All_trends_values[i+4]-3)
						// 	count++
						// 	bol = true
						// }
						// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]-1].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+4]+1].Value2, a2) {
						// 	start = append(start, All_trends_values[i]-1)
						// 	end = append(end, All_trends_values[i+4]+1)
						// 	count++
						// 	bol = true
						// }
						// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]-2].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+4]+2].Value2, a2) {
						// 	start = append(start, All_trends_values[i]-2)
						// 	end = append(end, All_trends_values[i+4]+2)
						// 	count++
						// 	bol = true
						// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]-3].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+4]+3].Value2, a2) {
						// 	start = append(start, All_trends_values[i]-3)
						// 	end = append(end, All_trends_values[i+4]+3)
						// 	count++
						// 	bol = true
						// }
						// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]+1].Value2, a2) {
						// 	start = append(start, All_trends_values[i])
						// 	end = append(end, All_trends_values[i+3]+2)
						// 	count++
						// 	bol = true
						// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]+2].Value2, a2) {
						// 	start = append(start, All_trends_values[i])
						// 	end = append(end, All_trends_values[i+3]+3)
						// 	count++
						// 	bol = true
						// }
						// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]+3].Value2, a2) {
						// 	start = append(start, All_trends_values[i])
						// 	end = append(end, All_trends_values[i+3]+4)
						// 	count++
						// 	bol = true
						// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]-1].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]+1].Value2, a2) {
						// 	start = append(start, All_trends_values[i]-1)
						// 	end = append(end, All_trends_values[i+3]+2)
						// 	count++
						// 	bol = true
						// }
						// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]-2].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]+2].Value2, a2) {
						// 	start = append(start, All_trends_values[i]-2)
						// 	end = append(end, All_trends_values[i+3]+3)
						// 	count++
						// 	bol = true
						// } else if IsValidCategoryu(l.Date_struct[All_trends_values[i]-3].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]+3].Value2, a2) {
						// 	start = append(start, All_trends_values[i]-3)
						// 	end = append(end, All_trends_values[i+3]+4)
						// 	count++
						// 	bol = true
						// }
						// start = append(start, All_trends_values[i])
						// end = append(end, All_trends_values[i+3+1])
						// bol = true

					}
				}
			}

		}

	}
	// fmt.Println("Count =", count)
	return start, end, count, bol
}
func (l *Date_struct) PatternW() ([]int, []int, int, bool) {
	Peaks := Peaks(l.Date_struct)
	// fmt.Println("Peaks =", Peaks)
	Troughs := Troughs(l.Date_struct)
	// fmt.Println("Troughs =", Troughs)
	// l.P = []string{"Upward", "Downward", "Upward", "Downward"}
	var All_trends_string []string
	var mixed_unsorted []int
	for i := range Peaks {
		mixed_unsorted = append(mixed_unsorted, Peaks[i])
	}
	for i := range Troughs {
		mixed_unsorted = append(mixed_unsorted, Troughs[i])
	}
	mixed_sorted := utils.Sort(mixed_unsorted)
	var All_trends_values []int
	for _, i := range mixed_sorted {
		if P_up(l.Date_struct[i : i+2]) {
			All_trends_string = append(All_trends_string, "Upward")
			All_trends_values = append(All_trends_values, i)

		} else {
			All_trends_string = append(All_trends_string, "Downward")
			All_trends_values = append(All_trends_values, i)
		}
	}
	// fmt.Println("All_trends_string =", All_trends_string)
	var bol bool //to check whether the code is of the patter that we expect it to be.
	// fmt.Println("Mix =", mixed_sorted)
	// fmt.Println("Pattern =", All_trends_string)
	var start []int
	var end []int
	var count int
	for i := range All_trends_string {

		if All_trends_string[i] == "Downward" && i != len(All_trends_string)-1 {
			if All_trends_string[i+1] == "Upward" && i+1 != len(All_trends_string)-1 {
				if All_trends_string[i+2] == "Downward" && i+2 != len(All_trends_string)-1 {
					if All_trends_string[i+3] == "Upward" && i+3 != len(All_trends_string)-1 {
						// fmt.Println("Count =", count)
						// fmt.Println("Starting Point =", l.Date_struct[All_trends_values[i]])
						// fmt.Println("Ending Point =", l.Date_struct[All_trends_values[i+3]+1])
						// && IsValidCategoryu(All_trends_values[i+2], a1)
						// && IsValidCategoryu(All_trends_values[i+3], a2)
						var a1 []int //end point
						var x int

						for i1 := 0; i1 < 80; i1++ {
							x = l.Date_struct[All_trends_values[i]].Value3 + i1
							a1 = append(a1, x)
						}
						for i2 := 0; i2 < 80; i2++ {
							x = l.Date_struct[All_trends_values[i]].Value3 - i2
							a1 = append(a1, x)
						}
						var a2 []int //start point
						var x1 int

						for i1 := 0; i1 < 80; i1++ {
							x1 = l.Date_struct[All_trends_values[i+1]].Value2 + i1
							a2 = append(a2, x1)
						}
						for i2 := 0; i2 < 80; i2++ {
							x1 = l.Date_struct[All_trends_values[i+1]].Value2 - i2
							a2 = append(a2, x1)
						}
						if IsValidCategoryu(l.Date_struct[All_trends_values[i+2]].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]].Value2, a2) {
							start = append(start, All_trends_values[i])
							end = append(end, All_trends_values[i+3]+2)
							count++
							bol = true
						} else if IsValidCategoryu(l.Date_struct[All_trends_values[i+1]].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+3]].Value2, a2) {
							start = append(start, All_trends_values[i]-2)
							end = append(end, All_trends_values[i+3+1])
							count++
							bol = true
						} else if IsValidCategoryu(l.Date_struct[All_trends_values[i]-2].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+2]].Value2, a2) {
							start = append(start, All_trends_values[i]-2)
							end = append(end, All_trends_values[i+3+1])
							count++
							bol = true
						} else if IsValidCategoryu(l.Date_struct[All_trends_values[i]].Value3, a1) && IsValidCategoryu(l.Date_struct[All_trends_values[i+4]].Value2, a2) {
							start = append(start, All_trends_values[i]+2)
							end = append(end, All_trends_values[i+3+1]+4)
							count++
							bol = true
						}
						// start = append(start, All_trends_values[i])
						// end = append(end, All_trends_values[i+3+1])
						// bol = true

					}
				}
			}

		}

	}
	// fmt.Println("Count =", count)
	return start, end, count, bol
}
func Peaks(pts []Date) (a []int) {
	var b []int
	for i := range pts {
		if pts[i].Value2 == pts[0].Value2 {
			continue
		} else if pts[i].Value2 == pts[len(pts)-1].Value2 {
			continue
		} else if pts[i-1].Value2 < pts[i].Value2 && pts[i+1].Value2 < pts[i].Value2 && pts[i+1] != pts[len(pts)-1] {
			b = append(b, i)
			continue

		}
	}
	return b
}

func Troughs(pts []Date) (a []int) {
	var b []int
	for i := range pts {
		if pts[i].Value3 == pts[0].Value3 {
			continue
		} else if pts[i].Value3 == pts[len(pts)-1].Value3 {
			continue
		} else if pts[i-1].Value3 > pts[i].Value3 && pts[i+1].Value3 > pts[i].Value3 && pts[i-1] != pts[0] && pts[i+1] != pts[len(pts)-1] {

			b = append(b, i)
			continue

		}

	}
	return b

}
