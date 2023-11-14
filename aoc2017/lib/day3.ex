defmodule Day3 do
  def puzzle_input do
    277678
  end

  def part1(num) do
    circle = :math.ceil(:math.sqrt(num)) / 2
    circle_zero = :math.pow((circle * 2 - 1), 2)
    centers = Enum.map([1, 3, 5, 7], fn x -> (circle * x) + circle_zero end)
    distance = circle + Enum.min(Enum.map(centers, fn center -> abs(num - center) end))
    trunc(distance)
  end

  def get_next_coords(x, y) do
    case {x, y} do
      {x, y} when (x == 0 and y == 0) ->
        {1, 0}
      {x, y} when (y > -x and x > y) ->
        {x, y + 1}
      {x, y} when (y > -x and y >= x) ->
        {x - 1, y}
      {x, y} when (y <= -x and x < y) ->
        {x, y - 1}
      {x, y} when (y <= -x and x >= y) ->
        {x + 1, y}
    end
  end
  
  # list of x+i, y+j
  def create_neighbor_map(x, y) do
    nmap = Enum.map(-1..1, fn i ->
      Enum.map(-1..1, fn j ->
        "#{x + i},#{y + j}"
      end)
    end)

    Enum.reduce(nmap, [], fn row, acc ->
      acc ++ row
    end)
  end

  def get_neighbor_values(map, x, y) do
    neighbor_map = create_neighbor_map(x, y)
    Enum.map(neighbor_map, fn coord ->
      Map.get(map, coord, 0)
    end)
  end

  def part2(num, x, y, spiral) do
    if Map.get(spiral, "#{x},#{y}", 0) > num do
      Map.get(spiral, "#{x},#{y}")
    else
      {x1, y1} = get_next_coords(x, y)
      new_spiral = Map.put(spiral, "#{x1},#{y1}",
        Enum.sum(
          get_neighbor_values(spiral, x1, y1)
        )
      )
      part2(num, x1, y1, new_spiral)
    end
  end

  def run do
    puzzle_input = puzzle_input()
    IO.puts " %-----% Day 3 %-----% "
    IO.puts "  Part 1: #{part1(puzzle_input)}"
    IO.puts "  Part 2: #{part2(puzzle_input, 0, 0, %{"0,0" => 1})}"
    IO.puts " %-------------------% "
  end
end
