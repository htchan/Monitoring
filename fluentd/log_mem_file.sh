id=$(for i in {1..30}; do echo -n "?"; done)

mem_files=$(ls /log/resources/memory/docker/$id*/memory.stat)

for file in $mem_files
do
    content=$(cat $file | grep -E '^cache |^rss ' | sed "s/[a-z]//g")
    sum=0
    for num in $content
    do
	    sum=$(expr $sum + $num)
    done
    echo $file $sum
done
