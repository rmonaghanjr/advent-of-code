defmodule Day1 do
  
  def puzzle_input do
    case File.read("inputs/day1.txt") do
      {:ok, contents} -> String.replace(contents, "\n", "") |> String.graphemes() |> Enum.map(&String.to_integer/1)
      {:error, reason} -> throw(reason)
    end
  end

  def part1(nums) do
    cyclic_nums = nums ++ [Enum.at(nums, 0)]
    Enum.reduce(Enum.with_index(cyclic_nums), 0, fn {curr_num, idx}, acc ->
      if curr_num == Enum.at(cyclic_nums, idx + 1) do
        acc + curr_num
      else
        acc
      end
    end)
  end
  
  def part2(nums) do
    cyclic_nums = Stream.cycle(nums)
    Enum.reduce(Enum.with_index(nums), 0, fn {curr_num, idx}, acc ->
      if curr_num == Enum.at(cyclic_nums, idx + trunc(Enum.count(nums) / 2)) do
        acc + curr_num
      else
        acc
      end
    end)
  end

  def run do
    puzzle_input = puzzle_input()
    IO.puts " %-----% Day 1 %-----% "
    IO.puts "  Part 1: #{part1(puzzle_input)}"
    IO.puts "  Part 2: #{part2(puzzle_input)}"
    IO.puts " %-------------------% "
  end

end
