"""
    Script to cleanup input files for mchains.iofile
"""
fname = "la-princesse-de-cleve-la-fayette.txt"
with open(fname, "rb") as f:
    text = f.read()


import unidecode
unaccented_string = unidecode.unidecode(text.decode("utf-8"))


with open(fname, "w") as f:
    f.write(unaccented_string)

with open(fname, "r") as f:
    l = f.readlines()

newl = list()
for lines in l:
    for i in lines :
        if i.lower() in """abcdefghijklmnopqrstuvwxyz 123456789.?:,;!'<>/\n""":
            newl.append(i)

# with open(fname, "w") as f:
#     f.writelines(newl)
