
from functools import reduce
from itertools import permutations


def load_and_clean_data(filename):
  # process and clean each line
  lines_cells = []

  with open(filename, 'r') as f:
    file_lines = f.readlines()

    for line in file_lines:
      line_cells = (i for i in line.split(' ') if i.strip() != '')
      rm_end_char_line_cells = (int(i.replace('\n', '').strip()) for i in line_cells)
      to_int_line_cells = [int(i) for i in rm_end_char_line_cells]

      lines_cells.append(to_int_line_cells)
  
  return lines_cells


def find_min_current_permutation(lines_cells, x_positions):
  min_cell_from_lines = {}
  used_positions = []
  # number of lines
  y_len = len(lines_cells)

  for x in x_positions:
    last_y = -1

    for y in range(y_len):
      # fetch last stored value
      stored_value = min_cell_from_lines.get(x, 0)
      # get cell from line
      cell = lines_cells[y][x]
      # current cell is less than stored value
      if (cell < stored_value or stored_value == 0) and y not in used_positions:
        min_cell_from_lines[x] = cell
        last_y = y

    # save last position
    used_positions.append(last_y)

  min_cell_values = (min_cell_from_lines[k] for k in min_cell_from_lines)

  return reduce(lambda a, b: a+b, min_cell_values)


lines_cells = load_and_clean_data('../code/code-quests/matrix-sum/5x5.txt')

# fetch x,y length
x_len = len(lines_cells[0])
result_elements = []

# permutaions will generate all possibilities (1, 2, 3) - (1, 3, 2) - (2, 1, 3) - ...
count_perm = 0
for permutation in permutations(list(range(x_len))):
  result_elements.append(find_min_current_permutation(lines_cells, permutation))
  count_perm += 1
  if count_perm % 1000 == 1:
    print(f"{count_perm}")

print(min(result_elements))
