all_files=$(ls /log/sys/fs/cgroup/cpu,cpuacct/docker/*/cpuacct.usage)

for file in $all_files
do
    update_filename=$(echo $file | sed -e 's/\//\\\//g')
    sed "s/^/$update_filename /" $file | cat
done

