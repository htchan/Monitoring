file_pattern="/log/resources/docker-*.scope/cpu.stat"

cpu_files=$(ls $file_pattern)

for file in $cpu_files
do
    replace_string=$(echo "$file_pattern" | sed "s/*/(.*)/g" | sed -e 's/\//\\\//g')
    docker_id=$(echo "$file" | sed -r "s/$replace_string/\1/g")
    content=$(cat $file | grep -E '^usage_usec ' | sed "s/usage_usec //g")

    echo $docker_id cpu $content
done
