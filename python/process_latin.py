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

def load_wiktionary(filename='./wiktionary-latin.xml'):
  #if not os.path.isfile(pickle_name):
  print('Parsing Wiktionary XML...')
  with Timer() as t:
    tree = et.parse(filename)
  print('Took {} secs'.format(t.interval))

  return tree

# Print out some stuff from the XML parse tree
tree = load_wiktionary()
root = tree.getroot()
third_level = {}
fourth_level = {}
meta_entries = {}

# Regular expression for third and fourth level headers
section_str = re.compile('(^===[^=]+?===$)', re.MULTILINE)
fourth_str = re.compile('(^====[^=]+?====$)', re.MULTILINE)

# Loop over extracted Latin entries
for c in root:
  # Skip all meta-entries, like Reconstruction:... or Appendix:...
  if c.text is not None and not ':' in c.attrib['title']:
    # Split on third-level header and loop over them
    sections = section_str.split(c.text)
    for idx, s in enumerate(sections):
      # If this a third level header...
      if section_str.match(s) is not None:
        h = s[3:-3]

        # DEBUG: Some contributed mistakenly put a space between === and header
        #if h[0] == ' ':
          #print(h)
          #print(c.attrib['title'])
          #raise Exception()

        # Skip reconstructed entries
        if h in third_level:
          third_level[h] += 1
        else:
          third_level[h] = 1

      # Otherwise split and loop on fourth level headers
      else:
        fourth_sections = fourth_str.split(c.text)
        for idx4, s4 in enumerate(fourth_sections):
          if fourth_str.match(s4) is not None:
            h4 = s[4:-4]
            if h4 in fourth_level:
              fourth_level[h] += 1
            else:
              fourth_level[h] = 1

        #if h == 'Latin':
        #  #print(c.text)
        #  print(c.attrib['title'])
        #if h == 'Pronuncation' or h == 'Prnonuciation':
        #  print('Incorrect pronunciation header:', c.attrib['title'])
        #print(idx, s.strip())
        
        """if h == 'Alternate forms':
          print('Alternate forms', c.attrib['title'])
        if h == 'Alternative form':
          print('Alternative form', c.attrib['title'])"""

        #if h == 'Letter':
        #  print('Letter', c.attrib['title'])

        #if h == 'Etymology 5':
        #  print('Etymology 5', c.attrib['title'])

        #if h == 'Prepositional phrase':
        #  print('Prepositional phrase', c.attrib['title'])

        #if h == 'Alternative forms':
        #  print('Alternative forms', c.attrib['title'])
        #  raise Exception()

    #print(len(sections))

  # DEBUG
  if ':' in c.attrib['title']:
    h = c.attrib['title'].split(':')[0]
    #print(h)
    #raise Exception()
    if h in meta_entries:
      meta_entries[h] += 1
    else:
      meta_entries[h] = 1
  #raise Exception()

# {'Appendix': 4, 'Wiktionary': 14, 'Help': 1, 'Reconstruction': 570, 'Thesaurus': 1, 'Module': 1}
#print('meta-entries', meta_entries)

# Headers
#print('third level headers')
#for key, value in sorted(third_level.items(), key=lambda x: x[0]): 
#  print("{} : {}".format(key, value))

# 'Abbreviation'
# 'Adjective'
# 'Adverb'
# 'Alternate forms'
# 'Alternative form'
# 'Alternative forms'
# 'Anagrams'
# 'Article'
# 'Circumfix'
# 'Citations'
# 'Conjunction'
# 'Contraction'
# 'Declension'
# 'Derived terms'
# 'Descendants'
# 'Determiner'
# 'Diacritical mark'
# 'Etymology'
# 'Etymology 1'
# 'Etymology 2'
# 'Etymology 3'
# 'Etymology 4'
# 'Etymology 5'
# 'External links'
# 'Further reading'
# 'Gerund'
# 'Infix'
# 'Inflection'
# 'Initialism'
# 'Interfix'
# 'Interjection'
# 'Letter'
# 'Noun'
# 'Numeral'
# 'Ordinal number'
# 'Participle'
# 'Particle'
# 'Phrase'
# 'Postposition'
# 'Prefix'
# 'Preposition'
# 'Prepositional phrase'
# 'Prnonuciation'
# 'Pronoun'
# 'Pronunciation'
# 'Pronunciation 1'
# 'Pronunciation 2'
# 'Proper noun'
# 'Proverb'
# 'Punctuation mark'
# 'Quotations'
# 'References'
# 'Related terms'
# 'See also'
# 'Suffix'
# 'Symbol'
# 'Synonyms'
# 'Usage notes'
# 'Verb'