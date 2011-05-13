DIRS = pheap pqueue queue splay stack trie tst

all: install

clean.dirs: $(addsuffix .clean, $(DIRS))
install.dirs: $(addsuffix .install, $(DIRS))
test.dirs: $(addsuffix .test, $(DIRS))

%.clean:
	+cd $* && gomake clean

%.install:
	+cd $* && gomake install

%.test:
	+cd $* && gomake test

clean: clean.dirs

install: install.dirs

test: test.dirs