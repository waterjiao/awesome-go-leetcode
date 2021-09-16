package _051_N_Queens

import "strings"

var results [][]string

func solveNQueens(n int) [][]string {
	results = make([][]string, 0)
	//board := initBoard(n)
	board := make([][]string, n)
	for i := 0; i < n; i++{
		board[i] = make([]string, n)
	}
	for i := 0; i < n; i++{
		for j := 0; j<n;j++{
			board[i][j] = "."
		}
	}

	backtrack(board, 0)
	return results
}

func initBoard(n int) [][]string {
	board := make([][]string, 0)
	cow := make([]string, 0)
	for i := 0; i < n; i++ {
		cow = append(cow, ".")
	}
	for i := 0; i < n; i++ {
		board = append(board, cow)
	}
	return board
}

func backtrack(board [][]string, cow int) {
	if cow == len(board) {
		cows := make([]string, len(board))
		//for _, c := range board {
		//	s := strings.Join(c, "")
		//	cows = append(cows, s)
		//}
		for i := 0; i<len(board);i++{
			cows[i] = strings.Join(board[i],"")
		}
		results = append(results, cows)
		return
	}
	col := len(board) - 1
	for i := 0; i <= col; i++ {
		if !isVaild(cow, i, board) {
			continue
		}
		board[cow][i] = "Q"
		backtrack(board, cow+1)
		board[cow][i] = "."
	}
}

func isVaild(row, col int, board [][]string) bool {
	//l, r := col, col
	//for cow-1 >= 0 && l >= 1 && r <= len(board)-2 {
	//	if board[cow-1][col] == "Q" {
	//		return false
	//	}
	//	if board[cow-1][l-1] == "Q" {
	//		return false
	//	}
	//	if board[cow-1][r+1] == "Q" {
	//		return false
	//	}
	//	l--
	//	r++
	//	cow--
	//}
	//return true
	n := len(board)
	for i:=0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	for i := 0; i < n; i++{
		if board[row][i] == "Q" {
			return false
		}
	}

	for i ,j := row, col; i >= 0 && j >=0 ; i, j = i - 1, j- 1{
		if board[i][j] == "Q"{
			return false
		}
	}
	for i, j := row, col; i >=0 && j < n; i,j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}

//func solveNQueens(n int) [][]string {
//	results = make([][]string, 0)
//	result := make([]string, 0)
//	flag := strings.Repeat(".", n)
//	for i:=0; i< n; i++ {
//		result = append(result, flag)
//	}
//	backtrack(result, 0)
//	return results
//}
//
//func backtrack(result []string, row int) {
//	if row == len(result) {
//		ans := make([]string, len(result))
//		copy(ans, result)
//		results = append(results, ans)
//	}
//	for i := range result {
//		if !isValid(i, row, result) {
//			continue
//		}
//		result[i] = "Q"
//		backtrack(result, row+1)
//		result[i] = "."
//	}
//}
//
//func isValid(index, row int, result []string) bool {
//	for row-1 >= 0 && index >= 0 && index < len(result) {
//		if result[row-1] == "Q" {
//
//		}
//	}
//	return false
//}
