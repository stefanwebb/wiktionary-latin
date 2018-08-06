import time, re
import pickle
import os.path
import xml.etree.ElementTree as et

# The purpose of this script is to extract Wiktionary entries that contain both English and Latin definitions, to compile a list of English words that have been taken directly from Latin

class Timer:
  def __enter__(self):
    self.start = time.clock()
    return self

  def __exit__(self, *args):
    self.end = time.clock()
    self.interval = self.end - self.start

def load_wiktionary(filename='enwiktionary-latest-pages-articles.xml'):  #, pickle_name='wiktionary-articles.pickle'):
  #if not os.path.isfile(pickle_name):
  print('Parsing Wiktionary XML...')
  with Timer() as t:
    tree = et.parse(filename)
  print('Took {} mins'.format(t.interval/60))

  # print('Saving to pickle...')
  # with open(pickle_name, 'wb') as f:
  #   with Timer() as t:
  #     pickle.dump(tree, f, pickle.HIGHEST_PROTOCOL)
  # print('Took {} secs'.format(t.interval))

  # else:
  #   print('Reading Wiktionary XML from pickle...')
  #   with Timer() as t:
  #     with open(pickle_name, 'rb') as f:
  #       tree = pickle.load(f)
  #   print('Took {} secs'.format(t.interval))

  return tree

# Print out some stuff from the XML parse tree
#tree = load_wiktionary(filename='test.xml', pickle_name='test.pickle')
#tree = load_wiktionary(filename='wiktionary-extract.xml')
tree = load_wiktionary()
root = tree.getroot()
output_root = et.Element(u'wiktionary')
count_found = 0

languages = {}

#print('root tag: ', root.tag)

print('Extracting entries containing both Latin and English entries...')
for c in root:
  tag = c.tag.split('}', 1)[1]
  #print('tag:', tag, 'attrib:', c.attrib)

  # <page> contains entries
  if tag == "page":
    done = False
    found_latin = False
    found_english = False
    entry = {'title': None}
    #entry_text = None
    for cc in c:
      tag = cc.tag.split('}', 1)[1]
      if tag == 'title':
        entry['title'] = cc.text
        # <page>/<revision> contains all the juicy stuff
      elif tag == 'revision':
        # Throw error if there is more than one revision
        assert(not done)
        done = True
        for ccc in cc:
          tag2 = ccc.tag.split('}', 1)[1]
          #if tag2 == 'timestamp':
          #  entry['timestamp'] = ccc.text
          if tag2 == 'text':
            section_str = re.compile('(^==[^=]+?==$)', re.MULTILINE)
            if ccc.text is not None:
              sections = section_str.split(ccc.text)
              # Break up the text at ==...== sections
              for idx, s in enumerate(sections):
                # If this one corresponds to a section header, then next must be text
                if section_str.match(s) is not None:
                  # Make record of languages (just curious and also checking for typos!)
                  if s.strip() not in languages:
                    languages[s.strip()] = 1
                  else:
                    languages[s.strip()] += 1

                  if s.strip() == '==Latin==':
                    if idx+1 < len(sections):
                      latin_text = sections[idx+1].strip()
                    else:
                      latin_text = u''
                    found_latin = True
                  elif s.strip() == '==English==':
                    if idx+1 < len(sections):
                      english_text = sections[idx+1].strip()
                    else:
                      english_text = u''
                    found_english = True
                
    if found_latin and found_english:
      #print(entry['title'])
      count_found += 1
      entry_node = et.SubElement(output_root, u'entry', entry)
      english_node = et.SubElement(entry_node, u'English')
      english_node.text = english_text
      latin_node = et.SubElement(entry_node, u'Latin')
      latin_node.text = latin_text

      if count_found % 1000 == 0 and count_found != 0:
        print('{} words found...'.format(count_found))

      #print(entry)

#print(languages)

print('Saving results...')
et.ElementTree(element=output_root).write(
    open(u'd:/onedrive/wiktionary-latin-english.xml', 'wb'), encoding='UTF-8')
print('Done!')

with open('d:/onedrive/wiktionary-language-tags.txt', 'w') as f:
  for k, v in sorted(languages.items(), key=lambda key_value: key_value[0]):
    print(k.encode('utf8'), v.encode('utf8'), file=f)

#with open(u'wiktionary-latin.xml', 'wb') as f:
#  f.write(et.tostring(output_root)) #.decode("utf-8"))
