package main

func SplitNull(arr []any) [][]any {
	var arr1 [][]any
	var temp []any
	for i, val := range arr {
		if i == 0 && val == nil {
			continue
		}

		if i == len(arr)-1 {
			if len(temp) > 1 {
				if val != nil {
					temp = append(temp, val)
				}
				arr1 = append(arr1, temp)
			}
			continue
		}

		if val == nil {
			arr1 = append(arr1, temp)
			temp = []any{}
			continue
		}

		temp = append(temp, val)
	}
	return arr1
}

func main() {
}
