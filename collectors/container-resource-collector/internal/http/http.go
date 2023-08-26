package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/htchan/ContainerResourceCollector/internal/container"
)

const SERVER_PORT = ":8000"

func ServeHTTP(containerMap *map[string]*container.Container) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, c := range *containerMap {
			// CPU usage
			fmt.Fprintf(w, "# HELP docker_container_cpu_usage Docker container CPU usage\n")
			fmt.Fprintf(w, "# TYPE docker_container_cpu_usage gauge\n")
			fmt.Fprintf(w, "docker_container_cpu_usage{container_id=\"%s\",project=\"%s\",service=\"%s\"} %d\n", c.ContainerID, c.Project, c.Service, c.Resources.CPU.Usage)

			// Memory usage
			fmt.Fprintf(w, "# HELP docker_container_memory_usage Docker container Memory usage\n")
			fmt.Fprintf(w, "# TYPE docker_container_memory_usage gauge\n")
			fmt.Fprintf(w, "docker_container_memory_usage{container_id=\"%s\",project=\"%s\",service=\"%s\"} %d\n", c.ContainerID, c.Project, c.Service, c.Resources.Memory.UsageBytes)

			// Memory Limit
			fmt.Fprintf(w, "# HELP docker_container_memory_limit Docker container Memory Limit\n")
			fmt.Fprintf(w, "# TYPE docker_container_memory_limit gauge\n")
			fmt.Fprintf(w, "docker_container_memory_limit{container_id=\"%s\",project=\"%s\",service=\"%s\"} %d\n", c.ContainerID, c.Project, c.Service, c.Resources.Memory.LimitBytes)

			for key, network := range c.Resources.Network {
				// Network Transmit
				fmt.Fprintf(w, "# HELP docker_container_network_receive Docker container Network Receive\n")
				fmt.Fprintf(w, "# TYPE docker_container_network_receive gauge\n")
				fmt.Fprintf(w, "docker_container_network_receive{container_id=\"%s\",project=\"%s\",service=\"%s\",interface=\"%s\"} %d\n", c.ContainerID, c.Project, c.Service, key, network.ReceiveBytes)

				// Network Receive
				fmt.Fprintf(w, "# HELP docker_container_network_transmit Docker container Network Transmit\n")
				fmt.Fprintf(w, "# TYPE docker_container_network_transmit gauge\n")
				fmt.Fprintf(w, "docker_container_network_transmit{container_id=\"%s\",project=\"%s\",service=\"%s\",interface=\"%s\"} %d\n", c.ContainerID, c.Project, c.Service, key, network.TransmitBytes)
			}

		}
	})

	log.Printf("Start server at port %s", SERVER_PORT)
	if err := http.ListenAndServe(SERVER_PORT, nil); err != nil {
		return err
	}

	return nil
}
