for i in `ls lists/bl_*`
do
    go run british_library_tag.go -in_file $i
done
