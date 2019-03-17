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

def load_wiktionary(filename='enwiktionary-latest-pages-articles.xml'):  #, pickle_name='wiktionary-articles.pickle'):
  #if not os.path.isfile(pickle_name):
  print('Parsing Wiktionary XML...')
  with Timer() as t:
    tree = et.parse(filename)
  print('Took {} secs'.format(t.interval))

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

#print('root tag: ', root.tag)

print('Extracting Latin entries...')
for c in root:
  tag = c.tag.split('}', 1)[1]
  #print('tag:', tag, 'attrib:', c.attrib)

  # <page> contains entries
  if tag == "page":
    done = False
    found_latin = False
    entry = {'title': None, 'timestamp': None}
    entry_text = None
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
          if tag2 == 'timestamp':
            entry['timestamp'] = ccc.text
          elif tag2 == 'text':
            section_str = re.compile('([^=]==[^=]+==[^=])')
            if ccc.text is not None:
              sections = section_str.split(ccc.text)
              # Break up the text at ==...== sections
              for idx, s in enumerate(sections):
                # If this one corresponds to a section header, then next must be text
                if section_str.match(s) is not None:
                  #print(s.strip())
                  if s.strip() == '==Latin==':
                    if idx+1 < len(sections):
                      #print(type(sections[idx+1].strip()))
                      #raise Exception()
                      entry_text = sections[idx+1].strip() #.decode('utf-8')
                    else:
                      entry_text = u''
                    found_latin = True
                    count_found += 1
                
    if found_latin:
      entry_node = et.SubElement(output_root, u'entry', entry)
      entry_node.text = entry_text

      if count_found % 1000 == 0 and count_found != 0:
        print('{} words found...'.format(count_found))

      #print(entry)

et.ElementTree(element=output_root).write(open(u'wiktionary-latin.xml', 'wb'), encoding='UTF-8')

#with open(u'wiktionary-latin.xml', 'wb') as f:
#  f.write(et.tostring(output_root)) #.decode("utf-8"))
