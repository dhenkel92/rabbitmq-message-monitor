package collector

import (
	"io"
	"strconv"

	"github.com/dhenkel92/rabbitmq-message-monitor/internal/helper"
	"github.com/olekukonko/tablewriter"
)

func (collector *Collector) PrintTable(out io.Writer) {
	table := tablewriter.NewWriter(out)
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	table.SetHeader([]string{"Routing Key", "Count", "Total Size"})

	collector.mu.Lock()
	for _, entry := range collector.routingKeyStats {
		table.Append([]string{entry.RoutingKey, strconv.FormatInt(int64(entry.Count), 10), helper.FormatBytesToMb(entry.TotalBytes)})
	}
	collector.mu.Unlock()

	table.Render()
}
