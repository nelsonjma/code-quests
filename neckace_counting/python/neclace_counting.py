import math
import functools as ft

# this scenario has problems with big numbers x/y had to be replaced by x//y in some operations
# and math.pow(x, y) does not work, you have to use x**y

def is_prime(num):
  # all numbers less than 2 are not prime

  if num <= 1:
    return False

  # see if num is divisible by any number
  for i in range(2, num):
    if num % i == 0:
      return False

  return True


def primes_that_divide_num(num: int) -> list:
  return [i for i in range(1, num+1) if is_prime(i) and (num/i)-int(num/i) == 0]


def calc_phi(num: int) -> int:
  """
    phi(12) = 12 * ((2-1)*(3-1)) / (2*3) = 12 * 2 / 6 = 4
    phi(a) = a * b / c
  """
  a = num

  ptd = primes_that_divide_num(num)

  if ptd == []:
    return 1

  # ((2-1)*(3-1))
  b = 1
  for p in ptd:
    b *= (p-1)

  c = ft.reduce(lambda a,b: a*b, ptd)

  # phi(a) = a * b / c
  return a * b // c


def factor_num_into_positive_ints(num: int) -> tuple:
  # factor 10 into two positive integers are 1x10, 2x5, 5x2, and 10x1
  # return [(10,1), (5,2), (2,5) (1,10)]
  return [(i, int(num/i)) for i in range(1, num+1) if (num/i) - int(num/i) == 0]

def calc_neclaces(k: int, n: int) -> int:
  # necklaces(k, n) = 1/n * Sum of (phi(a) k^b)
  # 1/10 (phi(1) 3^10 + phi(2) 3^5 + phi(5) 3^2 + phi(10) 3^1)
  # 1/10 (1 * 59049 + 1 * 243 + 4 * 9 + 4 * 3)
  # necklaces(k, n) = x * y

  # 1 / 10
  # x = 1/n this does not work in big numbers

  # (phi(1) 3^10 + phi(2) 3^5 + phi(5) 3^2 + phi(10) 3^1)
  y = 0
  for f in factor_num_into_positive_ints(n):
    y += calc_phi(f[1]) * k**f[0]
    # print(f'{f[1]}/{f[0]} {calc_phi(f[1])} * {k**f[0]}  ')

    # math.pow(k,f[0] this does not work for big number, for that you have to use k**f[0]
  
  # 1/n * ...
  return y//n


print(f'calc_neclaces(2, 12)          {calc_neclaces(2, 12)} => 352')
print(f'calc_neclaces(3, 7)           {calc_neclaces(3, 7)} => 315')
print(f'calc_neclaces(9, 4)           {calc_neclaces(9, 4)} => 1665')
print(f'calc_neclaces(21, 3)          {calc_neclaces(21, 3)} => 3101')
print(f'calc_neclaces(99, 2)          {calc_neclaces(99, 2)} => 4950')
print(f'calc_neclaces(12345678910, 3) {calc_neclaces(12345678910, 3)} => 627225458787209496560873442940')
