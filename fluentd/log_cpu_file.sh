id=$(for i in {1..30}; do echo -n "?"; done)

cpu_files=$(ls /log/resources/cpu,cpuacct/docker/$id*/cpuacct.usage)

for file in $cpu_files
do
    update_filename=$(echo $file | sed -e 's/\//\\\//g')
    sed "s/^/$update_filename /" $file | cat
done
