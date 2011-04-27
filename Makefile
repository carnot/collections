include $(GOROOT)/src/Make.inc

TARG=badgerodon/collections
GOFILES=\
  trie.go\
  ternary_search_tree.go\

include $(GOROOT)/src/Make.pkg
