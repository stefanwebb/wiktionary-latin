import time, re
import pickle
import os.path
import xml.etree.ElementTree as et

class Timer:
  def __enter__(self):
    self.start = time.clock()
    return self

  def __exit__(self, *args):
    self.end = time.clock()
    self.interval = self.end - self.start

def load_wiktionary(filename='./wiktionary-latin-english.xml'):
  #if not os.path.isfile(pickle_name):
  print('Parsing Wiktionary XML...')
  with Timer() as t:
    tree = et.parse(filename)
  print('Took {} secs'.format(t.interval))

  return tree

# Print out some stuff from the XML parse tree
tree = load_wiktionary()
root = tree.getroot()

titles = set()
for c in root:
  titles.add(c.attrib['title'])

with open('la-en-candidates.txt', 'w', encoding='utf-8') as f:
  for t in sorted(titles):
    print(t, file=f)