from string import ascii_uppercase

gematria = dict(
    zip(
        list(ascii_uppercase),
        range(1,26+1)
    )
)

sator = [
    "SATOR",
    "AREPO",
    "TENET",
    "OPERA",
    "ROTAS"
]

matrix = []

for word in sator:
    cypher = [ ]
    for letter in word:
        num = gematria[letter]
        if num < 10:
            cypher.append(f" {num}")
        else:
            cypher.append(str(num))
    matrix.append(cypher)

for row in matrix:
    print("   ".join(row))