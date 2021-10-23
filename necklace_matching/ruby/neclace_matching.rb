def b_post(start, len, idx)
  # lda
  # l | 2+1=3 | 3-3=0
  # d | 2+2=4 | 4-3=1

  pos = start + idx
  pos < len ? pos : pos - len
end

def main
  a = 'aabaaaaabaab'
  b = 'aabaabaabaaa'

  return if a.length != b.length

  b_len = b.length
  b_start_idxs = (0...b_len).select { |i| b[i] == a[0] }

  b_start_idxs.each do |b_start_idx|
    find_match = (1...a.length).map { |i| b_post(b_start_idx, b_len, i) }.map.with_index { |j, i| a[i + 1] == b[j] }
    count_match = find_match.select { |i| i == true }

    if find_match.length == a.length - 1 && count_match.length == find_match.length
      puts('true')
      return
    end
  end

  puts('false')
end

main