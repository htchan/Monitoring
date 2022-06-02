file_pattern="/log/resources/docker-*.scope/memory.current"

mem_files=$(ls $file_pattern)

for file in $mem_files
do
    replace_string=$(echo "$file_pattern" | sed "s/*/(.*)/g" | sed -e 's/\//\\\//g')
    docker_id=$(echo "$file" | sed -r "s/$replace_string/\1/g")
    content=$(cat $file)

    echo $docker_id mem $content
done
