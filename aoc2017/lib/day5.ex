defmodule Day5 do
  def puzzle_input do
    case File.read("inputs/day5.txt") do
      {:ok, contents} -> String.split(contents, "\n") |> Enum.filter(&(&1 != "")) |> Enum.map(&String.to_integer/1)
      {:error, reason} -> reason
    end
  end

  def part1(instruction_list, curr_instruction, steps_taken) do
    if curr_instruction >= length(instruction_list) do
      steps_taken
    else
      next_instruction = curr_instruction + Enum.at(instruction_list, curr_instruction)
      new_instruction_list = Enum.map(Enum.with_index(instruction_list), fn {instruction, idx} ->
        if idx == curr_instruction do
          instruction + 1
        else
          instruction
        end
      end)
      part1(new_instruction_list, next_instruction, steps_taken + 1)
    end
  end

  def part2(instruction_list, curr_instruction, steps_taken) do
    if curr_instruction >= length(instruction_list) do
      steps_taken
    else
      next_instruction = curr_instruction + Enum.at(instruction_list, curr_instruction)
      new_instruction_list = Enum.map(Enum.with_index(instruction_list), fn {instruction, idx} ->
        if idx == curr_instruction && instruction >= 3 do
          instruction - 1
        else
          if idx == curr_instruction do
            instruction + 1
          else
            instruction
          end
        end
      end)
      part2(new_instruction_list, next_instruction, steps_taken + 1)
    end
  end

  def run do
    puzzle_input = puzzle_input()
    IO.puts " %-----% Day 5 %-----% "
    IO.puts "  Part 1: #{part1(puzzle_input, 0, 0)}"
    IO.puts "  Part 2: #{part2(puzzle_input, 0, 0)}"
    IO.puts " %-------------------% "
  end
end


