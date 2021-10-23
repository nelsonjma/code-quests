def next_index(current_pos: int, letters_len: int) -> int:
  return 0 if current_pos >= letters_len - 1 else current_pos + 1


def check_next_letter(a: str, b: str, a_pos: int, b_pos) -> bool:
  if a_pos >= len(a):
    return True

  if a[a_pos] == b[b_pos]:
    b_pos = next_index(b_pos, len(b))
    return check_next_letter(a, b, a_pos+1, b_pos)
  
  return False


def main():
  a: str = 'ald'
  b: str = 'lda'

  a_pos = 0

  # if size does not match leave
  if len(a) != len(b):
    print(False)
    return

  # a ="" b=""
  if len(a) == len(b) and len(a) == 0:
    print(True)
    return

  # create list of possible start points
  b_pos_idxs = [i for i in range(len(b)) if b[i] == a[a_pos]]
  
  # no match found
  if b_pos_idxs == []:
    return

  for b_pos in b_pos_idxs:
    status = check_next_letter(a, b, a_pos, b_pos)
    if status:
      print(status)
      return
  print(False)


if __name__ == "__main__":
  main()
