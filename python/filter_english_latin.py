# Filtering the Wiktionary entries that have Latin and English definition to remove uncommon words
# I compare to a list of the 1/3 million most frequent words due to Peter Norvig
words = {}
with open('count_1w.txt', 'r') as f:
  lines = f.readlines()

for l in lines:
  s = l.split('\t')
  #print(l)
  #print(s)
  words[s[0].strip()] = int(s[1].strip())

  #print(words)
  #raise Exception()

with open('la-en-candidates.txt', 'r') as f:
  lines = f.readlines()

with open('la-en-filtered.txt', 'w', encoding='utf-8') as f:
  for l in lines:
    w = l.strip()
    # In top 100,000 words
    #if w in words and words[w] >= 99133:
    # Top 20,000 words
    #if w in words and words[w] >= 1602219:
    # Top 60,000 words
    if w in words and words[w] >= 243069:
      print(w, file=f)
    

#print(len(words))
#print('Done!')