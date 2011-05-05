include $(GOROOT)/src/Make.inc

TARG=badgerodon/collections
GOFILES=\
  interfaces.go\
  trie.go\
  ternary_tree.go\

include $(GOROOT)/src/Make.pkg
