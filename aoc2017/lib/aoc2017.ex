defmodule Aoc2017 do
  use Application

  def start(_type, _args) do
    Day1.run()
    Day2.run()
    Day3.run()
    Day4.run()
    # takes a while to run
    # Day5.run()
    Day6.run()

    {:ok, self()}
  end
end
