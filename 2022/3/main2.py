import fileinput

alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

i = 0
l = []
sum = 0
for line in fileinput.input():
	line = line.strip()
	l += [set(line)]
	i += 1
	if i == 3:
		c = l[0].intersection(l[1], l[2])
		sum += alpha.find(list(c)[0]) + 1
		l = []
		i = 0

print(sum)
