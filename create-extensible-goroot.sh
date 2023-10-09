
export GOBASEPATH=/usr/local/go
export GOSRCPATH=${GOBASEPATH}/src

export DUMMYGOROOT=/tmp/dummy-go-root
export DUMMYGOROOTSRC=${DUMMYGOROOT}/src

rm -rf ${DUMMYGOROOT}

mkdir -p ${DUMMYGOROOT}

for object in $(ls -A1 $GOBASEPATH | grep -v src) 
do
    ln -s $GOBASEPATH/$object $DUMMYGOROOT/$object 
done

mkdir -p ${DUMMYGOROOTSRC}

for object in $(ls -A1 $GOSRCPATH) 
do
    ln -s $GOSRCPATH/$object $DUMMYGOROOTSRC/$object 
done

export DUMMY_MACHINE=$(realpath $DUMMYGOROOTSRC)/machine
 
ln -s "$(pwd)/dummy/machine" ${DUMMY_MACHINE}

echo --------------------------------------------------
echo export GOROOT=$DUMMYGOROOT
echo --------------------------------------------------


