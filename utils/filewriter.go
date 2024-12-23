package utils

import (
	"bufio"
	"fmt"
	"os"
)

func Write2DSliceToFile(filePath string, grid [][]int) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, row := range grid {
		for j, val := range row {
			if j == len(row)-1 {
				_, err = writer.WriteString(fmt.Sprintf("%d", val))
			} else {
				_, err = writer.WriteString(fmt.Sprintf("%d ", val))
			}
			if err != nil {
				return fmt.Errorf("failed to write: %w", err)
			}
		}
		_, err = writer.WriteString("\n")
		if err != nil {
			return fmt.Errorf("failed to write newline to file: %w", err)
		}
	}
	_, err = writer.WriteString("------------------------------------------------------------")
	if err != nil {
		return fmt.Errorf("failed to write - to file: %w", err)
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush writer: %w", err)
	}

	return nil
}

// modified write to file for Day 14
func Write2DSliceToFileDay14(filePath string, grid [][]int) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, row := range grid {
		for j, val := range row {
			var symbol string
			if val == 0 {
				symbol = "."
			} else {
				symbol = "+"
			}

			if j == len(row)-1 {
				_, err = writer.WriteString(symbol)
			} else {
				_, err = writer.WriteString(symbol)
			}

			if err != nil {
				return fmt.Errorf("failed to write to file: %w", err)
			}
		}
		_, err = writer.WriteString("\n")
		if err != nil {
			return fmt.Errorf("failed to write newline to file: %w", err)
		}
	}
	_, err = writer.WriteString("\n\n\n\n")
	if err != nil {
		return fmt.Errorf("failed to write - to file: %w", err)
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush writer: %w", err)
	}

	return nil
}

func AppendIntegerToFile(filename string, number int) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%d\n", number))
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
