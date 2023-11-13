defmodule Aoc2017 do
  use Application

  def start(_type, _args) do
    Day1.run()

    {:ok, self()}
  end
end
