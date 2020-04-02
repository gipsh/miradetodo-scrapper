require 'base64'



s = "aHR0c!ArEovL29rLnJ1L3ZpZGVvLzE2MTg5O!ArEExNjE1N!ArEUN"


puts "original string [#{s}]"
t = s.split("!Ar")

for i in 1..t.size-1 
   puts "#{i}. #{t[i]}"
   t[i][0] = (t[i][0].ord - 1).chr
end 

puts t.join
puts "Download link is: #{Base64.decode64(t.join)}"

