defmodule Day2 do
  def puzzle_input do
    rows = case File.read("inputs/day2.txt") do
      # 2d array of integers
      {:ok, contents} -> String.split(contents, "\n")
      {:error, reason} -> throw(reason)
    end
    
    Enum.map(rows, fn row ->
      if row != "" do
        String.split(row, " ") |> Enum.map(&String.to_integer/1)
      else
        [0]
      end
    end)
  end

  def part1(spreadsheet) do
    Enum.reduce(spreadsheet, 0, fn row, acc ->
      acc + (Enum.max(row) - Enum.min(row))
    end)
  end

  def part2(spreadsheet) do
    Enum.reduce(spreadsheet, 0, fn row, acc ->
      acc + (Enum.reduce(row, 0, fn num, acc ->
        if Enum.any?(row, fn other_num -> num != other_num && rem(num, other_num) == 0 end) do
          acc + div(num, Enum.find(row, fn other_num -> num != other_num && rem(num, other_num) == 0 end))
        else
          acc
        end
      end))
    end)
  end

  def run do
    puzzle_input = puzzle_input()
    IO.puts " %-----% Day 2 %-----% "
    IO.puts "  Part 1: #{part1(puzzle_input)}"
    IO.puts "  Part 2: #{part2(puzzle_input)}"
    IO.puts " %-------------------% "
  end
end
