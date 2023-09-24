defmodule Hello do
  def say_hello do
    IO.puts("Hello gengs!!")
  end

  def add(x, n) do
    x + n
  end
end

Hello.say_hello()
result = Hello.add(5, 50)
IO.puts("sum of var result is #{result}")

# Creating anonymous function
sum = fn x, n -> x + n end
#other way to create anonym func
sum = &(&1+ &2)

res_sum =sum.(10, 10)
IO.puts(res_sum)

res_sum_times = fn res, times -> res * times end
res_sum_times = &(&1 * &2)

IO.puts(res_sum_times.(res_sum, 5))
#create a map and list
ex_list = ["kiw", "kiw", 1,2,3,4]#->tipe data ini tetap menyimpan int dan string, namum lompat ke io.puts
ex_list2 = [1, 23, 4, 5, 6] #-> ini tidak akan bisa di ioputs harus string joind dulu
IO.puts(ex_list)#ketika di io.puts ini akan dikonvert dan di merge menjadi satu string, jika list ada int tidak akan tereksekusi
my_map = %{Name: "John", Age: 30, Address: "New York"}
my_map = Map.put(my_map, :PhoneNumber, "921021743")#append map

#conditional
if true do
  IO.puts "yoo guys!!"
end
unless false do
  IO.puts "your condition is false broww!"
end
