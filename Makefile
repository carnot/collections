DIRS = queue set splay stack tst

all: install

clean.dirs: $(addsuffix .clean, $(DIRS))
install.dirs: $(addsuffix .install, $(DIRS))
test.dirs: $(addsuffix .test, $(DIRS))

%.clean:
	+cd $* && make clean

%.install:
	+cd $* && make install

%.test:
	+cd $* && make test

clean: clean.dirs

install: install.dirs

test: test.dirs