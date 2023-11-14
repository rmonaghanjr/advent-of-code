defmodule Day4 do
  def puzzle_input do
    case File.read("inputs/day4.txt") do
      {:ok, contents} -> String.split(contents, "\n") |> Enum.filter(&(&1 != ""))
      {:error, reason} -> reason
    end
  end

  def part1(passphrases) do
    Enum.filter(passphrases, fn passphrase ->
      words = String.split(passphrase, " ")
      length(words) == length(Enum.uniq(words))
    end) |> length()
  end

  def part2(passphrases) do
    Enum.filter(passphrases, fn passphrase ->
      words = String.split(passphrase, " ")
      length(words) == length(Enum.uniq(words)) &&
      length(words) == length(Enum.uniq(Enum.map(words, &String.codepoints(&1) |> Enum.sort)))
    end) |> length()
  end

  def run do
    puzzle_input = puzzle_input()
    IO.puts " %-----% Day 4 %-----% "
    IO.puts "  Part 1: #{part1(puzzle_input)}"
    IO.puts "  Part 2: #{part2(puzzle_input)}"
    IO.puts " %-------------------% "
  end
end

