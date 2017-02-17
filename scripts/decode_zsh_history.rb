#!/usr/bin/env ruby

(ARGV.empty? ? STDIN : File.read(ARGV[0])).each_line do |line|
  chars = []
  meta = false
  bytes = line.bytes
  x83 = bytes.find_index(0x83)

  if x83
    chars = bytes.slice!(0, x83)
  else
    puts line
    next
  end

  bytes.each do |char|
    if char == 0x83
      meta = true
    elsif meta
      chars << (char ^ 0x20)
      meta = false
    else
      chars << char
    end
  end

  puts chars.pack('c*')
end
