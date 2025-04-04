package main

import (
	"errors"
	"strings"
)

const rows, columns, empty = 9, 9, 0

var (
	ErrBounds           = errors.New("out of bounds")
	ErrDigit            = errors.New("invalid digit")
	ErrPlaced           = errors.New("already placed")
	ErrPlacedHorizontal = errors.New("already placed horizontal")
	ErrPlacedVertical   = errors.New("already placed vertical")
	ErrFixedDigit       = errors.New("fixed digit")
)

type SudokuError []error

func (err SudokuError) Error() string {
	var errors []string
	for _, err := range err {
		errors = append(errors, err.Error())
	}
	return strings.Join(errors, "\n")
}

type Cell struct {
	digit int8
	fixed bool
}

type Grid [rows][columns]Cell

func (grid *Grid) outBounds(row, column int8) bool {
	if row < 0 || row >= rows {
		return true
	}

	if column < 0 || column >= rows {
		return true
	}

	return false
}

func (grid *Grid) isFixed(row, column int8) bool {
	return grid[row][column].fixed
}

func (grid *Grid) place(row, column, digit int8) SudokuError {
	var errs SudokuError

	if grid.outBounds(row, column) {
		errs = append(errs, ErrBounds)
	}

	if grid.placed(row, column) {
		errs = append(errs, ErrPlaced)
	}

	if grid.isFixed(row, column) {
		errs = append(errs, ErrPlacedHorizontal)
	}

	if grid.inRow(row, digit) {
		errs = append(errs, ErrPlacedHorizontal)
	}

	if grid.inColum(column, digit) {
		errs = append(errs, ErrPlacedVertical)
	}

	if grid.inRegion(row, column, digit) {
		errs = append(errs, ErrPlaced)
	}

	if notValid(digit) {
		errs = append(errs, ErrDigit)
	}

	grid[row][column].digit = digit

	return errs
}

func (grid *Grid) placed(row, column int8) bool {
	if grid[row][column].digit != 0 {
		return true
	}
	return false
}

func (grid *Grid) inRow(row, digit int8) bool {
	for column := 0; column < columns; column++ {
		if grid[row][column].digit == digit {
			return true
		}
	}
	return false
}

func (grid *Grid) inColum(column, digit int8) bool {
	for row := 0; row < rows; row++ {
		if grid[row][column].digit == digit {
			return true
		}
	}
	return false
}

func (grid *Grid) inRegion(row, column, digit int8) bool {
	startRow, startColumn := row/3*3, column/3*3
	for r := startRow; r < startRow+3; r++ {
		for c := startColumn; c < startColumn+3; c++ {
			if grid[r][c].digit == digit {
				return true
			}
		}
	}
	return false
}

func newSudoku(digits [rows][columns]int8) *Grid {
	var grid Grid
	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			grid[row][column] = Cell{digits[row][column], digits[row][column] != empty}
		}
	}
	return &grid
}

func notValid(digit int8) bool {
	return digit < 0 || digit > 9
}
