for d in ../internal/*/ ; do
    echo "generating for " $d
    cd $d
    go generate
    cd -
done